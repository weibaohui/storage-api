package nfs

import (
	"nfs-api/pkg/api"
	"nfs-api/pkg/sg/common"
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
