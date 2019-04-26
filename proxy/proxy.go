package proxy

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

type (
	Proxy struct {
		Ip   string `mapstructure:"ip"`
		Port int    `mapstructure:"port"`
		Auth *Auth  `mapstructure:"auth"`
		Type Type   `mapstructure:"type"`
	}

	Auth struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	}
)

func (p *Proxy) Url() (*url.URL, error) {
	return url.Parse(p.String())
}

func (p *Proxy) LogPass() string {
	return fmt.Sprintf("%s:%s", p.Auth.User, p.Auth.Password)
}

func (p *Proxy) BasicAuth() string {
	if p.Auth == nil {
		return ""
	}

	return "Basic " + base64.StdEncoding.EncodeToString([]byte(p.LogPass()))
}

func (p *Proxy) String() string {
	if p.Auth != nil {
		return fmt.Sprintf("%s://%s:%s@%s:%d", p.Type, p.Auth.User, p.Auth.Password, p.Ip, p.Port)
	}
	return fmt.Sprintf("%s://%s:%d", p.Type, p.Ip, p.Port)
}
