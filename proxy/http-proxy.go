package proxy

import (
	"net/http"
	"net/url"
)

type (
	Http struct {
		proxy *Proxy
	}
)

func NewHttp(prx *Proxy) *Http {
	return &Http{
		proxy: prx,
	}
}

func (h *Http) HttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: h.Proxy,
		},
	}
}

func (h *Http) Proxy(req *http.Request) (*url.URL, error) {
	if req == nil {
		return h.proxy.Url()
	}

	if h.proxy.Auth != nil {
		req.Header.Add("Proxy-Authorization", h.proxy.BasicAuth())
	}

	return h.proxy.Url()
}

func (h *Http) Type() Type {
	return HttpType
}
