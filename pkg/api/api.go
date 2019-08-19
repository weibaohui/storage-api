package api

type NFSApi interface {
	CreateDirectory(path string) (ok bool, err error)
	DeleteDirectory(path string) (ok bool, err error)

	CreateQuota(path string, ips, ops, readBw, writeBw int) (ok bool, err error)
	DeleteQuota(id string) (ok bool, err error)
}
