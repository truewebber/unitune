package seeker

import (
	"testing"

	"lib/link-info"
)

func TestYandexMusic_Seek(t *testing.T) {
	tune, errL := link_info.NewSpotify("https://open.spotify.com/track/1Dr1fXbc2IxaK1Mu8P8Khz?si=U8nwPqpCT2GyXnkF1NW0Yg")
	if errL != nil {
		t.Fatal(errL.Error())
	}

	ym := NewYandexMusic()
	link, err := ym.Seek(tune)
	if err != nil {
		t.Fatal(err.Error())
	}

	if link == nil {
		t.Log("no link")

		return
	}

	t.Log("link", *link)
}
