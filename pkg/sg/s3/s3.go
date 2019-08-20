package s3

import (
	"encoding/json"
	"fmt"
	"storage-api/pkg/api"
	"storage-api/pkg/sg"
)

type CertificateList struct {
	AccountID       string                 `json:"account_id"`
	CertificateInfo []*api.CertificateInfo `json:"certificate_info"`
	CertificateNr   int                    `json:"certificate_nr"`
}
type CertificateListResult struct {
	sg.ErrorMsg
	Data *CertificateList `json:"result"`
}
type s3AccountList struct {
	Accounts    []*api.Account `json:"accounts"`
	HasMoreData bool           `json:"has_more_data"`
	Total       int            `json:"total"`
}
type s3AccountListResult struct {
	sg.ErrorMsg
	Data s3AccountList `json:"result"`
}

type s3AccountResult struct {
	sg.ErrorMsg
	S3Account *api.Account `json:"result"`
}

type s3Certificate struct {
	AccountID     string `json:"account_id"`
	CertificateID string `json:"certificate_id"`
	CreateDate    string `json:"create_date"`
	SecretKey     string `json:"secret_key"`
	State         string `json:"state"`
}
type s3CertificateResult struct {
	sg.ErrorMsg
	Certificate *s3Certificate `json:"result"`
}

//创建S3账户
//quota 单位GB
func (i *instance) CreateAccount(name string, quota int) (accountID string, err error) {
	url := i.common.Command("/commands/add_account.action")
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{
	"account_name":"%s",
	"account_email":"%s@email.com",
	"account_quota":%d
	}
	`, name, name, quota)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return "", err
	}
	result := &s3AccountResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return "", err
	}
	if result.ErrNo != 0 {
		return "", result.Error()
	}
	return result.S3Account.AccountID, nil
}

//创建S3账户
func (i *instance) CreateCertificate(accountID string) (ak, sk string, err error) {
	url := i.common.Command("/commands/add_certificate.action")
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{"account_id":"%s"}`, accountID)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return "", "", err
	}
	result := &s3CertificateResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return "", "", err
	}
	if result.ErrNo != 0 {
		return "", "", result.Error()
	}
	return result.Certificate.CertificateID, result.Certificate.SecretKey, nil
}

//POST 查询账户
//todo 分页最大1000
func (i *instance) ListAccount() ([]*api.Account, error) {
	url := i.common.Command("/commands/list_accounts_sort.action")
	params := make(map[string]string)
	params["params"] = `{
	"limit":0,"start":0,
	"sort":"",
	"start_account_name":".",
	"number":1000
	}
	`
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	result := &s3AccountListResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result.Data.Accounts, nil
}

//查询S3账户接入证书AK\SK
//POST
func (i *instance) ListCertificate(accountID string) ([]*api.CertificateInfo, error) {
	url := i.common.Command("/commands/list_certificate.action")
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{
	"limit":1000,"start":0,"sort":"",
	"account_id":"%s"
	}
	`, accountID)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	result := &CertificateListResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result.Data.CertificateInfo, nil
}

//删除证书
func (i *instance) DeleteCertificate(certificateID string) (ok bool, err error) {
	url := i.common.Command("/commands/delete_certificate.action")
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{
	"certificate_id":"%s"
	}
	`, certificateID)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &sg.ErrorMsg{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}
	return true, nil
}

//删除S3账户
//POST
func (i *instance) DeleteAccount(accountID string) (ok bool, err error) {
	//1 检查是否有证书，有的话全删除
	//2 删除s3账户
	infos, err := i.ListCertificate(accountID)
	if err != nil {
		return false, err
	}
	if len(infos) > 0 {
		for _, v := range infos {
			i.DeleteCertificate(v.CertificateID)
		}
	}

	//2 删除账户
	url := i.common.Command("/commands/delete_account.action")
	params := make(map[string]string)
	params["params"] = fmt.Sprintf(`{
	"account_id":"%s"
	}
	`, accountID)
	str, err := i.common.PostWithLoginSession(url, params)
	if err != nil {
		return false, err
	}
	result := &sg.ErrorMsg{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return false, err
	}
	if result.ErrNo != 0 {
		return false, result.Error()
	}
	return true, nil
}
