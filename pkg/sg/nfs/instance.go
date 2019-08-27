package nfs

import (
	"errors"
	"storage-api/pkg/api"
	"storage-api/pkg/sg/common"
)

type instance struct {
	common *common.Instance
}

func NewInstance(config *api.Config) (api.NFSApi, error) {
	if config.StoragePoolName == "" {
		return nil, errors.New("缺少StoragePoolName，请设置")
	}
	instance := &instance{
		common: common.NewInstance(config),
	}
	return instance, nil
}
