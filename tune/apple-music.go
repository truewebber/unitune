package tune

import (
	"github.com/truewebber/unitune/streamer"
)

type (
	appleMusicTune struct {
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

func newAppleMusicTune(link string) (*appleMusicTune, error) {
	return &appleMusicTune{
		trackLink: link,
	}, nil
}

func (a *appleMusicTune) Link() string {
	return a.trackLink
}

func (a *appleMusicTune) Artist() string {
	return a.artistTitle
}

func (a *appleMusicTune) Album() string {
	return a.albumTitle
}

func (a *appleMusicTune) AlbumType() string {
	return a.albumType
}

func (a *appleMusicTune) Track() string {
	return a.trackTitle
}

func (a *appleMusicTune) StreamerType() streamer.Type {
	return streamer.TypeAppleMusic
}
