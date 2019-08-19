package api

type S3Api interface {
	CreateAccount(name string, quota int) (accountID string, err error)
	CreateCertificate(accountID string) (ak, sk string, err error)
}
