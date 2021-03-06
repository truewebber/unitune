package tune

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/pkg/errors"

	"github.com/truewebber/unitune/streamer"
)

type (
	spotifyResponse struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					SSpotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			ExternalUrls struct {
				SSpotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				SSpotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		DiscNumber  int  `json:"disc_number"`
		DurationMs  int  `json:"duration_ms"`
		Explicit    bool `json:"explicit"`
		ExternalIds struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			SSpotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		IsPlayable  bool   `json:"is_playable"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	}

	spotifyTune struct {
		trackLink string

		artistId    int64
		artistTitle string

		albumId    int64
		albumTitle string
		albumType  string
		albumPic   string

		trackId    int64
		trackTitle string
	}
)

var (
	spotifyPayloadRegex = regexp.MustCompile("<script>\\s*Spotify\\s*=\\s*{};\\s*Spotify.Entity\\s*=\\s*(.*);\\s*</script>")
)

func newSpotifyTune(link string) (*spotifyTune, error) {
	resp, err := http.DefaultClient.Get(link)
	if err != nil {
		return nil, errors.Errorf("Error request Spotify link info, link: `%s`, error: %s",
			link, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("Spotify link return non-200 status code, link: `%s`, code: %d",
			link, resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := spotifyPayloadRegex.FindSubmatch(data)
	if len(results) != 2 {
		return nil, errors.Errorf("Error parse Spotify payload, body: %s", string(data))
	}

	obj := new(spotifyResponse)
	err = json.Unmarshal(results[1], obj)
	if err != nil {
		return nil, errors.Errorf("Error unmarshal Spotify json payload, error: %s, data: %s", err.Error(),
			string(results[1]))
	}

	img := ""
	if len(obj.Album.Images) != 0 {
		img = obj.Album.Images[0].URL
	}

	return &spotifyTune{
		trackLink:   link,
		artistTitle: obj.Artists[0].Name,
		albumTitle:  obj.Album.Name,
		albumType:   obj.Album.AlbumType,
		albumPic:    img,
		trackTitle:  obj.Name,
	}, nil
}

func (s *spotifyTune) Link() string {
	return s.trackLink
}

func (s *spotifyTune) Artist() string {
	return s.artistTitle
}

func (s *spotifyTune) Album() string {
	return s.albumTitle
}

func (s *spotifyTune) AlbumType() string {
	return s.albumType
}

func (s *spotifyTune) AlbumPic() string {
	return s.albumPic
}

func (s *spotifyTune) Track() string {
	return s.trackTitle
}

func (s *spotifyTune) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}
