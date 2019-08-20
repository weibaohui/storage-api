package s3

import (
	"nfs-api/pkg/api"
	"nfs-api/pkg/sg/common"
)

type instance struct {
	common *common.Instance
}

func NewInstance(config *api.Config) api.S3Api {
	instance := &instance{
		common: common.NewInstance(config),
	}
	return instance
}
