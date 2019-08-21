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
type Certificate struct {
	AccountID     string `json:"account_id"`
	CertificateID string `json:"certificate_id"`
	CreateDate    string `json:"create_date"`
	SecretKey     string `json:"secret_key"`
	State         string `json:"state"`
}
type CertificateResult struct {
	sg.ErrorMsg
	Certificate *Certificate `json:"result"`
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
	result := &CertificateResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return "", "", err
	}
	if result.ErrNo != 0 {
		return "", "", result.Error()
	}
	return result.Certificate.CertificateID, result.Certificate.SecretKey, nil
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
