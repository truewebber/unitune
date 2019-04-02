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
	socksPrx, err := proxy.SOCKS5("tcp", prx.String(), nil, proxy.Direct)
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
