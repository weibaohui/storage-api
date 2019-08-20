package nfs

import (
	"storage-api/pkg/api"
	"storage-api/pkg/sg/common"
)

type instance struct {
	common *common.Instance
}

func NewInstance(config *api.Config) api.NFSApi {
	instance := &instance{
		common: common.NewInstance(config),
	}
	return instance
}
