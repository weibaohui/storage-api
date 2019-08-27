// 存储系统管理，包括 列表
package common

import (
	"encoding/json"
	"errors"
	"log"
	"storage-api/pkg/sg"
)

type connectResult struct {
	sg.ErrorMsg
	Data struct {
		sg.ErrorMsg
		Cluster interface{} `json:"result"`
	} `json:"result"`
}

type ClusterList struct {
	sg.ErrorMsg
	Data []*Cluster `json:"result"`
}

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
type Cluster struct {
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

type ZkServers struct {
	Port  int      `json:"port"`
	ZkID  int      `json:"zkId"`
	ZkIps []string `json:"zkIps"`
}

//列表当前系统可用的集群
//需要先进行连接，连接后才能列表展示
//没有连接的话，可以执行连接
func (i *Instance) ListCluster() (*ClusterList, error) {
	url := i.FullURL("/install/getStorageSystemsOverview.action?user_name=" + i.Username)
	str, err := i.PostWithLoginSession(url, nil)
	if err != nil {
		return nil, err
	}
	result := &ClusterList{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result, nil
}

//默认的第一个存储集群
func (i *Instance) DefaultCluster() (*Cluster, error) {
	list, err := i.ListCluster()
	if err != nil {
		return nil, err
	}
	if len(list.Data) == 0 {
		ok, err := i.connectToExistCluster()
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("程序尝试连接存储集群但失败了。请先在界面连接存储集群。")
		}
		return i.DefaultCluster()
	}
	return list.Data[0], nil
}

//连接到存储系统
func (i *Instance) connectToExistCluster() (ok bool, err error) {
	log.Println("connectToExistCluster:", i.Config.Host)
	url := i.FullURL("/install/connectToExistCluster.action?user_name=" + i.Username)
	params := make(map[string]string)
	params["oJmgsIPs"] = i.Config.Host
	str, err := i.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &connectResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}
	if result.Data.Cluster != nil {
		return true, nil
	}

	return false, errors.New("连接到存储系统失败")
}
