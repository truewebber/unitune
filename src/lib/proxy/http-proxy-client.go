package proxy

import (
	"net"
	"net/http"
	"net/url"
)

type (
	HttpProxyClient interface {
		HttpClient() *http.Client
		Type() Type
	}

	SocksProxy interface {
		HttpClient() *http.Client
		Dial(string, string) (net.Conn, error)
	}

	HttpProxy interface {
		HttpClient() *http.Client
		Proxy(*http.Request) (*url.URL, error)
	}
)
