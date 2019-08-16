package sg

import (
	"errors"
	"github.com/weibaohui/go-kit/httpkit"
	"net/http"
)

//操作使用的cookie
func (r *Robot) AuthCookie() (cookies []*http.Cookie, err error) {
	return r.mainPageCookie()
}

// 登录cookie
// POST
// strUserName: optadmin
// strPassword: adminadmin
// language: zh_CN
func (r *Robot) loginCookie() ([]*http.Cookie, error) {
	url := r.fullURL(LoginUrl)
	post := httpkit.Post(url)
	post.Param("strUserName", r.Username)
	post.Param("strPassword", r.Password)
	post.Param("language", "zh_CN")
	post.EnableCookie()

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

//主界面的Cookie
func (r *Robot) mainPageCookie() (cookies []*http.Cookie, err error) {

	url := r.fullURL(MainPageUrl)
	req := httpkit.Get(url)
	SetSkipSSLVerify(req)
	for _, v := range r.loginCookies {
		cookies = append(cookies, v)
		req.SetCookie(v)
	}
	getResp, err := req.Response()
	if err != nil {
		return nil, err
	}
	mainCookies := getResp.Cookies()
	if mainCookies == nil {
		return nil, errors.New("服务器没有返回主页面cookies")
	}

	for _, v := range mainCookies {
		cookies = append(cookies, v)
	}
	return cookies, nil
}
