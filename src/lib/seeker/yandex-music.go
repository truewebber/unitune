package seeker

import (
	"lib/link-info"
	"lib/streamer"
)

// curl "https://music.yandex.ru/handlers/music-search.jsx?text=retrovision+get+up&type=all&ncrnd=0.36734637463&lang=ru&external-domain=music.yandex.ru&overembed=false"

type (
	YandexMusic struct {
	}
)

func NewYandexMusic() *YandexMusic {
	return &YandexMusic{}
}

func (y *YandexMusic) Seek(tune link_info.Tune) (string, error) {
	return "", nil
}

func (y *YandexMusic) StreamerType() streamer.Type {
	return streamer.TypeYandexMusic
}
