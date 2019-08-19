// 配额管理
package sg

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Quota struct {
	ID                            int    `json:"id"`
	Fsid                          int    `json:"fsid"`
	Key                           int    `json:"key"`
	Path                          string `json:"path"`
	AuthProviderID                int    `json:"auth_provider_id"`
	Description                   string `json:"description"`
	FilenrGraceTime               int    `json:"filenr_grace_time"`
	FilenrHardThreshold           int    `json:"filenr_hard_threshold"`
	FilenrQuotaCalType            string `json:"filenr_quota_cal_type"`
	FilenrSoftThreshold           int    `json:"filenr_soft_threshold"`
	FilenrSoftThresholdOverTime   string `json:"filenr_soft_threshold_over_time"`
	FilenrSuggestThreshold        int    `json:"filenr_suggest_threshold"`
	FilenrUsedNr                  int    `json:"filenr_used_nr"`
	IpsQuota                      int    `json:"ips_quota"`
	IpsReal                       int    `json:"ips_real"`
	LogicalGraceTime              int    `json:"logical_grace_time"`
	LogicalHardThreshold          int    `json:"logical_hard_threshold"`
	LogicalQuotaCalType           string `json:"logical_quota_cal_type"`
	LogicalSoftThreshold          int    `json:"logical_soft_threshold"`
	LogicalSoftThresholdOverTime  string `json:"logical_soft_threshold_over_time"`
	LogicalSuggestThreshold       int    `json:"logical_suggest_threshold"`
	LogicalUsedCapacity           int    `json:"logical_used_capacity"`
	OpsQuota                      int    `json:"ops_quota"`
	OpsReal                       int    `json:"ops_real"`
	PhysicalCountRedundantSpace   bool   `json:"physical_count_redundant_space"`
	PhysicalCountSnapshot         bool   `json:"physical_count_snapshot"`
	PhysicalGraceTime             int    `json:"physical_grace_time"`
	PhysicalHardThreshold         int    `json:"physical_hard_threshold"`
	PhysicalQuotaCalType          string `json:"physical_quota_cal_type"`
	PhysicalSoftThreshold         int    `json:"physical_soft_threshold"`
	PhysicalSoftThresholdOverTime string `json:"physical_soft_threshold_over_time"`
	PhysicalSuggestThreshold      int    `json:"physical_suggest_threshold"`
	PhysicalUsedCapacity          int    `json:"physical_used_capacity"`
	ReadBandwidthQuota            int    `json:"read_bandwidth_quota"`
	ReadBandwidthReal             int    `json:"read_bandwidth_real"`
	State                         string `json:"state"`
	UserOrGroupID                 int    `json:"user_or_group_id"`
	UserOrGroupName               string `json:"user_or_group_name"`
	UserType                      string `json:"user_type"`
	Version                       int    `json:"version"`
	WriteBandwidthQuota           int    `json:"write_bandwidth_quota"`
	WriteBandwidthReal            int    `json:"write_bandwidth_real"`
}
type QuotaQuery struct {
	Limit    int           `json:"limit"`
	Quotas   []*Quota      `json:"quotas"`
	Searches []interface{} `json:"searches"`
	Sort     string        `json:"sort"`
	Start    int           `json:"start"`
	Total    int           `json:"total"`
}
type QuotasList struct {
	ErrorMsg
	Data *QuotaQuery `json:"result"`
}

// 查询配额列表
//POST
//登录cookie
//https://192.168.3.60:6080/commands/get_quota.action?cmd_id=0.5387214431814484&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
func (r *Robot) QuotaList() (*QuotasList, error) {
	url := r.fullURL("/commands/get_quota.action?user_name=" + r.Username + "&uuid=" + r.uuid)

	params := make(map[string]string)
	params["params"] = "{\"limit\":10000,\"start\":0,\"sort\":\"\",\"data\":[]}"
	j, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	list := &QuotasList{}
	err = json.Unmarshal([]byte(j), list)
	if err != nil {
		return nil, err
	}
	if list.ErrNo != 0 {
		return nil, errors.New(list.ErrorString())
	}
	return list, nil
}

//设置配额
//POST
//登录cookie
//https://192.168.3.60:6080/commands/create_quota.action?cmd_id=0.5181687999132814&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//readBw writeBw Mb/s
func (r *Robot) CreateQuota(dir string, ips, ops, readBw, writeBw int) (*JobID, error) {
	url := r.fullURL("/commands/create_quota.action?user_name=" + r.Username + "&uuid=" + r.uuid)
	params := make(map[string]string)

	s := `{"quotas":[{
	"path":"%s",
	"auth_provider_id":"0",
	"user_type":"USERTYPE_NONE",
	"user_or_group_name":"",
	"description":"",
	"logical_quota_cal_type":"QUOTA_NONE",
	"logical_hard_threshold":0,
	"logical_soft_threshold":0,
	"logical_grace_time":"0",
	"logical_suggest_threshold":0,
	"filenr_quota_cal_type":"QUOTA_NONE",
	"filenr_hard_threshold":0,
	"filenr_soft_threshold":0,
	"filenr_grace_time":"0",
	"filenr_suggest_threshold":0,
	"physical_quota_cal_type":"QUOTA_NONE",
	"physical_hard_threshold":0,
	"physical_soft_threshold":0,
	"physical_grace_time":"0",
	"physical_suggest_threshold":0,
	"physical_count_redundant_space":false,
	"physical_count_snapshot":false,
	"ips_quota":"%d",
	"ops_quota":"%d",
	"read_bandwidth_quota":"%d",
	"write_bandwidth_quota":"%d",
	"user_or_group_id":""
	}]}`
	path := fmt.Sprintf("%s:%s", r.storeName, dir)
	config := fmt.Sprintf(s, path, ips, ops, readBw, writeBw)
	params["params"] = config
	str, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	jobIDResult := &JobIDResult{}
	err = json.Unmarshal([]byte(str), jobIDResult)
	if err != nil {
		return nil, err
	}
	if jobIDResult.ErrNo != 0 {
		return nil, errors.New(jobIDResult.ErrorString())
	}
	return jobIDResult.Data, nil
}

//删除配额
//POST
//https://192.168.3.60:6080/commands/delete_quota.action?cmd_id=0.5855324522870262&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//rand:
//params: {"ids":[6]}
func (r *Robot) DeleteQuota(id int) (*JobID, error) {
	url := r.fullURL("/commands/delete_quota.action?user_name=" + r.Username + "&uuid=" + r.uuid)
	params := make(map[string]string)
	params["params"] = fmt.Sprintf("{\"ids\":[%d]}", id)
	str, err := r.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	jobIDResult := &JobIDResult{}
	err = json.Unmarshal([]byte(str), jobIDResult)
	if err != nil {
		return nil, err
	}
	if jobIDResult.ErrNo != 0 {
		return nil, errors.New(jobIDResult.ErrorString())
	}
	return jobIDResult.Data, nil
}
