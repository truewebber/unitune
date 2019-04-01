package seeker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/proxy"

	"lib/link-info"
	"lib/streamer"
)

type (
	YMSearchResponse struct {
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

	YandexMusic struct {
		client *http.Client
	}
)

const (
	yandexMusicScheme   = "https"
	yandexMusicHost     = "music.yandex.ru"
	yandexMusicPath     = "/handlers/music-search.jsx"
	yandexMusicTemplate = "https://music.yandex.ru/album/%d/track/%d"
)
const (
	yandexProxyAddress = "46.16.13.212:3001"
)

func NewYandexMusic() *YandexMusic {
	socksPrx, err := proxy.SOCKS5("tcp", yandexProxyAddress, nil, proxy.Direct)
	if err != nil {
		fmt.Println("Error connecting to proxy:", err)
	}

	return &YandexMusic{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, network, addr string) (net.Conn, error) {
					return socksPrx.Dial(network, addr)
				},
			},
		},
	}
}

func (y *YandexMusic) request(r *http.Request) (*http.Response, error) {
	var (
		resp    *http.Response
		respErr error
	)

	max := 3
	for i := 0; i <= max; i++ {
		var err error
		resp, err = y.client.Do(r)

		if err != nil {
			respErr = errors.Errorf("Error do request search tune in YandexMusic, error: %s", err.Error())

			if i != max {
				time.Sleep(time.Duration(i+1) * time.Second)
			}
			continue
		}

		if resp.StatusCode != http.StatusOK {
			respErr = errors.Errorf("YandexMusic search return non-200 status code, code: %d", resp.StatusCode)

			if i != max {
				time.Sleep(time.Duration(i+1) * time.Second)
			}
			continue
		}

		break
	}

	return resp, respErr
}

func (y *YandexMusic) Seek(tune link_info.Tune) (*string, error) {
	rUrl := url.URL{
		Scheme: yandexMusicScheme,
		Host:   yandexMusicHost,
		Path:   yandexMusicPath,
	}

	query := fmt.Sprintf("%s - %s", tune.Artist(), tune.Track())

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
		return nil, errors.Errorf("Error create search request YandexMusic, error: %s, tune: %v", err.Error(), tune)
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("DNT", "1")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en,ru;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")

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

	resp, err := y.request(req)
	if err != nil {
		return nil, errors.Errorf("%s, tune: %v", err.Error(), tune)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error read YandexMusic search response, error: %s, tune: %v", err.Error(), tune)
	}

	obj := new(YMSearchResponse)
	err = json.Unmarshal(body, obj)
	if err != nil {
		return nil, errors.Errorf("Error unmarshal YandexMusic search response, error: %s, body: %s",
			err.Error(), string(body))
	}

	if obj.Counts.Tracks != 0 {
		track := obj.Tracks.Items[0]
		album := track.Albums[0]
		artist := track.Artists[0]

		if strings.ToLower(track.Title) != strings.ToLower(tune.Track()) {
			println(strings.ToLower(track.Title), strings.ToLower(tune.Track()))

			return nil, nil
		}

		if strings.ToLower(artist.Name) != strings.ToLower(tune.Artist()) {
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

func (y *YandexMusic) StreamerType() streamer.Type {
	return streamer.TypeYandexMusic
}
