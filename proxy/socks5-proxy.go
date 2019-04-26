package proxy

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/proxy"
)

type (
	Socks5 struct {
		dialer proxy.Dialer
	}
)

func NewSocks5(prx *Proxy) *Socks5 {
	var auth *proxy.Auth
	if prx.Auth != nil {
		auth = &proxy.Auth{
			User:     prx.Auth.User,
			Password: prx.Auth.Password,
		}
	}
	socksPrx, err := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%d", prx.Ip, prx.Port), auth, proxy.Direct)
	if err != nil {
		fmt.Println("Error connecting to proxy:", err)
	}

	return &Socks5{
		dialer: socksPrx,
	}
}

func (s *Socks5) HttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, addr string) (conn net.Conn, e error) {
				return s.dialer.Dial(network, addr)
			},
		},
	}
}

func (s *Socks5) Dial(network string, addr string) (net.Conn, error) {
	return s.dialer.Dial(network, addr)
}

func (s *Socks5) Type() Type {
	return Socks5Type
}
