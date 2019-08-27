package common

import (
	"encoding/json"
	"storage-api/pkg/sg"
)

type ClusterStatusResult struct {
	sg.ErrorMsg
	Data *Cluster `json:"result"`
}

//集群状态
func (i *Instance) ClusterStatus() (status *Cluster, err error) {
	url := i.Command("/commands/get_cluster_overview.action")
	str, err := i.PostWithLoginSession(url, nil)
	result := &ClusterStatusResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	return result.Data, nil
}
