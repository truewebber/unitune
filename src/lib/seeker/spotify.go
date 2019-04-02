package seeker

import (
	"lib/streamer"
	"lib/tune"
)

type (
	Spotify struct {
	}
)

func NewSpotify() *Spotify {
	return &Spotify{}
}

func (s *Spotify) Seek(tune tune.Tune) (*string, error) {
	return nil, nil
}

func (s *Spotify) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}
