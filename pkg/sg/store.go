// 存储系统管理，包括 列表
package sg

import (
	"encoding/json"
	"errors"
)

type StoreList struct {
	ErrorMsg
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
func (r *Robot) StoreList() (*StoreList, error) {
	url := r.fullURL("/install/getStorageSystemsOverview.action?user_name=" + r.Username)

	j, err := r.PostWithLoginSession(url, nil)
	if err != nil {
		return nil, err
	}
	list := &StoreList{}
	err = json.Unmarshal([]byte(j), list)
	if err != nil {
		return nil, err
	}
	if list.ErrNo != 0 {
		return nil, errors.New(list.ErrorString())
	}
	return list, nil
}

func (r *Robot) DefaultStore() (*Store, error) {
	list, err := r.StoreList()
	if err != nil {
		return nil, err
	}
	return list.Data[0], nil
}
