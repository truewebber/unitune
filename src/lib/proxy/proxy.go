package proxy

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"net/url"
)

type (
	HttpProxyClient interface {
		GetHttpClient() *http.Client
		GetModuleName() string
		SetHostChecker(hosts []string)
	}

	SocksProxyFactory interface {
		GetHttpClient() *http.Client
		Dial(string, string) (net.Conn, error)
	}

	HttpProxyFactory interface {
		Proxy(*http.Request) (*url.URL, error)
	}

	Proxy struct {
		Geo      string
		Ip       string
		Port     int
		User     string
		Password string
	}
)

func (p *Proxy) Url() (*url.URL, error) {
	return url.Parse(p.String())
}

func (p *Proxy) LogPass() string {
	return fmt.Sprintf("%s:%s", p.User, p.Password)
}

func (p *Proxy) BasicAuth() string {
	if len(p.User) == 0 {
		return ""
	}
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(p.LogPass()))
}

func (p *Proxy) String() string {
	if len(p.User) > 0 {
		return fmt.Sprintf("http://%s:%s@%s:%d", p.User, p.Password, p.Ip, p.Port)
	}
	return fmt.Sprintf("http://%s:%d", p.Ip, p.Port)
}
