package api

type NFSApi interface {
	//创建目录
	CreateDirectory(path string) (ok bool, err error)
	//删除目录
	DeleteDirectory(path string) (ok bool, err error)
	//创建配额
	CreateQuota(path string, ips, ops, readBw, writeBw int) (ok bool, quotaID string, err error)
	//删除配额
	DeleteQuota(id string) (ok bool, err error)
}
