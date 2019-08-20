package common

import (
	"fmt"
	"nfs-api/pkg/api"
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

func TestStoreList(t *testing.T) {
	list, err := common.ListStore()
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range list.Data {
		fmt.Println(v.Name, v.UUID)
		fmt.Println("ClusterDataState=", v.ClusterDataState)
		fmt.Println("ClusterRunningState=", v.ClusterRunningState)
		fmt.Println("ClusterHealthyState=", v.ClusterHealthyState)
	}

}

func TestDefaultStore(t *testing.T) {
	store, err := common.DefaultStore()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(store.Name, store.UUID)
	fmt.Println("ClusterDataState=", store.ClusterDataState)
	fmt.Println("ClusterRunningState=", store.ClusterRunningState)
	fmt.Println("ClusterHealthyState=", store.ClusterHealthyState)
}
