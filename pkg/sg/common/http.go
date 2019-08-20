package common

import (
	"crypto/tls"
	"errors"
	"github.com/weibaohui/go-kit/httpkit"
	"log"
	"net/http"
	"strings"
)

func (i *Instance) PostWithLoginSession(fullURL string, params map[string]string) (str string, err error) {
	req := httpkit.Post(fullURL)
	SetSkipSSLVerify(req)

	for _, v := range i.loginCookies {
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
		if i.retryLogin <= 3 {
			log.Println("进行重新登录", i.retryLogin)
			//自动进行重新登录
			i.retryLogin += 1
			i.connect()
			return i.PostWithLoginSession(fullURL, params)
		}

		return "", errors.New("cookie失效，请检查登录参数")
	}
	//重置重新登录次数
	if i.retryLogin > 0 {
		i.retryLogin = 0
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
