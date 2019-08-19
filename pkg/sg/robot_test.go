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
		fmt.Println(v.ID)
	}
}

func TestCreateListDeleteQuota(t *testing.T) {
	robot := FakeRobot()
	jobID, err := robot.CreateQuota("/nfs", 55, 66, 77, 88)
	if err != nil {
		fmt.Println("创建配额失败", err.Error())
		t.Fatal(err.Error())
	}
	done, err := robot.IsJobDone(jobID.JobIDStr)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("创建配额结果", done)
	if done {
		list, err := robot.QuotaList()
		if err != nil {
			t.Fatal(err.Error())
		}
		for _, v := range list.Data.Quotas {
			fmt.Println("删除配额,配额ID=", v.ID)
			jobID, err := robot.DeleteQuota(v.ID)
			if err != nil {
				fmt.Println("删除配额,配额ID=", v.ID, err.Error())
				t.Fatal(err.Error())
			}
			done, err := robot.IsJobDone(jobID.JobIDStr)
			if err != nil {
				t.Fatal(err.Error())
			}
			fmt.Println("删除结果", done)
		}
	}
}
func TestDeleteQuota(t *testing.T) {
	robot := FakeRobot()
	jobID, err := robot.DeleteQuota(9)
	if err != nil {
		fmt.Println("删除", err.Error())
		return
	}
	done, err := robot.IsJobDone(jobID.JobIDStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("删除结果", done)
}
