package s3

import (
	"encoding/json"
	"fmt"
	"nfs-api/pkg/api"
	"nfs-api/pkg/sg"
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
//POST
//https://192.168.3.60:6080/commands/add_account.action?cmd_id=0.8293892534223559&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//rand:
//params: {"account_name":"mingcheng","account_email":"email@email.com","account_quota":0}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"account_id":"2PO1B8F4QBKHSRZW2SA3A9G2BFJ2EKB6","account_name":"mingcheng","create_date":"2019-08-19 16:53:08"},"sync":true,"time_stamp":1566204788677,"time_zone_offset":-480,"trace_id":"[348220645984 2]"}
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
//POST
//https://192.168.3.60:6080/commands/add_certificate.action?cmd_id=0.11054967319896059&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
// rand:
//params: {"account_id":"2PO1B8F4QBKHSRZW2SA3A9G2BFJ2EKB6"}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"account_id":"2PO1B8F4QBKHSRZW2SA3A9G2BFJ2EKB6","account_name":"mingcheng","create_date":"2019-08-19 16:53:08"},"sync":true,"time_stamp":1566204788677,"time_zone_offset":-480,"trace_id":"[348220645984 2]"}
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
//https://192.168.3.60:6080/commands/list_accounts_sort.action?cmd_id=0.272262162377473&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//params: {"limit":0,"start":0,"sort":"","start_account_name":".","number":20}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"accounts":[{"account_email":"4545@11.com","account_id":"2PO1B8F4QBKHSRZW2C21IZG6LI3FUMQ9","account_name":"555","account_quota":0,"bucket_number":0,"create_date":"2019-08-19 14:42:32","used_bytes":0},{"account_email":"email@email.com","account_id":"2PO1B8F4QBKHSRZW2SA3A9G2BFJ2EKB6","account_name":"mingcheng","account_quota":0,"bucket_number":0,"create_date":"2019-08-19 16:53:08","used_bytes":0},{"account_email":"sugontest@sugon.com","account_id":"2PO1B8F4QBKHSRZW2C21IZG6LI3FUMQB","account_name":"sugontest","account_quota":0,"bucket_number":0,"create_date":"2019-08-15 16:44:01","used_bytes":0}],"has_more_data":false,"total":3},"sync":true,"time_stamp":1566204788709,"time_zone_offset":-480,"trace_id":"[348220691257 2]"}
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
//https://192.168.3.60:6080/commands/list_certificate.action?cmd_id=0.6751997843285937&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//params: {"limit":20,"start":0,"sort":"","account_id":"2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCHQ"}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":{"account_id":"2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCHQ","certificate_info":[{"certificate_id":"2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCHQ2EL1TO9IPQEEFKML2SA3A9G2BFJ2EKBA","create_date":"2019-08-19 17:39:51","secret_key":"b4824f981d81a67be8d9f34e89acc60c1f27a7b9","state":"S3_CERTIFICATE_ENABLE"}],"certificate_nr":1},"sync":true,"time_stamp":1566207598651,"time_zone_offset":-480,"trace_id":"[351030643522 2]"}
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
//POST
//https://192.168.3.60:6080/commands/delete_certificate.action?cmd_id=0.01966998131719655&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//params: {"certificate_id":"2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCHQ2EL1TO9IPQEEFKML2SA3A9G2BFJ2EKBA"}
//{"detail_err_msg":"","err_msg":"","err_no":0,"sync":true,"time_stamp":1566208239975,"time_zone_offset":-480,"trace_id":"[351671949124 2]"}
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
//https://192.168.3.60:6080/commands/delete_account.action?cmd_id=0.42864007450337804&user_name=optadmin&uuid=9fdc9c55-cb34-4e40-9da9-ada6d5334a6c
//params: {"account_id":"2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCHQ"}
//{"detail_err_msg":"","err_msg":"","err_no":0,"result":"2PO1B8F4QBKHSRZW31W1ZLS4U5AIGCHQ","sync":true,"time_stamp":1566208473514,"time_zone_offset":-480,"trace_id":"[351905484684 2]"}
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
