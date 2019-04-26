package tune

import (
	"github.com/pkg/errors"

	"github.com/truewebber/unitune/link"
	"github.com/truewebber/unitune/proxy"
	"github.com/truewebber/unitune/streamer"
)

type (
	Tune interface {
		StreamerType() streamer.Type
		Link() string
		Artist() string
		Album() string
		AlbumType() string
		Track() string
	}

	Tunner struct {
		proxyList []proxy.HttpProxyClient
	}
)

var (
	UnknownType = errors.Errorf("Unknown type of link")
)

func NewTunner(proxyList []proxy.HttpProxyClient) *Tunner {
	if proxyList == nil {
		proxyList = make([]proxy.HttpProxyClient, 0)
	}

	return &Tunner{
		proxyList: proxyList,
	}
}

func (t *Tunner) Tune(trackLink string) (Tune, error) {
	switch {
	case link.IsYandexMusic(trackLink):
		{
			return newYandexMusicTune(t.proxyList, trackLink)
		}
	case link.IsSpotify(trackLink):
		{
			return newSpotifyTune(trackLink)
		}
	case link.IsAppleMusic(trackLink):
		{
			return newAppleMusicTune(trackLink)
		}
	}

	return nil, UnknownType
}
