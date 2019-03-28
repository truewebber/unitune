package link_info

import (
	"lib/streamer"
)

type (
	AppleMusic struct {
		trackLink string

		artistId    int64
		artistTitle string

		albumId    int64
		albumTitle string
		albumType  string

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

func (a *AppleMusic) Artist() string {
	return a.artistTitle
}

func (a *AppleMusic) Album() string {
	return a.albumTitle
}

func (a *AppleMusic) AlbumType() string {
	return a.albumType
}

func (a *AppleMusic) Track() string {
	return a.trackTitle
}

func (a *AppleMusic) StreamerType() streamer.Type {
	return streamer.TypeAppleMusic
}
