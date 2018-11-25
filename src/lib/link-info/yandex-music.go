package link_info

import (
	"encoding/json"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type (
	YMResponse struct {
		Type        string `json:"@type"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Duration    string `json:"duration"`
		URL         string `json:"url"`
		Context     string `json:"@context"`
		InAlbum     struct {
			Context     string `json:"@context"`
			Type        string `json:"@type"`
			Description string `json:"description"`
			Name        string `json:"name"`
			NumTracks   int    `json:"numTracks"`
			Genre       string `json:"genre"`
			Image       string `json:"image"`
			URL         string `json:"url"`
			Track       []struct {
				Type     string `json:"@type"`
				Duration string `json:"duration"`
				Name     string `json:"name"`
				URL      string `json:"url"`
			} `json:"track"`
			ByArtist struct {
				Type string `json:"@type"`
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"byArtist"`
			PotentialAction struct {
				Type   string `json:"@type"`
				Target []struct {
					Type        string `json:"@type"`
					URLTemplate string `json:"urlTemplate"`
					InLanguage  string `json:"inLanguage"`
					//ActionPlatform    []string `json:"actionPlatform"`
					ActionApplication struct {
						Type                string `json:"@type"`
						ID                  string `json:"@id"`
						Name                string `json:"name"`
						ApplicationCategory string `json:"applicationCategory"`
						OperatingSystem     string `json:"operatingSystem"`
					} `json:"actionApplication,omitempty"`
				} `json:"target"`
				ExpectsAcceptanceOf struct {
					Type           string `json:"@type"`
					EligibleRegion []struct {
						Type string `json:"@type"`
						Name string `json:"name"`
					} `json:"eligibleRegion"`
					Category string `json:"category"`
				} `json:"expectsAcceptanceOf"`
			} `json:"potentialAction"`
		} `json:"inAlbum"`
		ByArtist struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"byArtist"`
		PotentialAction struct {
			Type   string `json:"@type"`
			Target []struct {
				Type        string `json:"@type"`
				URLTemplate string `json:"urlTemplate"`
				InLanguage  string `json:"inLanguage"`
				//ActionPlatform    []string `json:"actionPlatform"`
				ActionApplication struct {
					Type                string `json:"@type"`
					ID                  string `json:"@id"`
					Name                string `json:"name"`
					ApplicationCategory string `json:"applicationCategory"`
					OperatingSystem     string `json:"operatingSystem"`
				} `json:"actionApplication,omitempty"`
			} `json:"target"`
			ExpectsAcceptanceOf struct {
				Type           string `json:"@type"`
				EligibleRegion []struct {
					Type string `json:"@type"`
					Name string `json:"name"`
				} `json:"eligibleRegion"`
				Category string `json:"category"`
			} `json:"expectsAcceptanceOf"`
		} `json:"potentialAction"`
	}

	YandexMusic struct {
		trackLink string

		ActorId    int64
		AlbomId    int64
		TrackId    int64
		ActorTitle string
		AlbomTitle string
		TrackTitle string
	}
)

func NewYandexMusic(link string) (*YandexMusic, error) {
	resp, err := http.DefaultClient.Get(link)
	if err != nil {
		return nil, errors.Errorf("Error request Yandex Music link info, link: `%s`, error: %s",
			link, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("Yandex Music link return non-200 status code, link: `%s`, code: %d",
			link, resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error parse Yandex Music response body, link: `%s`, error: %s",
			link, err.Error())
	}

	selection := doc.Find("script.light-data")
	if len(selection.Nodes) == 0 {
		return nil, errors.Errorf("No info while parse Yandex Music response body, link: `%s`", link)
	}

	obj := new(YMResponse)
	err = json.Unmarshal([]byte(selection.Nodes[0].FirstChild.Data), obj)
	if err != nil {
		return nil, errors.Errorf("Error parse YM json response, link: `%s`, error: %s", link, err.Error())
	}

	return &YandexMusic{
		trackLink:  link,
		ActorTitle: obj.ByArtist.Name,
		AlbomTitle: obj.InAlbum.Name,
		TrackTitle: obj.Name,
	}, nil
}

func (y *YandexMusic) GetLink() string {
	return y.trackLink
}

func (y *YandexMusic) GetActor() string {
	return y.ActorTitle
}

func (y *YandexMusic) GetAlbom() string {
	return y.AlbomTitle
}

func (y *YandexMusic) GetTrack() string {
	return y.TrackTitle
}
