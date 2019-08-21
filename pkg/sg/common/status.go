package common

import (
	"encoding/json"
	"storage-api/pkg/sg"
)

//"cluster_data_state":"SYSTEM_NORMAL",// 存储数据状态
//      SYSTEM_FAULT:故障
//      SYSTEM_DEGRADE:降级
//      SYSTEM_NORMAL:正常
//"cluster_healthy_state":"NORMAL",// 存储健康状态
//       NORMAL:健康
//		 ABNORMAL:存在告警即为非健康
//"cluster_running_state":"SYSTEM_RUNNING",//存储运行状态
//		SYSTEM_RUNNNING:运行中
//		SYSTEM_SHUTDOWN:已关闭
//		SYSTEM_STARTING:启动中
//		SYSTEM_PREPARE_SHUTDOWN:准备关闭
//		SYSTEM_SHUTTING_DOWN:关闭中
//		SYSTEM_SHUTDOWN_FAILED:关闭失败
//		SYSTEM_SHUTTING_DOWN_FORCE:强制关闭中
//		SYSTEM_UPGRADING:升级中

type ClusterStatus struct {
	Name                string       `json:"name"` //集群名称
	UUID                string       `json:"uuid"`
	ActiveAlarmsNum     int          `json:"active_alarms_num"`     //实时告警
	ClusterDataState    string       `json:"cluster_data_state"`    //存储数据状态
	ClusterHealthyState string       `json:"cluster_healthy_state"` //存储健康状态
	ClusterRunningState string       `json:"cluster_running_state"` //存储运行状态
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
