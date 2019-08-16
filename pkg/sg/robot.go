package sg

import (
	"fmt"
	"log"
	"net/http"
)

type Robot struct {
	Protocol     string         //协议
	Host         string         //主机
	Port         string         //端口
	Username     string         //管理用户名
	Password     string         //管理密码
	uuid         string         //磁阵系统UUID
	storeName    string         //存储名称
	loginCookies []*http.Cookie //登录cookie
}

func FakeRobot() *Robot {
	newRobot := NewRobot("https", "192.168.3.60", "6080", "optadmin", "adminadmin")
	robot := newRobot.Connect()
	return robot
}

func NewRobot(protocol, host, port, username, password string) *Robot {
	return &Robot{
		Protocol: protocol,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}
func (r *Robot) fullURL(path string) (fullURL string) {
	return fmt.Sprintf("%s://%s:%s%s", r.Protocol, r.Host, r.Port, path)
}
func (r *Robot) Connect() *Robot {
	cookies, err := r.loginCookie()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.loginCookies = cookies

	store, err := r.DefaultStore()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.storeName = store.Name
	r.uuid = store.UUID
	return r
}
