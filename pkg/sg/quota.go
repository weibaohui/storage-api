// 配额管理
package sg

import (
	"encoding/json"
	"errors"
)

type Quota struct {
	AuthProviderID                int    `json:"auth_provider_id"`
	Description                   string `json:"description"`
	FilenrGraceTime               int    `json:"filenr_grace_time"`
	FilenrHardThreshold           int    `json:"filenr_hard_threshold"`
	FilenrQuotaCalType            string `json:"filenr_quota_cal_type"`
	FilenrSoftThreshold           int    `json:"filenr_soft_threshold"`
	FilenrSoftThresholdOverTime   string `json:"filenr_soft_threshold_over_time"`
	FilenrSuggestThreshold        int    `json:"filenr_suggest_threshold"`
	FilenrUsedNr                  int    `json:"filenr_used_nr"`
	Fsid                          int    `json:"fsid"`
	ID                            int    `json:"id"`
	IpsQuota                      int    `json:"ips_quota"`
	IpsReal                       int    `json:"ips_real"`
	Key                           int    `json:"key"`
	LogicalGraceTime              int    `json:"logical_grace_time"`
	LogicalHardThreshold          int    `json:"logical_hard_threshold"`
	LogicalQuotaCalType           string `json:"logical_quota_cal_type"`
	LogicalSoftThreshold          int    `json:"logical_soft_threshold"`
	LogicalSoftThresholdOverTime  string `json:"logical_soft_threshold_over_time"`
	LogicalSuggestThreshold       int    `json:"logical_suggest_threshold"`
	LogicalUsedCapacity           int    `json:"logical_used_capacity"`
	OpsQuota                      int    `json:"ops_quota"`
	OpsReal                       int    `json:"ops_real"`
	Path                          string `json:"path"`
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
	Data           *QuotaQuery `json:"result"`
	Sync           bool        `json:"sync"`
	TimeZoneOffset int         `json:"time_zone_offset"`
	TraceID        string      `json:"trace_id"`
}

//POST
//登录cookie
//https://192.168.3.60:6080/commands/get_quota.action?cmd_id=0.5387214431814484&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
func (r *Robot) QuotaList(uuid string) (*QuotasList, error) {
	url := r.fullURL("/commands/get_quota.action?user_name=" + r.Username + "&uuid=" + uuid)

	params := make(map[string]string)
	params["rand"] = ""
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
