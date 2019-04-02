package seeker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mgutz/logxi/v1"
	"github.com/pkg/errors"

	"lib/proxy"
	"lib/streamer"
	"lib/tune"
)

type (
	ymSearchResponse struct {
		Text   string `json:"text"`
		Albums struct {
			Items []struct {
				ID                  int    `json:"id"`
				StorageDir          string `json:"storageDir"`
				OriginalReleaseYear int    `json:"originalReleaseYear"`
				Year                int    `json:"year"`
				Type                string `json:"type"`
				Artists             []struct {
					ID    int `json:"id"`
					Cover struct {
						Type   string `json:"type"`
						Prefix string `json:"prefix"`
						URI    string `json:"uri"`
					} `json:"cover"`
					Composer   bool          `json:"composer"`
					Name       string        `json:"name"`
					Various    bool          `json:"various"`
					Decomposed []interface{} `json:"decomposed"`
				} `json:"artists"`
				CoverURI                 string   `json:"coverUri"`
				TrackCount               int      `json:"trackCount"`
				Genre                    string   `json:"genre"`
				Available                bool     `json:"available"`
				AvailableForPremiumUsers bool     `json:"availableForPremiumUsers"`
				Title                    string   `json:"title"`
				Regions                  []string `json:"regions"`
				Version                  string   `json:"version,omitempty"`
			} `json:"items"`
			Total   int `json:"total"`
			PerPage int `json:"perPage"`
		} `json:"albums"`
		Tracks struct {
			Items []struct {
				ID                       int  `json:"id"`
				Available                bool `json:"available"`
				AvailableAsRbt           bool `json:"availableAsRbt"`
				AvailableForPremiumUsers bool `json:"availableForPremiumUsers"`
				Albums                   []struct {
					ID                       int           `json:"id"`
					StorageDir               string        `json:"storageDir"`
					OriginalReleaseYear      int           `json:"originalReleaseYear"`
					Year                     int           `json:"year"`
					Artists                  []interface{} `json:"artists"`
					CoverURI                 string        `json:"coverUri"`
					TrackCount               int           `json:"trackCount"`
					Genre                    string        `json:"genre"`
					Available                bool          `json:"available"`
					AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
					Title                    string        `json:"title"`
					TrackPosition            struct {
						Volume int `json:"volume"`
						Index  int `json:"index"`
					} `json:"trackPosition"`
				} `json:"albums"`
				StorageDir string `json:"storageDir"`
				DurationMs int    `json:"durationMs"`
				Explicit   bool   `json:"explicit"`
				Title      string `json:"title"`
				Artists    []struct {
					ID    int `json:"id"`
					Cover struct {
						Type   string `json:"type"`
						Prefix string `json:"prefix"`
						URI    string `json:"uri"`
					} `json:"cover"`
					Composer   bool          `json:"composer"`
					Name       string        `json:"name"`
					Various    bool          `json:"various"`
					Decomposed []interface{} `json:"decomposed"`
				} `json:"artists"`
				Regions []string `json:"regions"`
				Version string   `json:"version,omitempty"`
			} `json:"items"`
			Total   int `json:"total"`
			PerPage int `json:"perPage"`
		} `json:"tracks"`
		Artists struct {
			Items []interface{} `json:"items"`
		} `json:"artists"`
		Videos struct {
			Items []struct {
				Cover    string `json:"cover"`
				URL      string `json:"url"`
				Title    string `json:"title"`
				Duration int    `json:"duration"`
				Embed    string `json:"embed"`
				Text     string `json:"text"`
			} `json:"items"`
			Total   int `json:"total"`
			PerPage int `json:"perPage"`
		} `json:"videos"`
		Playlists struct {
			Items []interface{} `json:"items"`
		} `json:"playlists"`
		Users struct {
			Items []interface{} `json:"items"`
		} `json:"users"`
		RequestID       string `json:"request_id"`
		SearchRequestID string `json:"searchRequestId"`
		ExtendedParams  struct {
		} `json:"extendedParams"`
		Misspell struct {
			Corrected bool `json:"corrected"`
			Nocorrect bool `json:"nocorrect"`
		} `json:"misspell"`
		Best struct {
			Type string `json:"type"`
			Item struct {
				ID                       int  `json:"id"`
				Available                bool `json:"available"`
				AvailableAsRbt           bool `json:"availableAsRbt"`
				AvailableForPremiumUsers bool `json:"availableForPremiumUsers"`
				Albums                   []struct {
					ID                       int           `json:"id"`
					StorageDir               string        `json:"storageDir"`
					OriginalReleaseYear      int           `json:"originalReleaseYear"`
					Year                     int           `json:"year"`
					Artists                  []interface{} `json:"artists"`
					CoverURI                 string        `json:"coverUri"`
					TrackCount               int           `json:"trackCount"`
					Genre                    string        `json:"genre"`
					Available                bool          `json:"available"`
					AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
					Title                    string        `json:"title"`
					TrackPosition            struct {
						Volume int `json:"volume"`
						Index  int `json:"index"`
					} `json:"trackPosition"`
				} `json:"albums"`
				StorageDir string `json:"storageDir"`
				DurationMs int    `json:"durationMs"`
				Explicit   bool   `json:"explicit"`
				Title      string `json:"title"`
				Artists    []struct {
					ID    int `json:"id"`
					Cover struct {
						Type   string `json:"type"`
						Prefix string `json:"prefix"`
						URI    string `json:"uri"`
					} `json:"cover"`
					Composer   bool          `json:"composer"`
					Name       string        `json:"name"`
					Various    bool          `json:"various"`
					Decomposed []interface{} `json:"decomposed"`
				} `json:"artists"`
				Regions []string `json:"regions"`
			} `json:"item"`
		} `json:"best"`
		Counts struct {
			Artists   int `json:"artists"`
			Albums    int `json:"albums"`
			Tracks    int `json:"tracks"`
			Videos    int `json:"videos"`
			Playlists int `json:"playlists"`
			Users     int `json:"users"`
		} `json:"counts"`
	}

	yandexMusic struct {
		proxyList []proxy.HttpProxyClient
	}
)

