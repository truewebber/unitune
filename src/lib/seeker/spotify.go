package seeker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	sptf "github.com/zmb3/spotify"

	"lib/streamer"
	"lib/tune"
)

type (
	spotify struct {
		auth  sptf.Authenticator
		ch    chan *sptf.Client
		state string
	}

	spotifyResponse struct {
		BestMatch struct {
			Items []struct {
				Album struct {
					AlbumType string `json:"album_type"`
					Artists   []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						ID   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						URI  string `json:"uri"`
					} `json:"artists"`
					ExternalUrls struct {
						Spotify string `json:"spotify"`
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
						Spotify string `json:"spotify"`
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
					Spotify string `json:"spotify"`
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
			} `json:"items"`
		} `json:"best_match"`
		Albums struct {
			Href     string        `json:"href"`
			Items    []interface{} `json:"items"`
			Limit    int           `json:"limit"`
			Next     interface{}   `json:"next"`
			Offset   int           `json:"offset"`
			Previous interface{}   `json:"previous"`
			Total    int           `json:"total"`
		} `json:"albums"`
		Artists struct {
			Href     string        `json:"href"`
			Items    []interface{} `json:"items"`
			Limit    int           `json:"limit"`
			Next     interface{}   `json:"next"`
			Offset   int           `json:"offset"`
			Previous interface{}   `json:"previous"`
			Total    int           `json:"total"`
		} `json:"artists"`
		Tracks struct {
			Href  string `json:"href"`
			Items []struct {
				Album struct {
					AlbumType string `json:"album_type"`
					Artists   []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						ID   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						URI  string `json:"uri"`
					} `json:"artists"`
					ExternalUrls struct {
						Spotify string `json:"spotify"`
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
						Spotify string `json:"spotify"`
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
					Spotify string `json:"spotify"`
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
			} `json:"items"`
			Limit    int         `json:"limit"`
			Next     interface{} `json:"next"`
			Offset   int         `json:"offset"`
			Previous interface{} `json:"previous"`
			Total    int         `json:"total"`
		} `json:"tracks"`
		Playlists struct {
			Href  string `json:"href"`
			Items []struct {
				Collaborative bool `json:"collaborative"`
				ExternalUrls  struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name  string `json:"name"`
				Owner struct {
					DisplayName  string `json:"display_name"`
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"owner"`
				PrimaryColor interface{} `json:"primary_color"`
				Public       interface{} `json:"public"`
				SnapshotID   string      `json:"snapshot_id"`
				Tracks       struct {
					Href  string `json:"href"`
					Total int    `json:"total"`
				} `json:"tracks"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"items"`
			Limit    int         `json:"limit"`
			Next     interface{} `json:"next"`
			Offset   int         `json:"offset"`
			Previous interface{} `json:"previous"`
			Total    int         `json:"total"`
		} `json:"playlists"`
		Shows struct {
			Href     string        `json:"href"`
			Items    []interface{} `json:"items"`
			Limit    int           `json:"limit"`
			Next     interface{}   `json:"next"`
			Offset   int           `json:"offset"`
			Previous interface{}   `json:"previous"`
			Total    int           `json:"total"`
		} `json:"shows"`
		Episodes struct {
			Href     string        `json:"href"`
			Items    []interface{} `json:"items"`
			Limit    int           `json:"limit"`
			Next     interface{}   `json:"next"`
			Offset   int           `json:"offset"`
			Previous interface{}   `json:"previous"`
			Total    int           `json:"total"`
		} `json:"episodes"`
	}
)

const (
	spotifyScheme = "https"
	spotifyHost   = "api.spotify.com"
	spotifyPath   = "/v1/search"

	spotifyAuth = "Bearer BQDJNPrj5CFsXKRU0fG3MzaeUX62n8btrUkQjwNXU06YdoKyte5JHiYS9BP3rFxtJOTzdgIO48vDL8J0OLzj5p4FUpgmx43MSdrEFA211d7PKVzMh2e6JOHPugBFVg4npHp4wP5lN99QwgzIO2s8_CJO7Gi2xpEPyPaDf6J8fYCdFRbGgPr8u9Mt8D2xtJfHThEuEXk1FWmR11ENh0VTgbwUnhay6W8eVI6jdG_KUEtzusgiJ6zVMkCzNoKJKEFOApy1IskXLdko3AChZ7nwKk2cMmfNOQiu"
)

func newSpotify() *spotify {
	s := &spotify{
		auth:  sptf.NewAuthenticator("http://localhost:8080/callback", sptf.ScopeUserReadPrivate),
		ch:    make(chan *sptf.Client),
		state: "abc123",
	}

	go func() {
		http.HandleFunc("/callback", s.completeAuth)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Got request for:", r.URL.String())
		})
		go http.ListenAndServe(":8080", nil)

		authUrl := s.auth.AuthURL(s.state)
		fmt.Println("Please log in to Spotify by visiting the following page in your browser:", authUrl)

		// wait for auth to complete
		client := <-s.ch

		// use the client to make calls that require authorization
		user, err := client.CurrentUser()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You are logged in as:", user.ID)
	}()

	return s
}

func (s *spotify) completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := s.auth.Token(s.state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != s.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, s.state)
	}
	// use the token to get an authenticated client
	client := s.auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	s.ch <- &client
}

func (s *spotify) Seek(t tune.Tune) (*string, error) {
	rUrl := url.URL{
		Scheme: spotifyScheme,
		Host:   spotifyHost,
		Path:   spotifyPath,
	}

	query := fmt.Sprintf("%s - %s", t.Artist(), t.Track())

	q := url.Values{}
	q.Add("type", "album,artist,playlist,track,show_audio,episode_audio")
	q.Add("q", fmt.Sprintf("%s*", query))
	q.Add("decorate_restrictions", "true")
	q.Add("best_match", "true")
	q.Add("include_external", "audio")
	q.Add("limit", "50")
	q.Add("userless", "false")
	q.Add("market", "from_token")
	rUrl.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, rUrl.String(), nil)
	if err != nil {
		return nil, errors.Errorf("Error create search request Spotify, error: %s, tune: %v", err.Error(), t)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Referer", fmt.Sprintf("https://open.spotify.com/search/results/%s", query))
	req.Header.Add("Origin", "https://open.spotify.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Authorization", spotifyAuth)
	req.Header.Add("DNT", "1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Errorf("Error request Spotify response: %s, tune: %v", err.Error(), t)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error read Spotify search response, error: %s, tune: %v", err.Error(), t)
	}

	if resp.StatusCode >= 400 {
		return nil, errors.Errorf("Spotify return non-200 code: %d, body: %s", resp.StatusCode, string(body))
	}

	obj := new(spotifyResponse)
	err = json.Unmarshal(body, obj)
	if err != nil {
		return nil, errors.Errorf("Error unmarshal Spotify error: %s, body: %s", err.Error(), string(body))
	}

	for i := 0; i < len(obj.Tracks.Items); i++ {
		got := obj.Tracks.Items[i]

		if t.Track() == got.Name && t.Artist() == got.Artists[0].Name {
			return &got.ExternalUrls.Spotify, nil
		}
	}

	return nil, nil
}

func (s *spotify) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}

func login() error {
	u := url.URL{
		Scheme: "https",
		Host:   "accounts.spotify.com",
		Path:   "/api/login",
	}

	v := url.Values{}
	v.Add("remember", "false")
	v.Add("username", "kish94@mail.ru")
	v.Add("password", "mystery41")

	// recaptchaToken=03AOLTBLTAkomS4op1jt5MoXemMtBvsVsjDp2UPygWTPu0wWLVmrM3fVqlni_u-b5v7qhwUAqfuPVioM1ECn2021Y-EmB
	// ZlbSlCZme0LHJzdKPKVcXhzfdCo7NYOgBWgtlUsrivxr6bxsVoUodz2YoGyXQ6lD04xOPd2crdj3WJGdJyEZw7EvKiDIEToq9mVIE1I5_4
	// cBGrprYwWpg8eSKxEfQGuQEOl7pwksXgcIe1zBrDGdB88c5Yq8pI-4qLYPbujlJ1kIKHQK9BnI8eglR4aji0EYciowVr-RAuTuDtrYv4
	// R8jrYq9iCGkXMim_JTUVTd8NwfA_OZ8
	//
	// csrf_token=AQBhQsEbdCaYYrUU7kEiMfr6gZ6PtCvYcA-BDo0YoEY9gWyJ55rB5kPwLQ0UmoAf3tWWvbbJsYubZrykEw

	r, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Accept", "application/json, text/plain, */*")
	r.Header.Set("Origin", "https://accounts.spotify.com")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	r.Header.Set("DNT", "1")
	r.Header.Set("Referer", "https://accounts.spotify.com/en/login?continue=https:%2F%2Fopen.spotify.com%2Fbrowse%2Ffeatured")
	r.Header.Set("Accept-Language", "en,ru;q=0.9")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return errors.Errorf("Error request Auth Spotify response: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Errorf("Error read Auth Spotify response, error: %s", err.Error())
	}

	fmt.Printf("%v\n", resp.Header)
	println(string(body))

	return nil
}

// Login
//
// curl
// -H 'Host: accounts.spotify.com'
// -H 'Accept: application/json, text/plain, */*'
// -H 'Origin: https://accounts.spotify.com'
// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36'
// -H 'DNT: 1'
// -H 'Referer: https://accounts.spotify.com/en/login?continue=https:%2F%2Fopen.spotify.com%2Fbrowse%2Ffeatured'
// -H 'Accept-Language: en,ru;q=0.9'
// -H 'Cookie: __bon=MHwwfDE4MDgzMjcwODB8NzU5NDk3MzczNjB8MXwxfDF8MQ==; remember=kish94%40mail.ru; csrf_token=AQBhQsEbdCaYYrUU7kEiMfr6gZ6PtCvYcA-BDo0YoEY9gWyJ55rB5kPwLQ0UmoAf3tWWvbbJsYubZrykEw; fb_continue=https%3A%2F%2Fopen.spotify.com%2Fbrowse%2Ffeatured; _ga=GA1.2.674780686.1555665773; _gid=GA1.2.697624719.1555665773; _gat=1'
// --data "remember=true&username=kish94%40mail.ru&recaptchaToken=03AOLTBLTAkomS4op1jt5MoXemMtBvsVsjDp2UPygWTPu0wWLVmrM3fVqlni_u-b5v7qhwUAqfuPVioM1ECn2021Y-EmBZlbSlCZme0LHJzdKPKVcXhzfdCo7NYOgBWgtlUsrivxr6bxsVoUodz2YoGyXQ6lD04xOPd2crdj3WJGdJyEZw7EvKiDIEToq9mVIE1I5_4cBGrprYwWpg8eSKxEfQGuQEOl7pwksXgcIe1zBrDGdB88c5Yq8pI-4qLYPbujlJ1kIKHQK9BnI8eglR4aji0EYciowVr-RAuTuDtrYv4R8jrYq9iCGkXMim_JTUVTd8NwfA_OZ8&password=mystery41&csrf_token=AQBhQsEbdCaYYrUU7kEiMfr6gZ6PtCvYcA-BDo0YoEY9gWyJ55rB5kPwLQ0UmoAf3tWWvbbJsYubZrykEw"
// 'https://accounts.spotify.com/api/login'

// get token
//
// curl
// -H 'Host: open.spotify.com'
// -H 'Upgrade-Insecure-Requests: 1'
// -H 'DNT: 1'
// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36'
// -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3'
// -H 'Referer: https://accounts.spotify.com/en/login?continue=https:%2F%2Fopen.spotify.com%2Fbrowse%2Ffeatured'
// -H 'Accept-Language: en,ru;q=0.9'
// -H 'Cookie: sp_t=1e1b754f89af3a1e488c0653d57b02da; _ga=GA1.2.674780686.1555665773; _gid=GA1.2.697624719.1555665773; _gat=1; sp_dc=AQAHbZGDc9p80knfmxY7n0s3DQKX3Dr8DcdsUOy_MSj5CcLzdZmKS4BawhlYnbXFPBopZePX08sX1j9zjlnHDprbi5aRvZQqutQiU8r-Oho'
// 'https://open.spotify.com/browse/featured'
