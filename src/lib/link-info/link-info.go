package link_info

import (
	"github.com/pkg/errors"

	"lib/link"
	"lib/streamer"
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
)

var (
	UnknownType = errors.Errorf("Unknown type of link")
)

func GetLinkInfo(trackLink string) (Tune, error) {
	switch {
	case link.IsYandexMusic(trackLink):
		{
			return NewYandexMusic(trackLink)
		}
	case link.IsSpotify(trackLink):
		{
			return NewSpotify(trackLink)
		}
	case link.IsAppleMusic(trackLink):
		{
			return NewAppleMusic(trackLink)
		}
	}

	return nil, UnknownType
}
