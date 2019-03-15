package link_info

import (
	"github.com/pkg/errors"

	"lib/link"
)

type (
	Streamer interface {
		GetLink() string
		GetActor() string
		GetAlbom() string
		GetAlbomType() string
		GetTrack() string
	}
)

var (
	UnknownType = errors.Errorf("Unknown type of link")
)

func GetLinkInfo(trackLink string) (Streamer, error) {
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

		}
	}

	return nil, UnknownType
}
