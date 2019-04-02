package seeker

import (
	"lib/streamer"
	"lib/tune"
)

type (
	AppleMusic struct {
	}
)

func NewAppleMusic() *AppleMusic {
	return &AppleMusic{}
}

func (a *AppleMusic) Seek(tune tune.Tune) (*string, error) {
	return nil, nil
}

func (a *AppleMusic) StreamerType() streamer.Type {
	return streamer.TypeAppleMusic
}
