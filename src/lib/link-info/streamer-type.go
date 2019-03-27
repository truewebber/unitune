package link_info

type (
	StreamerType int
)

const (
	StreamTypeSpotify     StreamerType = 1
	StreamTypeAppleMusic  StreamerType = 2
	StreamTypeYandexMusic StreamerType = 3

	StreamTypeSpotifyString     = "spotify"
	StreamTypeAppleMusicString  = "applemusic"
	StreamTypeYandexMusicString = "yandexmusic"
)

func (st StreamerType) String() string {
	switch st {
	case StreamTypeSpotify:
		return StreamTypeSpotifyString
	case StreamTypeAppleMusic:
		return StreamTypeAppleMusicString
	case StreamTypeYandexMusic:
		return StreamTypeYandexMusicString
	}

	return ""
}
