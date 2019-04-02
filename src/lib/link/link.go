package link

import "regexp"

const (
	yandexMusicRegex = `^https?:\/\/music\.yandex\.ru\/album\/[0-9]+\/track\/[0-9]+$`
	spotifyLink      = `^https?:\/\/open\.spotify\.com\/track\/\w+\??.*`
	appleMusicLink   = ``
)

var (
	isYMRegex      = regexp.MustCompile(yandexMusicRegex)
	isSpotifyRegex = regexp.MustCompile(spotifyLink)
	//isAMRegex      = regexp.MustCompile(appleMusicLink)
)

func IsMusicLink(link string) bool {
	if IsYandexMusic(link) || IsAppleMusic(link) || IsSpotify(link) {
		return true
	}

	return false
}

// https://music.yandex.ru/album/5948145/track/30065762
func IsYandexMusic(link string) bool {
	return isYMRegex.MatchString(link)
}

// https://open.spotify.com/track/3F1TIgyzrfdU4dF8c4C75U?si=rLTMSukeTGeAJmFj8HdFLw
func IsSpotify(link string) bool {
	return isSpotifyRegex.MatchString(link)
}

// https://itunes.apple.com/ru/album/last-stop-before-heaven/1327950611?i=1327950740
func IsAppleMusic(link string) bool {
	return false
	//return isAMRegex.MatchString(link)
}
