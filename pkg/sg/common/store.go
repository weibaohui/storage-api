// 存储系统管理，包括 列表
package common

import (
	"encoding/json"
	"storage-api/pkg/sg"
)

type StoreList struct {
	sg.ErrorMsg
	Data []*Store `json:"result"`
}

type Store struct {
	ActiveAlarmsNum     int          `json:"activeAlarmsNum"`
	ClusterDataState    string       `json:"clusterDataState"`
	ClusterHealthyState string       `json:"clusterHealthyState"`
	ClusterRunningState string       `json:"clusterRunningState"`
	Name                string       `json:"name"`
	OjmgsIps            string       `json:"ojmgs_ips"`
	Sysid               int          `json:"sysid"`
	UUID                string       `json:"uuid"`
	Version             string       `json:"version"`
	ZkServers           []*ZkServers `json:"zk_servers"`
}

type ZkServers struct {
	Port  int      `json:"port"`
	ZkID  int      `json:"zkId"`
	ZkIps []string `json:"zkIps"`
}

//POST
//登录cookie
func (i *Instance) ListStore() (*StoreList, error) {
	url := i.FullURL("/install/getStorageSystemsOverview.action?user_name=" + i.Username)
	str, err := i.PostWithLoginSession(url, nil)
	if err != nil {
		return nil, err
	}
	result := &StoreList{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result, nil
}

func (i *Instance) DefaultStore() (*Store, error) {
	list, err := i.ListStore()
	if err != nil {
		return nil, err
	}
	return list.Data[0], nil
}