const (
	yandexMusicScheme   = "https"
	yandexMusicHost     = "music.yandex.ru"
	yandexMusicPath     = "/handlers/music-search.jsx"
	yandexMusicTemplate = "https://music.yandex.ru/album/%d/track/%d"
)

var (
	ymTryAgainError = errors.New("YandexMusic 451 code or proxy error")
)

func newYandexMusic(proxyList []proxy.HttpProxyClient) *yandexMusic {
	if len(proxyList) == 0 {
		proxyList = []proxy.HttpProxyClient{proxy.NewNull()}
	}

	return &yandexMusic{
		proxyList: proxyList,
	}
}

func (y *yandexMusic) proxyRequest(r *http.Request) (*http.Response, error) {
	var (
		resp    *http.Response
		respErr error
	)

	for _, prx := range y.proxyList {
		log.Debug("YM request", "proxy", prx.Type().String())

		resp, respErr = y.request(prx.HttpClient(), r)
		if errors.Cause(respErr) == ymTryAgainError {
			continue
		}

		break
	}

	return resp, respErr
}

func (y *yandexMusic) request(client *http.Client, r *http.Request) (*http.Response, error) {
	var (
		resp    *http.Response
		respErr error
	)

	max := 3
	for i := 0; i <= max; i++ {
		var err error
		resp, err = client.Do(r)

		if err != nil {
			respErr = errors.Wrapf(ymTryAgainError, "Error do request search tune in YandexMusic, error: %s", err.Error())

			if i != max {
				time.Sleep(time.Duration(i+1) * time.Second)
			}
			continue
		}

		if resp.StatusCode != http.StatusOK {
			if resp.StatusCode == http.StatusUnavailableForLegalReasons {
				respErr = errors.Wrapf(ymTryAgainError, "YandexMusic search return non-200 status code, code: %d", resp.StatusCode)

				if i != max {
					time.Sleep(time.Duration(i+1) * time.Second)
				}
				continue
			}

			respErr = errors.Errorf("YandexMusic search return non-200 status code, code: %d", resp.StatusCode)
		}

		break
	}

	return resp, respErr
}

func (y *yandexMusic) Seek(t tune.Tune) (*string, error) {
	rUrl := url.URL{
		Scheme: yandexMusicScheme,
		Host:   yandexMusicHost,
		Path:   yandexMusicPath,
	}

	query := fmt.Sprintf("%s - %s", t.Artist(), t.Track())

	q := url.Values{}
	q.Add("text", query)
	q.Add("type", "all")
	q.Add("lang", "ru")
	q.Add("external-domain", "music.yandex.ru")
	q.Add("overembed", "false")
	rand.Seed(time.Now().UnixNano())
	q.Add("ncrnd", strconv.FormatFloat(rand.Float64(), 'f', 10, 64))
	rUrl.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, rUrl.String(), nil)
	if err != nil {
		return nil, errors.Errorf("Error create search request YandexMusic, error: %s, tune: %v", err.Error(), t)
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("DNT", "1")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en,ru;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")

	refUrl := url.URL{
		Scheme: yandexMusicScheme,
		Host:   yandexMusicHost,
		Path:   "/search",
	}
	qRef := url.Values{}
	qRef.Add("text", query)
	refUrl.RawQuery = qRef.Encode()

	req.Header.Add("X-Retpath-Y", refUrl.String())
	req.Header.Add("Referer", refUrl.String())

	resp, err := y.proxyRequest(req)
	if err != nil {
		return nil, errors.Errorf("%s, tune: %v", err.Error(), t)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error read YandexMusic search response, error: %s, tune: %v", err.Error(), t)
	}

	obj := new(ymSearchResponse)
	err = json.Unmarshal(body, obj)
	if err != nil {
		return nil, errors.Errorf("Error unmarshal YandexMusic search response, error: %s, body: %s",
			err.Error(), string(body))
	}

	if obj.Counts.Tracks != 0 {
		track := obj.Tracks.Items[0]
		album := track.Albums[0]
		artist := track.Artists[0]

		if strings.ToLower(track.Title) != strings.ToLower(t.Track()) {
			println(strings.ToLower(track.Title), strings.ToLower(t.Track()))

			return nil, nil
		}

		if strings.ToLower(artist.Name) != strings.ToLower(t.Artist()) {
			return nil, nil
		}

		link := fmt.Sprintf(yandexMusicTemplate, album.ID, track.ID)

		return &link, nil
	}

	if obj.Counts.Videos != 0 {
		video := obj.Videos.Items[0]

		link := video.URL

		return &link, nil
	}

	return nil, nil
}

func (y *yandexMusic) StreamerType() streamer.Type {
	return streamer.TypeYandexMusic
}
