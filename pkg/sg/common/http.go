package common

import (
	"crypto/tls"
	"errors"
	"github.com/weibaohui/go-kit/httpkit"
	"net/http"
	"strings"
)

func (r *Instance) PostWithLoginSession(fullURL string, params map[string]string) (str string, err error) {
	req := httpkit.Post(fullURL)
	SetSkipSSLVerify(req)

	for _, v := range r.loginCookies {
		req.SetCookie(v)
	}
	for k, v := range params {
		req.Param(k, v)
	}

	req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	str, err = req.String()
	if err != nil {
		return "", err
	}
	if strings.Contains(str, "<script>window.top.location='/'</script>") {
		if r.retryLogin <= 3 {
			//自动进行重新登录
			r.retryLogin += 1
			r.connect()
			return r.PostWithLoginSession(fullURL, params)
		}
		return "", errors.New("cookie失效，请检查登录参数")
	}

	return str, err
}

func SetSkipSSLVerify(req *httpkit.HTTPRequest) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	req.SetTransport(tr)
}
