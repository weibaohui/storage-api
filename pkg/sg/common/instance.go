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

func (r *Instance) connect() {
	cookies, err := r.loginCookie()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.loginCookies = cookies

	store, err := r.DefaultStore()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.StoreName = store.Name
	r.UUID = store.UUID
}

func (r *Instance) FullURL(path string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s", r.Protocol, r.Host, r.Port, path)
}

func (r *Instance) Command(command string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s?user_name=%s&uuid=%s", r.Protocol, r.Host, r.Port, command, r.Username, r.UUID)
}
