package seeker

import (
	"github.com/pkg/errors"

	"lib/proxy"
	"lib/streamer"
	"lib/tune"
)

type (
	Seeker interface {
		Seek(tune.Tune) (*string, error)
		StreamerType() streamer.Type
	}
)

var (
	seekers = []Seeker{
		NewSpotify(),
		NewAppleMusic(),
		NewYandexMusic([]proxy.HttpProxyClient{}),
	}
)

func LookUpTune(tune tune.Tune) ([]string, []error) {
	errList := make([]error, 0, 2)
	links := make([]string, 0, 2)

	for _, seeker := range seekers {
		if seeker.StreamerType() == tune.StreamerType() {
			continue
		}

		link, err := seeker.Seek(tune)
		if err != nil {
			errList = append(errList, errors.Errorf("Error seek `%s` tune in `%s`, link: %s, error: %s",
				tune.StreamerType(), seeker.StreamerType(), tune.Link(), err.Error()))

			continue
		}

		if link == nil {
			continue
		}

		links = append(links, *link)
	}

	return links, errList
}
