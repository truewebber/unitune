package seeker

import (
	"lib/streamer"
	"lib/tune"
)

type (
	spotify struct {
	}
)

func newSpotify() *spotify {
	return &spotify{}
}

func (s *spotify) Seek(tune tune.Tune) (*string, error) {
	return nil, nil
}

func (s *spotify) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}
