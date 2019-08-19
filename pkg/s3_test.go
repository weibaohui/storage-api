package pkg

import (
	"fmt"
	"nfs-api/pkg/sg"
	"testing"
)

func TestDeleteAccount(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	id := ""
	if ok, err := api.DeleteAccount(id); err == nil {
		fmt.Println("删除账户结果", ok)
	} else {
		t.Fatal("删除账户失败", err.Error())
	}
}

func TestDeleteCertificate(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	id := "/test/4455"
	if ok, err := api.DeleteCertificate(id); err == nil {
		fmt.Println("删除证书结果", ok)
	} else {
		t.Fatal("删除证书失败")
	}
}
func TestRunS3(t *testing.T) {
	api := sg.NewInstance("https", "192.168.3.60", "6080", "optadmin", "adminadmin")

	accountName := "zhangsanfeng"
	if accountID, err := api.CreateAccount(accountName, 0); err == nil {
		fmt.Println("创建账户结果", accountID)
		if ak, sk, err := api.CreateCertificate(accountID); err == nil {
			fmt.Println("创建证书成功", ak, sk)
			if ok, err := api.DeleteCertificate(ak); err == nil {
				fmt.Println("删除证书", ok)
				if ok, err := api.DeleteAccount(accountID); err == nil {
					fmt.Println("删除账户", ok)
				}
			}
		}

	} else {
		t.Fatal("创建账户失败", err.Error())
	}

}
