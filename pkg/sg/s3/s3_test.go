package s3

import (
	"fmt"
	"nfs-api/pkg/api"
	"testing"
)

var s3api api.S3Api

func init() {
	config := &api.Config{
		Protocol: "https",
		Host:     "192.168.3.60",
		Port:     "6080",
		Username: "optadmin",
		Password: "adminadmin",
	}
	s3api = NewInstance(config)
}
func TestListS3Accounts(t *testing.T) {
	accounts, err := s3api.ListAccount()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range accounts {
		fmt.Println(v.AccountID, v.AccountName, v.AccountQuota, v.UsedBytes)
	}
}
func TestCreateS3Accounts(t *testing.T) {

	accountID, err := s3api.CreateAccount("testttt", 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("创建S3账户ID=", accountID)
}
func TestCreateCertificate(t *testing.T) {

	ak, sk, err := s3api.CreateCertificate("2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCGX")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("创建S3账户", ak, sk)
}
func TestListCertificate(t *testing.T) {

	infos, err := s3api.ListCertificate("2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCGX")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range infos {
		fmt.Println(v.State, v.CertificateID, v.SecretKey)
	}
}
func TestDeleteAccount(t *testing.T) {

	ok, err := s3api.DeleteAccount("2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCGX")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("删除S3账户结果", ok)

}

func TestRunS3(t *testing.T) {

	accountName := "zhangsanfeng"
	if accountID, err := s3api.CreateAccount(accountName, 0); err == nil {
		fmt.Println("创建账户结果", accountID)
		if ak, sk, err := s3api.CreateCertificate(accountID); err == nil {
			fmt.Println("创建证书成功", ak, sk)
			if ok, err := s3api.DeleteCertificate(ak); err == nil {
				fmt.Println("删除证书", ok)
				if ok, err := s3api.DeleteAccount(accountID); err == nil {
					fmt.Println("删除账户", ok)
				}
			}
		}

	} else {
		t.Fatal("创建账户失败", err.Error())
	}

}
