package seeker

import (
	"testing"

	"lib/proxy"
	"lib/tune"
)

func TestYandexMusic_Seek(t *testing.T) {
	tunner := tune.NewTunner(nil)

	tn, errL := tunner.Tune("https://open.spotify.com/track/1Dr1fXbc2IxaK1Mu8P8Khz?si=U8nwPqpCT2GyXnkF1NW0Yg")
	if errL != nil {
		t.Fatal(errL.Error())
	}

	ym := NewYandexMusic([]proxy.HttpProxyClient{})
	link, err := ym.Seek(tn)
	if err != nil {
		t.Fatal(err.Error())
	}

	if link == nil {
		t.Log("no link")

		return
	}

	t.Log("link", *link)
}
