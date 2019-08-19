package api

type NFSApi interface {
	CreateDirectory(path string) (bool, error)
	DeleteDirectory(path string) (bool, error)

	CreateQuota(dir string, ips, ops, readBw, writeBw int) (bool, error)
	DeleteQuota(id string) (bool, error)
}
