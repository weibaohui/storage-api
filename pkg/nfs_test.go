package pkg

import (
	"fmt"
	"nfs-api/pkg/sg"
	"testing"
	"time"
)

func TestDeleteDirectory(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	path := "/test/4455"
	if ok, err := api.DeleteDirectory(path); err == nil {
		fmt.Println("删除目录结果", ok)
	} else {
		t.Fatal("删除目录失败", err.Error())
	}
}

func TestDeleteQuota(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	quotaId := "/test/4455"
	if ok, err := api.DeleteQuota(quotaId); err == nil {
		fmt.Println("删除配额结果", ok)
	} else {
		t.Fatal("删除配额失败")
	}
}
func TestRun(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	path := "/test/4455"

	if ok, err := api.CreateDirectory(path); err == nil {
		fmt.Println("创建目录结果", ok)
	} else {
		t.Fatal("创建目录失败", err.Error())
	}

	if ok, quotaId, err := api.CreateQuota(path, 91, 92, 93, 94); err == nil {
		fmt.Println("创建配额结果", ok, "quotaID=", quotaId)
		if ok, err := api.DeleteQuota(quotaId); err == nil {
			fmt.Println("删除配额结果", ok)
			time.Sleep(time.Millisecond * 180)
		} else {
			t.Fatal("删除配额失败", err.Error())
		}
	} else {
		t.Fatal("创建配额失败", err.Error())
	}

	if ok, err := api.DeleteDirectory(path); err == nil {
		fmt.Println("删除目录结果", ok)
	} else {
		t.Fatal("删除目录失败", err.Error())
	}
}
