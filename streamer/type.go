package streamer

type (
	Type int
)

const (
	TypeSpotify     Type = 1
	TypeAppleMusic  Type = 2
	TypeYandexMusic Type = 3

	TypeSpotifyString     = "spotify"
	TypeAppleMusicString  = "applemusic"
	TypeYandexMusicString = "yandexmusic"
)

func (t Type) String() string {
	switch t {
	case TypeSpotify:
		return TypeSpotifyString
	case TypeAppleMusic:
		return TypeAppleMusicString
	case TypeYandexMusic:
		return TypeYandexMusicString
	}

	return ""
}
