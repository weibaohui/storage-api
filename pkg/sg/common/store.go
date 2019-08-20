// 存储系统管理，包括 列表
package common

import (
	"encoding/json"
	"nfs-api/pkg/sg"
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
//https://192.168.3.60:6080/install/getStorageSystemsOverview.action?cmd_id=0.4124075744799065&user_name=optadmin
func (r *Instance) ListStore() (*StoreList, error) {
	url := r.FullURL("/install/getStorageSystemsOverview.action?user_name=" + r.Username)
	str, err := r.PostWithLoginSession(url, nil)
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

func (r *Instance) DefaultStore() (*Store, error) {
	list, err := r.ListStore()
	if err != nil {
		return nil, err
	}
	return list.Data[0], nil
}
