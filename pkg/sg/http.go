package sg

import (
	"crypto/tls"
	"errors"
	"github.com/weibaohui/go-kit/httpkit"
	"net/http"
	"strings"
)

func (r *Robot) PostWithLoginSession(fullURL string, params map[string]string) (string, error) {
	req := httpkit.Post(fullURL)
	SetSkipSSLVerify(req)
	loginCookies, err := r.loginCookie()
	if err != nil {
		return "", err
	}
	for _, v := range loginCookies {
		req.SetCookie(v)
	}
	for k, v := range params {
		req.Param(k, v)
	}

	str, err := req.String()
	if err != nil {
		return "", err
	}
	if strings.Contains(str, "<script>window.top.location='/'</script>") {
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
