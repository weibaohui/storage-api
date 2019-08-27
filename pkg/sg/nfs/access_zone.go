package nfs

import (
	"errors"
	"fmt"
)

//列表显示目录
//列表显示NFS可以使用的AccessZone
func (i *instance) ListNFSAccessZone() (string, error) {
	zones, err := i.common.ListAccessZones()
	if err != nil {
		return "", err
	}
	for _, v := range zones {
		if v.EnableNfs {
			return fmt.Sprintf("%d", v.ID), nil
		}
	}

	return "", errors.New("没有找到NFS可以使用的分区")
}
