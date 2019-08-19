package sg

import (
	"fmt"
	"testing"
)

func TestCookies(t *testing.T) {
	robot := FakeRobot4Test()
	cookies, err := robot.AuthCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(cookies)
}

func TestStoreList(t *testing.T) {
	robot := FakeRobot4Test()
	list, err := robot.ListStore()
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
	robot := FakeRobot4Test()
	store, err := robot.DefaultStore()
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(store.Name, store.UUID)
	fmt.Println("ClusterDataState=", store.ClusterDataState)
	fmt.Println("ClusterRunningState=", store.ClusterRunningState)
	fmt.Println("ClusterHealthyState=", store.ClusterHealthyState)
}

func TestCreateListDeleteQuota(t *testing.T) {
	robot := FakeRobot4Test()
	done, err := robot.CreateQuota("/nfs", 55, 66, 77, 88)
	if err != nil {
		fmt.Println("创建配额失败", err.Error())
		t.Fatal(err.Error())
	}

	fmt.Println("创建配额结果", done)
	if done {
		list, err := robot.ListQuota()
		if err != nil {
			t.Fatal(err.Error())
		}
		for _, v := range list.Data.Quotas {
			fmt.Println("删除配额,配额ID=", v.ID)
			done, err := robot.DeleteQuota(fmt.Sprintf("%d", v.ID))
			if err != nil {
				fmt.Println("删除配额,配额ID=", v.ID, err.Error())
				t.Fatal(err.Error())
			}
			fmt.Println("删除结果", done)
		}
	}
}
func TestQuotaList(t *testing.T) {
	robot := FakeRobot4Test()
	list, err := robot.ListQuota()
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, v := range list.Data.Quotas {
		fmt.Println(v.ID)
	}
}
func TestCreateQuota(t *testing.T) {
	robot := FakeRobot4Test()
	done, err := robot.CreateQuota("/nfs", 55, 66, 77, 88)
	if err != nil {
		fmt.Println("创建配额失败", err.Error())
		t.Fatal(err.Error())
	}
	fmt.Println("创建配额结果", done)
}
func TestDeleteQuota(t *testing.T) {
	robot := FakeRobot4Test()
	done, err := robot.DeleteQuota(43)
	if err != nil {
		fmt.Println("删除", err.Error())
		return
	}
	fmt.Println("删除结果", done)
}

func TestCreateDeleteDirectory(t *testing.T) {
	robot := FakeRobot4Test()
	path := "/test5/dddttt"
	created, err := robot.CreateDirectory(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if created {
		fmt.Println(path, "创建成功")
		deleted, err := robot.DeleteDirectory(path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(path, "删除结果", deleted)
	}
	fmt.Println(created)
}
func TestCreateDirectory(t *testing.T) {
	robot := FakeRobot4Test()
	path := "/test/dddttt55d"
	created, err := robot.CreateDirectoryWithPermission(path, "rwxrw-rw-")
	if err != nil {
		fmt.Println(err)
		return
	}
	if created {
		fmt.Println(path, "创建成功")

	}
}
func TestListDirectory(t *testing.T) {
	robot := FakeRobot4Test()
	path := "/nfs/"
	list, err := robot.ListDirectory(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		fmt.Printf("%s\t%s\t%s \t%s \n", v.PosixPath, v.Path, v.PosixPermission, v.Type)
	}
}
func TestListDirectoryWithFiles(t *testing.T) {
	robot := FakeRobot4Test()
	path := "/nfs/"
	list, err := robot.ListDirectoryWithFiles(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		fmt.Printf("%s\t%s\t%s \t%s \n", v.PosixPath, v.Path, v.PosixPermission, v.Type)
	}
}
