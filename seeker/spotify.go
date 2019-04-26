package seeker

import (
	"fmt"
	"os"
	"strings"

	"github.com/mgutz/logxi/v1"
	"github.com/pkg/errors"
	sptf "github.com/zmb3/spotify"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/truewebber/unitune/streamer"
	"github.com/truewebber/unitune/tune"
)

type (
	spotify struct {
		client sptf.Client
	}
)

const (
	spotifyDefaultExternalURLKey = "spotify"
)

func newSpotify() *spotify {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     sptf.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatal("Error login Spotify", "error", err.Error())
	}

	client := sptf.Authenticator{}.NewClient(token)

	return &spotify{
		client: client,
	}
}

func (s *spotify) Seek(t tune.Tune) (*string, error) {
	query := fmt.Sprintf("%s - %s", t.Artist(), t.Track())

	r, err := s.client.Search(query, sptf.SearchTypeTrack)
	if err != nil {
		return nil, errors.Errorf("Error request Spotify response: %s, tune: %v", err.Error(), t)
	}

	for i := 0; i < len(r.Tracks.Tracks); i++ {
		got := r.Tracks.Tracks[i]

		if strings.ToLower(t.Track()) == strings.ToLower(got.Name) &&
			strings.ToLower(t.Artist()) == strings.ToLower(got.Artists[0].Name) {
			if len(got.ExternalURLs) > 0 {
				if link, ok := got.ExternalURLs[spotifyDefaultExternalURLKey]; ok {
					return &link, nil
				}

				log.Debug("Spotify return non-empty ExternalURLs, but there are no default key, tune: %v, links: %#v",
					t, got.ExternalURLs)
			}
		}
	}

	log.Debug("Spotify not found", "query", query, "total results", r.Tracks.Total)

	return nil, nil
}

func (s *spotify) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}
