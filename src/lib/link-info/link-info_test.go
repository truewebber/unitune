package link_info

import "testing"

func TestGetLinkInfo(t *testing.T) {
	//info, err := GetLinkInfo("https://music.yandex.ru/album/5948145/track/30065762")
	info, err := GetLinkInfo("https://open.spotify.com/track/3F1TIgyzrfdU4dF8c4C75U?si=rLTMSukeTGeAJmFj8HdFLw")
	if err != nil {
		t.Error(err.Error())

		return
	}

	t.Log(info)
}
