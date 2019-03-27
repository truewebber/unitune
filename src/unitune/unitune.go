package main

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mgutz/logxi/v1"
	"github.com/mvdan/xurls"

	"lib/config"
	"lib/link-info"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.GetEnvParam("TOKEN"))
	if err != nil {
		log.Error("Error create new bot", "error", err.Error())

		return
	}

	log.Debug("Authorized", "_", bot.Self.UserName)

	u := tgbotapi.NewUpdate(20)
	u.Timeout = 5

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Error("Error get updates", "error", err.Error())
	}

	for u := range updates {
		if u.Message == nil {
			log.Debug("NOT MSG")

			continue
		}

		msg := u.Message

		// got message
		log.Debug("MSG", "chat", msg.Chat.Title, "from", msg.From.UserName, "_", msg.Text)

		// parse url
		links := xurls.Strict().FindAllString(msg.Text, -1)
		if len(links) == 0 {
			continue
		}

		// got links from msg
		log.Debug("LINKS", "_", links)

		// check link
		trackInfo, err := link_info.GetLinkInfo(links[0])
		if err != nil {
			if err == link_info.UnknownType {
				// it is not music link
				log.Debug("No music links")

				continue
			}

			// unknown error
			log.Error("Error get streamer", "error", err.Error())

			continue
		}

		// found music track
		log.Debug("FOUND TRACK LINK", "Actor", trackInfo.Actor(), "Albom", trackInfo.Albom(),
			"Title", trackInfo.Track())

		text := fmt.Sprintf("Found in %s!\nActor: %s\nAlbom: %s\nTitle: %s", trackInfo.StreamerType(),
			trackInfo.Actor(), trackInfo.Albom(), trackInfo.Track())

		replyMsg := tgbotapi.NewMessage(msg.Chat.ID, text)
		replyMsg.ReplyToMessageID = msg.MessageID

		_, err = bot.Send(replyMsg)
		if err != nil {
			log.Error("Error send msg to channel", "error", err.Error())
		}
	}
}
