package seeker

import (
	"github.com/truewebber/unitune/streamer"
	"github.com/truewebber/unitune/tune"
)

type (
	appleMusic struct {
	}
)

func newAppleMusic() *appleMusic {
	return &appleMusic{}
}

func (a *appleMusic) Seek(tune tune.Tune) (*string, error) {
	return nil, nil
}

func (a *appleMusic) StreamerType() streamer.Type {
	return streamer.TypeAppleMusic
}
