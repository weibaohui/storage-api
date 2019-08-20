package common

import (
	"fmt"
	"log"
	"net/http"
	"nfs-api/pkg/api"
)

type Instance struct {
	*api.Config
	UUID         string         //磁阵系统UUID
	StoreName    string         //存储系统名称
	loginCookies []*http.Cookie //登录cookie
	retryLogin   int            //当登录失败时尝试重新登录次数
}

func NewInstance(config *api.Config) *Instance {
	instance := &Instance{
		Config:       config,
		UUID:         "",
		StoreName:    "",
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

	store, err := i.DefaultStore()
	if err != nil {
		log.Fatal(err.Error())
	}
	i.StoreName = store.Name
	i.UUID = store.UUID
}

func (i *Instance) FullURL(path string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s", i.Protocol, i.Host, i.Port, path)
}

func (i *Instance) Command(command string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s?user_name=%s&uuid=%s", i.Protocol, i.Host, i.Port, command, i.Username, i.UUID)
}
