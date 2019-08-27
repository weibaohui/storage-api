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
		Store *Store `json:"result"`
	} `json:"result"`
}

type StoreList struct {
	sg.ErrorMsg
	Data []*Store `json:"result"`
}

type Store struct {
	ActiveAlarmsNum     int    `json:"activeAlarmsNum"`
	ClusterDataState    string `json:"clusterDataState"`
	ClusterHealthyState string `json:"clusterHealthyState"`
	ClusterRunningState string `json:"clusterRunningState"`
	Name                string `json:"name"`
	OjmgsIps            string `json:"ojmgs_ips"`
	//Sysid               int          `json:"sysid"`
	UUID      string       `json:"uuid"`
	Version   string       `json:"version"`
	ZkServers []*ZkServers `json:"zk_servers"`
}

type ZkServers struct {
	Port  int      `json:"port"`
	ZkID  int      `json:"zkId"`
	ZkIps []string `json:"zkIps"`
}

//列表当前系统可用的集群
//需要先进行连接，连接后才能列表展示
//没有连接的话，可以执行连接
func (i *Instance) ListCluster() (*StoreList, error) {
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

//默认的第一个存储集群
func (i *Instance) DefaultCluster() (*Store, error) {
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
	if result.Data.Store != nil {
		return true, nil
	}

	return false, errors.New("连接到存储系统失败")
}
