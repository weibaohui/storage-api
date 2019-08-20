package s3

import (
	"fmt"
	"nfs-api/pkg/api"
	"nfs-api/pkg/sg/common"
)

type instance struct {
	*api.Config
	common *common.Robot
}

func NewInstance(config *api.Config) api.S3Api {
	instance := &instance{
		Config: config,
		common: common.NewInstance(config),
	}
	return instance
}
func (r *instance) fullURL(path string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s", r.Protocol, r.Host, r.Port, path)
}
