package common

import (
	"errors"
	"github.com/weibaohui/go-kit/httpkit"
	"net/http"
)

//操作使用的cookie
func (i *Instance) AuthCookie() (cookies []*http.Cookie, err error) {
	return i.loginCookie()
}

// 登录cookie
// POST
// strUserName: optadmin
// strPassword: adminadmin
// language: zh_CN
func (i *Instance) loginCookie() ([]*http.Cookie, error) {
	url := i.FullURL("/login/loginAuth.action")
	post := httpkit.Post(url)
	post.Param("strUserName", i.Username)
	post.Param("strPassword", i.Password)
	post.Param("language", "zh_CN")

	SetSkipSSLVerify(post)
	response, err := post.Response()
	if err != nil {
		return nil, err
	}
	cookies := response.Cookies()
	if i == nil {
		return nil, errors.New("服务器没有返回登录cookies")
	}
	return cookies, nil

}
