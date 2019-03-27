package seeker

import (
	"lib/link-info"
	"lib/streamer"
)

type (
	AppleMusic struct {
	}
)

func NewAppleMusic() *AppleMusic {
	return &AppleMusic{}
}

func (a *AppleMusic) Seek(tune link_info.Tune) (string, error) {
	return "", nil
}

func (a *AppleMusic) StreamerType() streamer.Type {
	return streamer.TypeAppleMusic
}
