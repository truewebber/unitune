package link_info

import (
	"lib/streamer"
)

type (
	AppleMusic struct {
		trackLink string

		actorId    int64
		actorTitle string

		albomId    int64
		albomTitle string
		albomType  string

		trackId    int64
		trackTitle string
	}
)

func NewAppleMusic(link string) (*AppleMusic, error) {
	return &AppleMusic{
		trackLink: link,
	}, nil
}

func (a *AppleMusic) Link() string {
	return a.trackLink
}

func (a *AppleMusic) Actor() string {
	return a.actorTitle
}

func (a *AppleMusic) Albom() string {
	return a.albomTitle
}

func (a *AppleMusic) AlbomType() string {
	return a.albomType
}

func (a *AppleMusic) Track() string {
	return a.trackTitle
}

func (a *AppleMusic) StreamerType() streamer.Type {
	return streamer.TypeAppleMusic
}
