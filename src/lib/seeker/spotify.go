package seeker

import (
	"lib/link-info"
	"lib/streamer"
)

type (
	Spotify struct {
	}
)

func NewSpotify() *Spotify {
	return &Spotify{}
}

func (s *Spotify) Seek(tune link_info.Tune) (*string, error) {
	return nil, nil
}

func (s *Spotify) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}
