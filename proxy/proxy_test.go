package proxy

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"golang.org/x/net/proxy"
)

func TestProxy_String(t *testing.T) {
	prx := Proxy{
		Ip:   "46.16.13.212",
		Port: 3001,
		Type: Socks5Type,
	}

	socksPrx, err := proxy.SOCKS5("tcp", prx.String(), nil, proxy.Direct)
	if err != nil {
		fmt.Println("Error connecting to proxy:", err)
	}

	c := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, addr string) (conn net.Conn, e error) {
				return socksPrx.Dial(network, addr)
			},
		},
	}

	resp, err := c.Get("https://wtfismyip.com/plain")
	if err != nil {
		t.Fatal(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(fmt.Sprintf("%d", resp.StatusCode))
	}

	t.Log("OK")
}
