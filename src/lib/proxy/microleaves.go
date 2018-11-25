package proxy

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"lib/config"
)

type (
	MicroLeaves struct {
		host string
		Geo  string
	}
)

const (
	MicroLeavesModuleName = "microleaves"
)

var (
	microleavesOnce sync.Once
	microLeavesMap  = make(map[string][]int)
)

func MicroLeavesInitialization(conf *config.Config) {
	conf.UnmarshalKey("proxy.microleaves.ports", &microLeavesMap)
}

func NewMicroLeaves(conf *config.Config, geo string) *MicroLeaves {
	microleavesOnce.Do(func() {
		MicroLeavesInitialization(conf)
	})

	return &MicroLeaves{
		host: conf.GetString("proxy.microleaves.host"),
		Geo:  geo,
	}
}

func (l *MicroLeaves) GetModuleName() string {
	return MicroLeavesModuleName
}

func (l *MicroLeaves) GetHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return l.Proxy(req)
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			MaxIdleConnsPerHost:   -1,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 60 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 1200 * time.Second,
	}
}

func (l *MicroLeaves) getProxy() (*Proxy, error) {
	ports, ok := microLeavesMap[strings.ToLower(l.Geo)]
	if !ok {
		return nil, errors.Errorf("(MicroLeaves) Unsupported geo: %s", l.Geo)
	}

	rand.Seed(time.Now().UnixNano())

	return &Proxy{
		Geo:  l.Geo,
		Ip:   l.host,
		Port: ports[rand.Intn(len(ports))],
	}, nil
}

func (l *MicroLeaves) Proxy(req *http.Request) (*url.URL, error) {
	prx, err := l.getProxy()
	if err != nil {
		return nil, err
	}

	if req == nil {
		return prx.Url()
	}

	return prx.Url()
}
