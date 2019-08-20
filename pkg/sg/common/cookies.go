package common

import (
	"errors"
	"github.com/weibaohui/go-kit/httpkit"
	"net/http"
)

//操作使用的cookie
func (r *Instance) AuthCookie() (cookies []*http.Cookie, err error) {
	return r.loginCookie()
}

// 登录cookie
// POST
// strUserName: optadmin
// strPassword: adminadmin
// language: zh_CN
func (r *Instance) loginCookie() ([]*http.Cookie, error) {
	url := r.FullURL("/login/loginAuth.action")
	post := httpkit.Post(url)
	post.Param("strUserName", r.Username)
	post.Param("strPassword", r.Password)
	post.Param("language", "zh_CN")

	SetSkipSSLVerify(post)
	response, err := post.Response()
	if err != nil {
		return nil, err
	}
	i := response.Cookies()
	if i == nil {
		return nil, errors.New("服务器没有返回登录cookies")
	}
	return i, nil

}
