package pkg

import (
	"fmt"
	"nfs-api/pkg/api"
	"nfs-api/pkg/sg/common"
	"testing"
	"time"
)

var config *api.Config
var nfsApi api.NFSApi

func init() {
	config = &api.Config{
		Protocol: "https",
		Host:     "192.168.3.60",
		Port:     "6080",
		Username: "optadmin",
		Password: "adminadmin",
	}
	apiBig := common.NewInstance(config)
	nfsApi = apiBig.(api.NFSApi)
}
func TestDeleteDirectory(t *testing.T) {
	path := "/test/4455"
	if ok, err := nfsApi.DeleteDirectory(path); err == nil {
		fmt.Println("删除目录结果", ok)
	} else {
		t.Fatal("删除目录失败", err.Error())
	}
}

func TestDeleteQuota(t *testing.T) {
	quotaId := "60"
	if ok, err := nfsApi.DeleteQuota(quotaId); err == nil {
		fmt.Println("删除配额结果", ok)
	} else {
		t.Fatal("删除配额失败", err.Error())
	}
}
func TestRun(t *testing.T) {
	path := "/test/4455"

	if ok, err := nfsApi.CreateDirectory(path); err == nil {
		fmt.Println("创建目录结果", ok)
	} else {
		t.Fatal("创建目录失败", err.Error())
	}

	if ok, quotaId, err := nfsApi.CreateQuota(path, 91, 92, 93, 94); err == nil {
		fmt.Println("创建配额结果", ok, "quotaID=", quotaId)
		if ok, err := nfsApi.DeleteQuota(quotaId); err == nil {
			fmt.Println("删除配额结果", ok)
			time.Sleep(time.Millisecond * 300)
		} else {
			t.Fatal("删除配额失败", err.Error())
		}
	} else {
		t.Fatal("创建配额失败", err.Error())
	}

	if ok, err := nfsApi.DeleteDirectory(path); err == nil {
		fmt.Println("删除目录结果", ok)
	} else {
		t.Fatal("删除目录失败", err.Error())
	}
}
