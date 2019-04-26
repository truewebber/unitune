package seeker

import (
	"github.com/pkg/errors"

	"github.com/truewebber/unitune/proxy"
	"github.com/truewebber/unitune/streamer"
	"github.com/truewebber/unitune/tune"
)

type (
	seeker interface {
		Seek(tune.Tune) (*string, error)
		StreamerType() streamer.Type
	}

	MasterSeeker struct {
		seekers []seeker
	}
)

func New(proxyList []proxy.HttpProxyClient) *MasterSeeker {
	return &MasterSeeker{
		seekers: []seeker{
			newSpotify(),
			newAppleMusic(),
			newYandexMusic(proxyList),
		},
	}
}

func (m *MasterSeeker) LookUpTune(tune tune.Tune) ([]string, []error) {
	errList := make([]error, 0, 2)
	links := make([]string, 0, 2)

	for _, seeker := range m.seekers {
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
