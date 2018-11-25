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

		log.Debug("MSG", "chat", msg.Chat.Title, "from", msg.From.UserName, "_", msg.Text)

		links := xurls.Strict().FindAllString(msg.Text, -1)
		if len(links) == 0 {
			continue
		}

		log.Debug("LINKS", "_", links)

		streamer, err := link_info.GetLinkInfo(links[0])
		if err != nil && err != link_info.UnknownType {
			log.Error("Error get streamer", "error", err.Error())

			continue
		} else if err == link_info.UnknownType {
			log.Debug("No music links")

			continue
		}

		log.Debug("FOUND TRACK LINK", "Actor", streamer.GetActor(), "Albom", streamer.GetAlbom(),
			"Title", streamer.GetTrack())

		text := fmt.Sprintf("Found!\nActor: %s\nAlbom: %s\nTitle: %s", streamer.GetActor(),
			streamer.GetAlbom(), streamer.GetTrack())

		replyMsg := tgbotapi.NewMessage(msg.Chat.ID, text)
		replyMsg.ReplyToMessageID = msg.MessageID

		bot.Send(replyMsg)
	}
}
