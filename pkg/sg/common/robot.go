package common

import (
	"fmt"
	"log"
	"net/http"
	"nfs-api/pkg/api"
)

type Robot struct {
	*api.Config
	UUID         string         //磁阵系统UUID
	StoreName    string         //存储系统名称
	loginCookies []*http.Cookie //登录cookie
}

func NewInstance(config *api.Config) *Robot {
	robot := &Robot{
		Config:       config,
		UUID:         "",
		StoreName:    "",
		loginCookies: nil,
	}
	robot.connect()
	return robot
}

func (r *Robot) connect() {
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

func (r *Robot) fullURL(path string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s", r.Protocol, r.Host, r.Port, path)
}