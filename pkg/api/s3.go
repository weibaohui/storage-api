package api

type S3Api interface {
	// 创建S3账户
	CreateAccount(name string, quota int) (accountID string, err error)
	// 删除S3账户
	DeleteAccount(accountID string) (ok bool, err error)
	// 创建S3账户对应的访问证书
	CreateCertificate(accountID string) (ak, sk string, err error)
	// 删除S3账户对应的访问证书
	DeleteCertificate(ak string) (ok bool, err error)
}
