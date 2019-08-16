package sg

import (
	"fmt"
	"testing"
)

func TestCookies(t *testing.T) {
	robot := FakeRobot()
	cookies, err := robot.AuthCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(cookies)
}

func TestStoreList(t *testing.T) {
	robot := FakeRobot()
	list, err := robot.StoreList()
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
	robot := FakeRobot()
	store, err := robot.DefaultStore()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(store.Name, store.UUID)
	fmt.Println("ClusterDataState=", store.ClusterDataState)
	fmt.Println("ClusterRunningState=", store.ClusterRunningState)
	fmt.Println("ClusterHealthyState=", store.ClusterHealthyState)
}

func TestQuotaList(t *testing.T) {
	robot := FakeRobot()
	list, err := robot.QuotaList()
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range list.Data.Quotas {
		fmt.Println(v)

	}

}

func TestCreateQuota(t *testing.T) {
	robot := FakeRobot()
	json, err := robot.CreateQuota("/nfs", 55, 66, 77, 88)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(json)
}
