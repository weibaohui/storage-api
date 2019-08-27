package common

import (
	"encoding/json"
	"errors"
	"storage-api/pkg/sg"
)

type AccessZone struct {
	ID              int    `json:"id"`
	Key             int    `json:"key"`
	AccessZoneState string `json:"access_zone_state"`
	AuthProvider    struct {
		ID      int    `json:"id"`
		Key     int    `json:"key"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		Version int    `json:"version"`
	} `json:"auth_provider"`
	AuthProviderID      int    `json:"auth_provider_id"`
	EnableFc            bool   `json:"enable_fc"`
	EnableFtp           bool   `json:"enable_ftp"`
	EnableHTTP          bool   `json:"enable_http"`
	EnableIscsi         bool   `json:"enable_iscsi"`
	EnableNfs           bool   `json:"enable_nfs"`
	EnableS3            bool   `json:"enable_s3"`
	EnableSan           bool   `json:"enable_san"`
	EnableSmb           bool   `json:"enable_smb"`
	LocalAuthProviderID int    `json:"local_auth_provider_id"`
	MaxNodeNumber       int    `json:"max_node_number"`
	Name                string `json:"name"`
	NasServiceEnabled   bool   `json:"nas_service_enabled"`
	NodeIds             []int  `json:"node_ids"`
	Nodes               []struct {
		AutoUmount   bool   `json:"auto_umount"`
		ManagerNode  bool   `json:"managerNode"`
		NodeID       int    `json:"node_id"`
		NodeName     string `json:"node_name"`
		ZkServerNode bool   `json:"zkServerNode"`
	} `json:"nodes"`
	S3ServiceEnabled  bool   `json:"s3_service_enabled"`
	SanProtocolState  string `json:"san_protocol_state"`
	SanServiceEnabled bool   `json:"san_service_enabled"`
	Version           int    `json:"version"`
}
type AccessZoneQuery struct {
	AccessZones []*AccessZone `json:"access_zones"`
	Limit       int           `json:"limit"`
	Searches    []interface{} `json:"searches"`
	Sort        string        `json:"sort"`
	Start       int           `json:"start"`
	Total       int           `json:"total"`
}
type AccessZoneResult struct {
	sg.ErrorMsg
	Data *AccessZoneQuery `json:"result"`
}

//列表显示当前的访问取
func (i *Instance) ListAccessZones() ([]*AccessZone, error) {
	url := i.Command("/commands/get_access_zones.action")
	params := make(map[string]string)
	params["params"] = `{"limit":99999,"start":0,"sort":""}`
	str, err := i.PostWithLoginSession(url, params)
	if err != nil {
		return nil, err
	}
	result := &AccessZoneResult{}
	err = json.Unmarshal([]byte(str), result)
	if err != nil {
		return nil, err
	}
	if result.ErrNo != 0 {
		return nil, result.Error()
	}
	if len(result.Data.AccessZones) > 0 {
		return result.Data.AccessZones, nil
	}
	return nil, errors.New("没有访问区")
}
