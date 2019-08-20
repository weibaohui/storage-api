package nfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"storage-api/pkg/api"
	"storage-api/pkg/sg"
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
func (i *instance) ListSnapshot() ([]*api.Snapshot, error) {
	config := fmt.Sprintf(`
	{"limit":999999,"start":0,"sort":""}
	`)
	return i.listSnapshot(config)
}

//列表显示快照，按名称查询，名称唯一
func (i *instance) ListSnapshotByName(name string) ([]*api.Snapshot, error) {
	config := fmt.Sprintf(`
	{"limit":999999,"start":0,"sort":"",
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
