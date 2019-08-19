package api

type CertificateInfo struct {
	CertificateID string `json:"certificate_id"`
	CreateDate    string `json:"create_date"`
	SecretKey     string `json:"secret_key"`
	State         string `json:"state"`
}
type S3Api interface {
	// 创建S3账户
	CreateAccount(name string, quota int) (accountID string, err error)
	// 删除S3账户
	DeleteAccount(accountID string) (ok bool, err error)
	// 创建S3账户对应的访问证书
	CreateCertificate(accountID string) (ak, sk string, err error)
	// 删除S3账户对应的访问证书
	DeleteCertificate(ak string) (ok bool, err error)
	//列表账户下的证书列表
	ListCertificate(accountID string) ([]*CertificateInfo, error)
}
