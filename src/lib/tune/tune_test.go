package tune

import (
	"testing"

	"lib/proxy"
)

func TestGetLinkInfo(t *testing.T) {
	//info, err := GetLinkInfo("https://music.yandex.ru/album/7014466/track/50664999") //single
	//info, err := GetLinkInfo("https://music.yandex.ru/album/1610161/track/14698882")

	//info, err := GetLinkInfo("https://open.spotify.com/track/7rc6L3UtM0uvwfpsl07GBL?si=KlfgJCZaSTS2GxdFViMUjA") //single
	//info, err := GetLinkInfo("https://open.spotify.com/track/5GjnIpUlLGEIYk052ISOw9?si=6gx17GPCTSCHsuvDX3OVLA") //albom

	// get proxyList; add default non-proxy http client
	proxyList := proxy.GenerateProxyList([]*proxy.Proxy{})
	//proxyList = append([]proxy.HttpProxyClient{proxy.NewNull()}, proxyList...)

	// init tunner
	tunner := NewTunner(proxyList)

	info, err := tunner.Tune("https://music.yandex.ru/album/7014466/track/50664999")

	if err != nil {
		t.Error(err.Error())

		return
	}

	t.Log(info)
}
