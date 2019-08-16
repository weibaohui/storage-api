package sg

import (
	"fmt"
)

type Robot struct {
	Protocol string //协议
	Host     string //主机
	Port     string //端口
	Username string //管理用户名
	Password string //管理密码
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
