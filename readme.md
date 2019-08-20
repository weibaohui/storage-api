- 存储操作的API
- 适用于ParaStor300S
------
```
   type NFSApi interface {
	//创建目录
	CreateDirectory(path string) (ok bool, err error)
	//删除目录
	DeleteDirectory(path string) (ok bool, err error)
	//创建配额
	CreateQuota(path string, ips, ops, readBw, writeBw int) (ok bool, quotaID string, err error)
	//删除配额
	DeleteQuota(id string) (ok bool, err error)

	//列表显示文件夹及文件
	ListDirectoryWithFiles(path string) ([]*DetailFiles, error)
	//列表显示文件夹
	ListDirectory(path string) ([]*DetailFiles, error)

	//列表快照
	ListSnapshot() ([]*Snapshot, error)
	//创建快照
	CreateSnapshot(name, path, desc string, expireTime int) (id string, err error)
	//快照回滚
	RevertSnapshot(id string) (ok bool, err error)
	//删除快照
	DeleteSnapshot(id string) (ok bool, err error)
}

```


```
type S3Api interface {
	//列表S3账户
	ListAccount() ([]*Account, error)
	// 创建S3账户
	CreateAccount(name string, quota int) (accountID string, err error)
	// 删除S3账户
	DeleteAccount(accountID string) (ok bool, err error)

	//列表账户下的证书列表
	ListCertificate(accountID string) ([]*CertificateInfo, error)
	// 创建S3账户对应的访问证书
	CreateCertificate(accountID string) (ak, sk string, err error)
	// 删除S3账户对应的访问证书
	DeleteCertificate(ak string) (ok bool, err error)
}
```