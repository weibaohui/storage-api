package nfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"nfs-api/pkg/api"
	"nfs-api/pkg/sg"
)

type ListSnapshotResult struct {
	sg.ErrorMsg
	Data struct {
		Limit     int             `json:"limit"`
		Searches  []interface{}   `json:"searches"`
		Snapshots []*api.Snapshot `json:"snapshots"`
		Sort      string          `json:"sort"`
		Start     int             `json:"start"`
		Total     int             `json:"total"`
	} `json:"result"`
}

//列表显示快照
///commands/get_snapshot.action
//params: {"limit":20,"start":0,"sort":""}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"limit":0,"searches":[],"snapshots":[{"create_time":1566271517,"create_user":"optadmin","description":"ddd描述","expire_time":1567049185,"id":1,"key":1,"name":"ttt","path":"ParaStor300S:/nfs","size":0,"state":"SNAPSHOT_WORKING"}],"sort":"NONE","start":0,"total":1},"sync":true,"time_stamp":1566279921297,"time_zone_offset":-480,"trace_id":"[423353242473 2]"}
func (i *instance) ListSnapshot() ([]*api.Snapshot, error) {
	config := fmt.Sprintf(`
	{"limit":20,"start":0,"sort":""}
	`)
	return i.listSnapshot(config)
}

//列表显示快照，按名称查询，名称唯一
func (i *instance) ListSnapshotByName(name string) ([]*api.Snapshot, error) {
	config := fmt.Sprintf(`
	{"limit":20,"start":0,"sort":"",
	"searches":[{"searchKey":"name","searchValue":"%s"}]}
	`, name)
	return i.listSnapshot(config)
}
func (i *instance) listSnapshot(config string) ([]*api.Snapshot, error) {
	url := i.common.Command("/commands/get_snapshot.action")
	params := make(map[string]string, 0)
	params["params"] = config
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	result := &ListSnapshotResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result.Data.Snapshots, nil
}

//创建快照
//commands/create_snapshot.action
//params: {"snapshots":[{"name":"ttt","path":"ParaStor300S:/nfs","description":"ddd描述","expire_time":"0","create_user":"optadmin"}]}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"job_id":1603862034062917,"job_id_str":"1603862034062917"},"sync":false,"time_stamp":1566271517649,"time_zone_offset":-480}
func (i *instance) CreateSnapshot(name, path, desc string, expireTime int) (id string, err error) {
	url := i.common.Command("/commands/create_snapshot.action")
	params := make(map[string]string, 0)
	params["params"] = fmt.Sprintf(`
	{"snapshots":[{
	"name":"%s",
	"path":"%s:%s",
	"description":"%s",
	"expire_time":"%d",
	"create_user":"%s"
	}]}
	`, name, i.common.StoreName, path, desc, expireTime, i.common.Username)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return "", err
	}
	result := &jobIDResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return "", err
	}
	if result.ErrNo != 0 {
		return "", result.Error()
	}

	//等待job执行完成
	_, err = i.isJobDone(result.Data.JobIDStr)
	if err != nil {
		return "", err
	}

	//找到快照ID
	snapshots, err := i.ListSnapshotByName(name)
	if err != nil {
		return "", err
	}
	for _, ss := range snapshots {
		if ss.Name == name {
			return fmt.Sprintf("%d", ss.ID), nil
		}
	}
	return "", errors.New("没有找到快照" + name)
}

//快照回滚
///commands/revert_snapshot.action
//params: {"snapshot_id":1}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"job_id":1603862185597684,"job_id_str":"1603862185597684"},"sync":false,"time_stamp":1566271665627,"time_zone_offset":-480}
func (i *instance) RevertSnapshot(id string) (ok bool, err error) {
	url := i.common.Command("/commands/revert_snapshot.action")
	params := make(map[string]string, 0)
	params["params"] = fmt.Sprintf(`{"snapshot_id":%s}`, id)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &jobIDResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}

	//等待job执行完成
	ok, err = i.isJobDone(result.Data.JobIDStr)
	if err != nil {
		return false, err
	}

	return ok, err
}

//快照删除
///commands/delete_snapshot.action
//params: {"ids":[3]}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"job_id":1603862185597684,"job_id_str":"1603862185597684"},"sync":false,"time_stamp":1566271665627,"time_zone_offset":-480}
func (i *instance) DeleteSnapshot(id string) (ok bool, err error) {
	url := i.common.Command("/commands/delete_snapshot.action")
	params := make(map[string]string, 0)
	params["params"] = fmt.Sprintf(`{"ids":[%s]}`, id)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &jobIDResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}

	//等待job执行完成
	ok, err = i.isJobDone(result.Data.JobIDStr)
	if err != nil {
		return false, err
	}

	return ok, err
}
