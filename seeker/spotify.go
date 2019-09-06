package seeker

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

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
		sync.Mutex
		config *clientcredentials.Config
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

	s := &spotify{
		config: config,
	}

	err := s.refreshSpotifyToken()
	if err != nil {
		log.Fatal("Error login Spotify", "error", err.Error())
	}

	return s
}

func (s *spotify) Seek(t tune.Tune) (*string, error) {
	return s.seek(t, 0)
}

func (s *spotify) seek(t tune.Tune, retry int) (*string, error) {
	query := fmt.Sprintf("%s - %s", t.Artist(), t.Track())

	r, err := s.client.Search(query, sptf.SearchTypeTrack)
	if err != nil {
		spotifyErr, ok := err.(sptf.Error)
		if ok && spotifyErr.Status == http.StatusUnauthorized && retry < 3 {
			if err := s.refreshSpotifyToken(); err != nil {
				log.Error("Error refresh Spotify token", "error", err.Error())
			}

			return s.seek(t, retry+1)
		}

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

func (s *spotify) refreshSpotifyToken() error {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	token, err := s.config.Token(context.Background())
	if err != nil {
		return err
	}
	s.client = sptf.Authenticator{}.NewClient(token)

	return nil
}
