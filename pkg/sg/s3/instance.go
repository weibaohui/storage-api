package s3

import (
	"storage-api/pkg/api"
	"storage-api/pkg/sg/common"
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
