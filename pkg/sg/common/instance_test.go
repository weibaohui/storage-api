package common

import (
	"fmt"
	"storage-api/pkg/api"
	"testing"
)

var common *Instance

func init() {
	config := &api.Config{
		Protocol: "https",
		Host:     "192.168.3.60",
		Port:     "6080",
		Username: "optadmin",
		Password: "adminadmin",
	}
	common = NewInstance(config)
}
func TestCookies(t *testing.T) {
	cookies, err := common.AuthCookie()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(cookies)
}

func TestListCluster(t *testing.T) {
	list, err := common.ListCluster()
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range list.Data {
		fmt.Println(v.Name, v.UUID)
	}

}

func TestDefaultCluster(t *testing.T) {
	store, err := common.DefaultCluster()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(store.Name, store.UUID)
}

func TestInstance_ClusterStatus(t *testing.T) {
	status, err := common.ClusterStatus()
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(status.Name)
	fmt.Println(status.UUID)
	fmt.Println(status.ActiveAlarmsNum)
	fmt.Println(status.ClusterDataState)
	fmt.Println(status.ClusterHealthyState)
	fmt.Println(status.ClusterRunningState)
	fmt.Println(status.NodesNumber)

}

func TestListAccessZones(t *testing.T) {
	list, err := common.ListAccessZones()
	if err != nil {
		t.Fatal(err.Error())
	}
	for k, v := range list {
		fmt.Println(k, v.ID, v.Name)
	}
}
