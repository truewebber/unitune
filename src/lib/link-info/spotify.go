package link_info

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

type (
	SpotifyResponse struct {
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

	Spotify struct {
		trackLink string

		ActorId    int64
		ActorTitle string

		AlbomId    int64
		AlbomTitle string
		AlbomType  string

		TrackId    int64
		TrackTitle string
	}
)

var (
	SpotifyPayloadRegex = regexp.MustCompile("<script>\\s*Spotify\\s*=\\s*{};\\s*Spotify.Entity\\s*=\\s*(.*);\\s*</script>")
)

func NewSpotify(link string) (*Spotify, error) {
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

	results := SpotifyPayloadRegex.FindSubmatch(data)
	if len(results) != 2 {
		return nil, errors.Errorf("Error parse Spotify payload, body: %s", string(data))
	}

	obj := new(SpotifyResponse)
	err = json.Unmarshal(results[1], obj)
	if err != nil {
		return nil, errors.Errorf("Error unmarshal Spotify json payload, error: %s, data: %s", err.Error(),
			string(results[1]))
	}

	return &Spotify{
		trackLink:  link,
		ActorTitle: obj.Artists[0].Name,
		AlbomTitle: obj.Album.Name,
		AlbomType:  obj.Album.AlbumType,
		TrackTitle: obj.Name,
	}, nil
}

func (s *Spotify) GetLink() string {
	return s.trackLink
}

func (s *Spotify) GetActor() string {
	return s.ActorTitle
}

func (s *Spotify) GetAlbom() string {
	return s.AlbomTitle
}

func (s *Spotify) GetAlbomType() string {
	return s.AlbomType
}

func (s *Spotify) GetTrack() string {
	return s.TrackTitle
}
