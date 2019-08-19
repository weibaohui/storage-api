package sg

import (
	"fmt"
	"testing"
)

func TestListS3Accounts(t *testing.T) {
	robot := FakeRobot4Test()
	accounts, err := robot.ListAccount()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range accounts {
		fmt.Println(v.AccountID, v.AccountName, v.AccountQuota, v.UsedBytes)
	}
}
func TestCreateS3Accounts(t *testing.T) {
	robot := FakeRobot4Test()

	accountID, err := robot.CreateAccount("testttt", 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("创建S3账户ID=", accountID)
}
func TestCreateCertificate(t *testing.T) {
	robot := FakeRobot4Test()

	ak, sk, err := robot.CreateCertificate("2PO1B8F4QBKHSRZW1L10GHBYOJTJ8MDK")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("创建S3账户", ak, sk)
}
func TestListCertificate(t *testing.T) {
	robot := FakeRobot4Test()

	infos, err := robot.ListCertificate("2PO1B8F4QBKHSRZW1L10GHBYOJTJ8MDK")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range infos {
		fmt.Println(v.State, v.CertificateID, v.SecretKey)
	}
}
func TestDeleteAccount(t *testing.T) {
	robot := FakeRobot4Test()

	ok, err := robot.DeleteAccount("2PO1B8F4QBKHSRZW1L10GHBYOJTJ8MDK")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("删除S3账户结果", ok)

}
