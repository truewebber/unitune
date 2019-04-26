package config

import (
	"encoding/json"
	"testing"

	"github.com/truewebber/unitune/proxy"
)

func TestNewConfig(t *testing.T) {
	cfgPath := "./mock/"
	cfg := NewConfig(cfgPath, "proxy")

	obj := make([]*proxy.Proxy, 0)
	err := cfg.UnmarshalKey("proxies", &obj)
	if err != nil {
		t.Fatal("Error unmarshal config:", err.Error())
	}

	result, _ := json.MarshalIndent(obj, "", "\t")
	t.Log(string(result))
}
