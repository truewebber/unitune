package link_info

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

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

	YMMetadataResponse struct {
		Experiments struct {
			Ugc                             string `json:"ugc"`
			SearchPrioritizeOwnContent      string `json:"searchPrioritizeOwnContent"`
			RotorSampleCandidates           string `json:"rotorSampleCandidates"`
			Musicgift                       string `json:"musicgift"`
			WebNewPlaylistsTabHide          string `json:"webNewPlaylistsTabHide"`
			PromoteOnYourWave               string `json:"promoteOnYourWave"`
			Adv                             string `json:"adv"`
			AudioAdsWhite                   string `json:"audioAdsWhite"`
			RotorNewSettings                string `json:"rotorNewSettings"`
			DjvuCandidates                  string `json:"djvuCandidates"`
			MusicCheckPass                  string `json:"musicCheckPass"`
			FeedTriggers                    string `json:"feedTriggers"`
			MusicPlaylistNoTrack            string `json:"musicPlaylistNoTrack"`
			MarketingMail100                string `json:"marketingMail_100"`
			MusicSyncQueue                  string `json:"musicSyncQueue"`
			MusicMobileWebLocked            string `json:"musicMobileWebLocked"`
			RotorSafeTopSize                string `json:"rotorSafeTopSize"`
			MusicCollectivePlaylist         string `json:"musicCollectivePlaylist"`
			WebSimilarArtistsInHead         string `json:"webSimilarArtistsInHead"`
			BranchLinks                     string `json:"branchLinks"`
			MusicNewGenres                  string `json:"musicNewGenres"`
			MusicVideoOnArtistPage          string `json:"musicVideoOnArtistPage"`
			WebAntiMusicTab                 string `json:"webAntiMusicTab"`
			MusicTestDebugProducts          string `json:"musicTestDebugProducts"`
			TurboQuasar                     string `json:"turboQuasar"`
			MusicLandingIntentPlaylistCount string `json:"musicLandingIntentPlaylistCount"`
			SearchQuorumSoftness            string `json:"searchQuorumSoftness"`
			WebPlaylistOfTheDayCounter      string `json:"webPlaylistOfTheDayCounter"`
			WebAutoPlaylistAnimated         string `json:"webAutoPlaylistAnimated"`
			MarketingMail                   string `json:"marketingMail"`
			RotorSimilarDither              string `json:"rotorSimilarDither"`
			WebMusicPreroll                 string `json:"webMusicPreroll"`
			WebSidebarUPbanner              string `json:"webSidebarUPbanner"`
			MusicLoginWall                  string `json:"musicLoginWall"`
			MusicSearchFormula              string `json:"musicSearchFormula"`
			MusicExperimentalPlayer         string `json:"musicExperimentalPlayer"`
			UserFeed                        string `json:"userFeed"`
			MusicUzbekistanLang             string `json:"musicUzbekistanLang"`
			MusicAdsVolumeIncrease          string `json:"musicAdsVolumeIncrease"`
			UseUserGroups                   string `json:"useUserGroups"`
			ReferralCounter                 string `json:"referralCounter"`
			MusicHighlightLyrics            string `json:"musicHighlightLyrics"`
			MusicErrorLogger                string `json:"musicErrorLogger"`
			SubStation                      string `json:"SubStation"`
			MusicHebrewLang                 string `json:"musicHebrewLang"`
			RotorRandomStation              string `json:"rotorRandomStation"`
			AutoPlaylistsAnimated           string `json:"autoPlaylistsAnimated"`
			WebAntiMusicBlockNaGlavnoi      string `json:"webAntiMusicBlockNaGlavnoi"`
			MusicTouchNewPleer              string `json:"musicTouchNewPleer"`
			MusicLoginBar                   string `json:"musicLoginBar"`
			MusicCspLogger                  string `json:"musicCspLogger"`
			MarketingMail10                 string `json:"marketingMail_10"`
			UgcPrivat                       string `json:"ugcPrivat"`
			MusicNYChart                    string `json:"musicNYChart"`
			PlusWebSale                     string `json:"plusWebSale"`
			WebPodcastShowUpdated           string `json:"webPodcastShowUpdated"`
			MusicStatsLogger                string `json:"musicStatsLogger"`
			MusicTakeEMail                  string `json:"musicTakeEMail"`
			MarketingMail95                 string `json:"marketingMail_95"`
			RotorLoginWall                  string `json:"rotorLoginWall"`
			MarketingMail25                 string `json:"marketingMail_25"`
			MiniBrick                       string `json:"miniBrick"`
			MusicYellowButton               string `json:"musicYellowButton"`
			MusicChartSwitch                string `json:"musicChartSwitch"`
			MarketingMail5                  string `json:"marketingMail_5"`
			MusicAutoFlow                   string `json:"musicAutoFlow"`
			PlusWeb                         string `json:"plusWeb"`
			MusicYellowButtonAuth           string `json:"musicYellowButtonAuth"`
			MusicPrice                      string `json:"musicPrice"`
			MusicKazakhstanLang             string `json:"musicKazakhstanLang"`
			MusicArtistStat                 string `json:"musicArtistStat"`
			RotorIosHappyNewYearDesign      string `json:"rotorIosHappyNewYearDesign"`
			RotorSafeTopTracksCount         string `json:"rotorSafeTopTracksCount"`
			Strm                            string `json:"strm"`
			MusicArmeniaLang                string `json:"musicArmeniaLang"`
			MusicSuggest                    string `json:"musicSuggest"`
			WebAutoplaylistsOnMain          string `json:"webAutoplaylistsOnMain"`
			CryProxy                        string `json:"cryProxy"`
		} `json:"experiments"`
		Hostname string `json:"hostname"`
		AuthData struct {
			User struct {
				Sign         string `json:"sign"`
				Sk           string `json:"sk"`
				DeviceID     string `json:"device_id"`
				OnlyDeviceID bool   `json:"onlyDeviceId"`
				IsHosted     bool   `json:"isHosted"`
				Name         string `json:"name"`
				Login        string `json:"login"`
				UID          int    `json:"uid"`
				HasEmail     bool   `json:"hasEmail"`
			} `json:"user"`
			Experiments string `json:"experiments"`
		} `json:"authData"`
		ServiceName string `json:"serviceName"`
		Settings    struct {
			Module       string `json:"module"`
			Date         int64  `json:"date"`
			Pathname     string `json:"pathname"`
			Tld          string `json:"tld"`
			MdaEnabled   bool   `json:"mdaEnabled"`
			IsDev        bool   `json:"isDev"`
			IsYandex     bool   `json:"isYandex"`
			IsCIS        bool   `json:"isCIS"`
			Compat       int    `json:"compat"`
			Lang         string `json:"lang"`
			BadRegion    bool   `json:"badRegion"`
			Region       int    `json:"region"`
			NativeRegion int    `json:"nativeRegion"`
			PortalRegion int    `json:"portalRegion"`
			Storage      string `json:"storage"`
			Avatars      string `json:"avatars"`
			Social       string `json:"social"`
			Passport     string `json:"passport"`
			PassportAPI  string `json:"passportApi"`
			Socket       string `json:"socket"`
			Regions      struct {
				Native  int         `json:"native"`
				Working int         `json:"working"`
				Show    int         `json:"show"`
				Cookie  interface{} `json:"cookie"`
				Portal  int         `json:"portal"`
			} `json:"regions"`
			Uatraits struct {
				IsTouch              bool   `json:"isTouch"`
				IsMobile             bool   `json:"isMobile"`
				PostMessageSupport   bool   `json:"postMessageSupport"`
				IsBrowser            bool   `json:"isBrowser"`
				HistorySupport       bool   `json:"historySupport"`
				WebPSupport          bool   `json:"WebPSupport"`
				SVGSupport           bool   `json:"SVGSupport"`
				BrowserBaseVersion   string `json:"BrowserBaseVersion"`
				BrowserEngine        string `json:"BrowserEngine"`
				OSFamily             string `json:"OSFamily"`
				BrowserEngineVersion string `json:"BrowserEngineVersion"`
				BrowserVersion       string `json:"BrowserVersion"`
				BrowserName          string `json:"BrowserName"`
				CSP1Support          bool   `json:"CSP1Support"`
				LocalStorageSupport  bool   `json:"localStorageSupport"`
				BrowserBase          string `json:"BrowserBase"`
				CSP2Support          bool   `json:"CSP2Support"`
				OSVersion            string `json:"OSVersion"`
			} `json:"uatraits"`
			Sandbox   string `json:"sandbox"`
			RequestID string `json:"requestId"`
			Csp       struct {
				AllowedHosts []string `json:"allowedHosts"`
				ReportURL    string   `json:"reportUrl"`
			} `json:"csp"`
			Theme string `json:"theme"`
			Bilet struct {
				URL             string `json:"url"`
				ClientKey       string `json:"clientKey"`
				MobileClientKey string `json:"mobileClientKey"`
			} `json:"bilet"`
			ServerSide bool `json:"serverSide"`
		} `json:"settings"`
		LibraryData struct {
			Value struct {
				Success bool   `json:"success"`
				Reason  string `json:"reason"`
			} `json:"value"`
		} `json:"libraryData"`
		Genres struct {
			Structure []struct {
				ID        string `json:"id"`
				SubGenres []struct {
					ID            string `json:"id"`
					HideInRegions []int  `json:"hideInRegions,omitempty"`
				} `json:"subGenres,omitempty"`
				HideInRegions []int `json:"hideInRegions,omitempty"`
			} `json:"structure"`
			Titles struct {
				All              string `json:"all"`
				Pop              string `json:"pop"`
				Indie            string `json:"indie"`
				Rock             string `json:"rock"`
				Metal            string `json:"metal"`
				Alternative      string `json:"alternative"`
				Electronics      string `json:"electronics"`
				Electronic       string `json:"electronic"`
				Dance            string `json:"dance"`
				Rap              string `json:"rap"`
				HipHop           string `json:"hip-hop"`
				Rnb              string `json:"rnb"`
				RNB              string `json:"r-n-b"`
				Jazz             string `json:"jazz"`
				Blues            string `json:"blues"`
				Reggae           string `json:"reggae"`
				Ska              string `json:"ska"`
				Punk             string `json:"punk"`
				Folk             string `json:"folk"`
				World            string `json:"world"`
				Classical        string `json:"classical"`
				Estrada          string `json:"estrada"`
				Shanson          string `json:"shanson"`
				Country          string `json:"country"`
				Soundtrack       string `json:"soundtrack"`
				Relax            string `json:"relax"`
				Easy             string `json:"easy"`
				Bard             string `json:"bard"`
				SingerSongwriter string `json:"singer-songwriter"`
				Forchildren      string `json:"forchildren"`
				ForChildren      string `json:"for-children"`
				Fairytales       string `json:"fairytales"`
				Other            string `json:"other"`
				Ruspop           string `json:"ruspop"`
				Disco            string `json:"disco"`
				Kpop             string `json:"kpop"`
				LocalIndie       string `json:"local-indie"`
				Rusrock          string `json:"rusrock"`
				Rnr              string `json:"rnr"`
				RockNRoll        string `json:"rock-n-roll"`
				Prog             string `json:"prog"`
				ProgRock         string `json:"prog-rock"`
				Postrock         string `json:"postrock"`
				PostRock         string `json:"post-rock"`
				Newwave          string `json:"newwave"`
				NewWave          string `json:"new-wave"`
				Ukrrock          string `json:"ukrrock"`
				Folkrock         string `json:"folkrock"`
				Stonerrock       string `json:"stonerrock"`
				Hardrock         string `json:"hardrock"`
				Classicmetal     string `json:"classicmetal"`
				Progmetal        string `json:"progmetal"`
				Numetal          string `json:"numetal"`
				Epicmetal        string `json:"epicmetal"`
				Folkmetal        string `json:"folkmetal"`
				Extrememetal     string `json:"extrememetal"`
				Industrial       string `json:"industrial"`
				Posthardcore     string `json:"posthardcore"`
				Hardcore         string `json:"hardcore"`
				Dubstep          string `json:"dubstep"`
				Experimental     string `json:"experimental"`
				House            string `json:"house"`
				Techno           string `json:"techno"`
				Trance           string `json:"trance"`
				Dnb              string `json:"dnb"`
				DrumNBass        string `json:"drum-n-bass"`
				Rusrap           string `json:"rusrap"`
				Foreignrap       string `json:"foreignrap"`
				Soul             string `json:"soul"`
				Funk             string `json:"funk"`
				Tradjazz         string `json:"tradjazz"`
				TradJass         string `json:"trad-jass"`
				Conjazz          string `json:"conjazz"`
				ModernJazz       string `json:"modern-jazz"`
				Reggaeton        string `json:"reggaeton"`
				Dub              string `json:"dub"`
				Rusfolk          string `json:"rusfolk"`
				Russian          string `json:"russian"`
				Tatar            string `json:"tatar"`
				Caucasian        string `json:"caucasian"`
				Celtic           string `json:"celtic"`
				Balkan           string `json:"balkan"`
				Eurofolk         string `json:"eurofolk"`
				European         string `json:"european"`
				Jewish           string `json:"jewish"`
				Eastern          string `json:"eastern"`
				African          string `json:"african"`
				Latinfolk        string `json:"latinfolk"`
				LatinAmerican    string `json:"latin-american"`
				Amerfolk         string `json:"amerfolk"`
				American         string `json:"american"`
				Romances         string `json:"romances"`
				Argentinetango   string `json:"argentinetango"`
				Vocal            string `json:"vocal"`
				Opera            string `json:"opera"`
				Modern           string `json:"modern"`
				ModernClassical  string `json:"modern-classical"`
				Rusestrada       string `json:"rusestrada"`
				Films            string `json:"films"`
				Tvseries         string `json:"tvseries"`
				TvSeries         string `json:"tv-series"`
				Animated         string `json:"animated"`
				AnimatedFilms    string `json:"animated-films"`
				Videogame        string `json:"videogame"`
				VideogameMusic   string `json:"videogame-music"`
				Musical          string `json:"musical"`
				Bollywood        string `json:"bollywood"`
				Lounge           string `json:"lounge"`
				Newage           string `json:"newage"`
				NewAge           string `json:"new-age"`
				Meditation       string `json:"meditation"`
				Meditative       string `json:"meditative"`
				Rusbards         string `json:"rusbards"`
				Foreignbard      string `json:"foreignbard"`
				Sport            string `json:"sport"`
				Holiday          string `json:"holiday"`
				Spoken           string `json:"spoken"`
				AudioBooks       string `json:"audio-books"`
			} `json:"titles"`
			CustomTitles struct {
				All            string `json:"all"`
				Dance          string `json:"dance"`
				Classical      string `json:"classical"`
				Forchildren    string `json:"forchildren"`
				ForChildren    string `json:"for-children"`
				Other          string `json:"other"`
				Ruspop         string `json:"ruspop"`
				Kpop           string `json:"kpop"`
				LocalIndie     string `json:"local-indie"`
				Folkrock       string `json:"folkrock"`
				Stonerrock     string `json:"stonerrock"`
				Hardrock       string `json:"hardrock"`
				Classicmetal   string `json:"classicmetal"`
				Progmetal      string `json:"progmetal"`
				Numetal        string `json:"numetal"`
				Epicmetal      string `json:"epicmetal"`
				Folkmetal      string `json:"folkmetal"`
				Extrememetal   string `json:"extrememetal"`
				Posthardcore   string `json:"posthardcore"`
				Hardcore       string `json:"hardcore"`
				Experimental   string `json:"experimental"`
				Foreignrap     string `json:"foreignrap"`
				Tradjazz       string `json:"tradjazz"`
				TradJass       string `json:"trad-jass"`
				Conjazz        string `json:"conjazz"`
				ModernJazz     string `json:"modern-jazz"`
				Rusfolk        string `json:"rusfolk"`
				Russian        string `json:"russian"`
				Tatar          string `json:"tatar"`
				Caucasian      string `json:"caucasian"`
				Celtic         string `json:"celtic"`
				Balkan         string `json:"balkan"`
				Eurofolk       string `json:"eurofolk"`
				European       string `json:"european"`
				Jewish         string `json:"jewish"`
				Eastern        string `json:"eastern"`
				African        string `json:"african"`
				Latinfolk      string `json:"latinfolk"`
				LatinAmerican  string `json:"latin-american"`
				Amerfolk       string `json:"amerfolk"`
				American       string `json:"american"`
				Romances       string `json:"romances"`
				Argentinetango string `json:"argentinetango"`
				Vocal          string `json:"vocal"`
				Opera          string `json:"opera"`
				Rusestrada     string `json:"rusestrada"`
				Films          string `json:"films"`
				Tvseries       string `json:"tvseries"`
				TvSeries       string `json:"tv-series"`
				Animated       string `json:"animated"`
				AnimatedFilms  string `json:"animated-films"`
				Videogame      string `json:"videogame"`
				VideogameMusic string `json:"videogame-music"`
				Meditation     string `json:"meditation"`
				Meditative     string `json:"meditative"`
				Rusbards       string `json:"rusbards"`
				Foreignbard    string `json:"foreignbard"`
				Sport          string `json:"sport"`
				Holiday        string `json:"holiday"`
			} `json:"customTitles"`
			CustomUrls struct {
				Electronics  string `json:"electronics"`
				Rap          string `json:"rap"`
				Rnb          string `json:"rnb"`
				Folk         string `json:"folk"`
				Relax        string `json:"relax"`
				Bard         string `json:"bard"`
				Forchildren  string `json:"forchildren"`
				Kpop         string `json:"kpop"`
				Rnr          string `json:"rnr"`
				Prog         string `json:"prog"`
				Postrock     string `json:"postrock"`
				Newwave      string `json:"newwave"`
				Hardrock     string `json:"hardrock"`
				Numetal      string `json:"numetal"`
				Posthardcore string `json:"posthardcore"`
				Dnb          string `json:"dnb"`
				Foreignrap   string `json:"foreignrap"`
				Tradjazz     string `json:"tradjazz"`
				Conjazz      string `json:"conjazz"`
				Rusfolk      string `json:"rusfolk"`
				Eurofolk     string `json:"eurofolk"`
				Latinfolk    string `json:"latinfolk"`
				Amerfolk     string `json:"amerfolk"`
				Vocal        string `json:"vocal"`
				Modern       string `json:"modern"`
				Tvseries     string `json:"tvseries"`
				Animated     string `json:"animated"`
				Videogame    string `json:"videogame"`
				Newage       string `json:"newage"`
				Meditation   string `json:"meditation"`
				Spoken       string `json:"spoken"`
			} `json:"customUrls"`
		} `json:"genres"`
		PageData struct {
			ID            int           `json:"id"`
			Title         string        `json:"title"`
			Type          string        `json:"type"`
			Year          int           `json:"year"`
			ReleaseDate   time.Time     `json:"releaseDate"`
			CoverURI      string        `json:"coverUri"`
			OgImage       string        `json:"ogImage"`
			Genre         string        `json:"genre"`
			Buy           []interface{} `json:"buy"`
			TrackCount    int           `json:"trackCount"`
			Recent        bool          `json:"recent"`
			VeryImportant bool          `json:"veryImportant"`
			Artists       []struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Various  bool   `json:"various"`
				Composer bool   `json:"composer"`
				Cover    struct {
					Type   string `json:"type"`
					Prefix string `json:"prefix"`
					URI    string `json:"uri"`
				} `json:"cover"`
				Genres []interface{} `json:"genres"`
			} `json:"artists"`
			Labels []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"labels"`
			Available                bool          `json:"available"`
			AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
			AvailableForMobile       bool          `json:"availableForMobile"`
			AvailablePartially       bool          `json:"availablePartially"`
			Bests                    []interface{} `json:"bests"`
			Prerolls                 []interface{} `json:"prerolls"`
			Volumes                  [][]struct {
				ID     string `json:"id"`
				RealID string `json:"realId"`
				Title  string `json:"title"`
				Major  struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"major"`
				Available                bool   `json:"available"`
				AvailableForPremiumUsers bool   `json:"availableForPremiumUsers"`
				DurationMs               int    `json:"durationMs"`
				StorageDir               string `json:"storageDir"`
				FileSize                 int    `json:"fileSize"`
				Normalization            struct {
					Gain float64 `json:"gain"`
					Peak int     `json:"peak"`
				} `json:"normalization"`
				Artists []struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Various  bool   `json:"various"`
					Composer bool   `json:"composer"`
					Cover    struct {
						Type   string `json:"type"`
						Prefix string `json:"prefix"`
						URI    string `json:"uri"`
					} `json:"cover"`
					Genres []interface{} `json:"genres"`
				} `json:"artists"`
				Albums []struct {
					Redirected               bool          `json:"redirected"`
					Volumes                  interface{}   `json:"volumes"`
					Prerolls                 []interface{} `json:"prerolls"`
					Bests                    []interface{} `json:"bests"`
					AvailablePartially       bool          `json:"availablePartially"`
					AvailableForMobile       bool          `json:"availableForMobile"`
					AvailableForPremiumUsers bool          `json:"availableForPremiumUsers"`
					Available                bool          `json:"available"`
					Labels                   []struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"labels"`
					Artists []struct {
						ID       int    `json:"id"`
						Name     string `json:"name"`
						Various  bool   `json:"various"`
						Composer bool   `json:"composer"`
						Cover    struct {
							Type   string `json:"type"`
							Prefix string `json:"prefix"`
							URI    string `json:"uri"`
						} `json:"cover"`
						Genres []interface{} `json:"genres"`
					} `json:"artists"`
					VeryImportant bool          `json:"veryImportant"`
					Recent        bool          `json:"recent"`
					TrackCount    int           `json:"trackCount"`
					Buy           []interface{} `json:"buy"`
					Genre         string        `json:"genre"`
					OgImage       string        `json:"ogImage"`
					CoverURI      string        `json:"coverUri"`
					ReleaseDate   time.Time     `json:"releaseDate"`
					Year          int           `json:"year"`
					Type          string        `json:"type"`
					Title         string        `json:"title"`
					ID            int           `json:"id"`
				} `json:"albums"`
				CoverURI        string `json:"coverUri"`
				OgImage         string `json:"ogImage"`
				LyricsAvailable bool   `json:"lyricsAvailable"`
				Best            bool   `json:"best"`
			} `json:"volumes"`
			Redirected bool `json:"redirected"`
		} `json:"pageData"`
		PageName string `json:"pageName"`
	}

	YandexMusic struct {
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
	YMMetadataRegex = regexp.MustCompile("<script\\s+nonce=\"[\\w/+=]+\">\\s*var\\s+Mu=(.*});\\s*</script>")
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
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

	results := YMMetadataRegex.FindSubmatch(data)
	if len(results) != 2 {
		return nil, errors.Errorf("Error parse YMMetadata payload, body: %s", string(data))
	}

	obj2 := new(YMMetadataResponse)
	err = json.Unmarshal(results[1], obj2)
	if err != nil {
		return nil, errors.Errorf("Error unmarshal YMMetadata json payload, error: %s, data: %s", err.Error(),
			string(results[1]))
	}

	return &YandexMusic{
		trackLink:  link,
		ActorTitle: obj.ByArtist.Name,
		AlbomTitle: obj.InAlbum.Name,
		AlbomType:  obj2.PageData.Type,
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

func (y *YandexMusic) GetAlbomType() string {
	return y.AlbomType
}

func (y *YandexMusic) GetTrack() string {
	return y.TrackTitle
}
