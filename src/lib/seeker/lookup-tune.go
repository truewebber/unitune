package seeker

import (
	"github.com/pkg/errors"

	"lib/link-info"
	"lib/streamer"
)

type (
	Seeker interface {
		Seek(link_info.Tune) (*string, error)
		StreamerType() streamer.Type
	}
)

var (
	seekers = []Seeker{
		NewSpotify(),
		NewAppleMusic(),
		NewYandexMusic(),
	}
)

func LookUpTune(tune link_info.Tune) ([]string, []error) {
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
