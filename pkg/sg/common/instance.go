package common

import (
	"fmt"
	"log"
	"net/http"
	"storage-api/pkg/api"
)

type Instance struct {
	*api.Config
	ClusterUUID  string         //磁阵系统UUID
	loginCookies []*http.Cookie //登录cookie
	retryLogin   int            //当登录失败时尝试重新登录次数
}

func NewInstance(config *api.Config) *Instance {
	instance := &Instance{
		Config:       config,
		ClusterUUID:  "",
		loginCookies: nil,
	}
	instance.connect()
	return instance
}

func (i *Instance) connect() {
	cookies, err := i.loginCookie()
	if err != nil {
		log.Fatal(err.Error())
	}
	i.loginCookies = cookies

	store, err := i.DefaultCluster()
	if err != nil {
		log.Fatal(err.Error())
	}
	i.StoragePoolName = i.Config.StoragePoolName
	i.ClusterUUID = store.UUID
}

func (i *Instance) FullURL(path string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s", i.Protocol, i.Host, i.Port, path)
}

func (i *Instance) Command(command string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s?user_name=%s&uuid=%s", i.Protocol, i.Host, i.Port, command, i.Username, i.ClusterUUID)
}
