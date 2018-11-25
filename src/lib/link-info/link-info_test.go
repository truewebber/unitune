package link_info

import "testing"

func TestGetLinkInfo(t *testing.T) {
	info, err := GetLinkInfo("https://music.yandex.ru/album/5948145/track/30065762")
	if err != nil {
		t.Error(err.Error())

		return
	}

	t.Log(info)
}
