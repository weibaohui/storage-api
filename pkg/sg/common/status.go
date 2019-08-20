package common

import (
	"encoding/json"
	"storage-api/pkg/sg"
)

type ClusterStatus struct {
	Name                string       `json:"name"` //集群名称
	UUID                string       `json:"uuid"`
	ActiveAlarmsNum     int          `json:"active_alarms_num"`     //实时告警
	ClusterDataState    string       `json:"cluster_data_state"`    //数据状态
	ClusterHealthyState string       `json:"cluster_healthy_state"` //健康状态
	ClusterRunningState string       `json:"cluster_running_state"` //运行状态
	MonoCacheMode       int          `json:"mono_cache_mode"`
	NodepoolsNumber     int          `json:"nodepools_number"`    //节点池数量
	NodesNumber         int          `json:"nodes_number"`        //节点数量
	StoragepoolsNumber  int          `json:"storagepools_number"` //存储池数量
	VolumesNumber       int          `json:"volumes_number"`      //存储卷数量
	ZkServers           []*ZkServers `json:"zk_servers"`
	Version             string       `json:"version"` //系统版本
	Sysid               int          `json:"sysid"`
	OjmgsIps            string       `json:"ojmgs_ips"`
	PackageTime         string       `json:"packageTime"`
}
type ClusterStatusResult struct {
	sg.ErrorMsg
	Data *ClusterStatus `json:"result"`
}

//集群状态
func (i *Instance) ClusterStatus() (status *ClusterStatus, err error) {
	url := i.Command("/commands/get_cluster_overview.action")
	str, err := i.PostWithLoginSession(url, nil)
	result := &ClusterStatusResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result.Data, nil
}
