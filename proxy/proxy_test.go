package proxy

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewHttp(t *testing.T) {
	prx := &Proxy{
		Ip:   "37.235.238.93",
		Port: 8080,
		Type: HttpType,
	}

	p := NewHttp(prx)

	resp, err := p.HttpClient().Get("https://wtfismyip.com/json")
	if err != nil {
		t.Error(err.Error())

		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("Error", "Bad status code", "_", resp.StatusCode)

		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(string(body))
}

func TestNewSocks5(t *testing.T) {
	prx := &Proxy{
		Ip:   "184.174.73.158",
		Port: 51724,
		Type: Socks5Type,
	}

	p := NewSocks5(prx)

	resp, err := p.HttpClient().Get("https://wtfismyip.com/json")
	if err != nil {
		t.Error(err.Error())

		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("Error", "Bad status code", "_", resp.StatusCode)

		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(string(body))
}
