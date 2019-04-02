package seeker

import (
	"lib/streamer"
	"lib/tune"
)

type (
	spotify struct {
	}
)

func newSpotify() *spotify {
	return &spotify{}
}

func (s *spotify) Seek(tune tune.Tune) (*string, error) {
	// https://api.spotify.com/v1/search?
	// type=album,artist,playlist,track,show_audio,episode_audio
	// q=armin blah bla*
	// decorate_restrictions=true
	// best_match=true
	// include_external=audio
	// limit=50
	// userless=false
	// market=from_token

	// Accept: application/json
	// Referer: https://open.spotify.com/search/results/armin%20blah%20bla
	// Origin: https://open.spotify.com
	// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36
	// Accept-Language: en
	// Authorization: Bearer BQB_QBiWrGeFoo9GTAqbuNDOvBIpZyA1JFC1eKYJ9zFiYRk6NmiOzDCXf7fxYSbUiaoFysCEZTsNkiRU8KrzA0YR7KNsFHZCl_9kUJ_9DyuP4S7iafdLvBEX7xuNaTe69cAmOxi5SPbJ-gfG0yWpUrvP05gue38AQ71K2XjQV7gYWxlX3_9e6t0jqa69xUFZEOOffe4Su_cjdORpgQWTc-D_-8L2YqMnXLGPvIgvW_X41FsUsqNLwFleFSsHE5MWmBmIYoIr2o2UVIMlKpm2IMXpO12D7sl7
	// DNT: 1

	return nil, nil
}

func (s *spotify) StreamerType() streamer.Type {
	return streamer.TypeSpotify
}
