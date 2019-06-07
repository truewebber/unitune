package main

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mgutz/logxi/v1"
	"github.com/mvdan/xurls"

	"github.com/truewebber/unitune/config"
	"github.com/truewebber/unitune/proxy"
	"github.com/truewebber/unitune/seeker"
	"github.com/truewebber/unitune/tune"
)

func main() {
	// bot
	bot, err := tgbotapi.NewBotAPI(config.EnvParam("TOKEN"))
	if err != nil {
		log.Error("Error create new bot", "error", err.Error())

		return
	}

	// bot enabled
	log.Debug("Authorized", "_", bot.Self.UserName)

	u := tgbotapi.NewUpdate(20)
	u.Timeout = 5

	// updates instance
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Error("Error get updates", "error", err.Error())
	}

	// get proxies from config
	proxies := make([]*proxy.Proxy, 0)
	err = config.Get().UnmarshalKey("proxies", &proxies)
	if err != nil {
		log.Error("Error parse proxies", "error", err.Error())
	}

	// get proxyList; add default non-proxy http client
	proxyList := proxy.GenerateProxyList(proxies)
	proxyList = append([]proxy.HttpProxyClient{proxy.NewNull()}, proxyList...)

	// init tunner
	tunner := tune.NewTunner(proxyList)

	// init master seeker
	ms := seeker.New(proxyList)

	// loop on updates
	for u := range updates {
		if u.Message == nil {
			log.Debug("NOT MSG")

			continue
		}

		msg := u.Message

		// got message
		log.Debug("MSG", "chat", msg.Chat.Title, "from", msg.From.UserName, "_", msg.Text)

		// parse url
		links := xurls.Strict.FindAllString(msg.Text, -1)
		if len(links) == 0 {
			continue
		}

		// got links from msg
		log.Debug("LINKS", "_", links)

		// check link
		t, err := tunner.Tune(links[0])
		if err != nil {
			if err == tune.UnknownType {
				// it is not music link
				log.Debug("No music links")

				continue
			}

			// unknown error
			log.Error("Error get tune", "error", err.Error())

			continue
		}

		if t == nil {
			continue
		}

		// found music track
		log.Debug("FOUND TRACK LINK", "Actor", t.Artist(), "Album", t.Album(), "Title", t.Track())

		links, errs := ms.LookUpTune(t)
		if len(errs) > 0 {
			for _, err := range errs {
				log.Error("Error lookup", "error", err.Error())
			}
		}

		text := fmt.Sprintf("Found in %s\n%s (%s - %s) - %s\n\n", t.StreamerType(), t.Track(), t.Album(),
			t.AlbumType(), t.Artist())
		for _, link := range links {
			text += fmt.Sprintf(" - %s\n", link)
		}

		replyMsg := tgbotapi.NewMessage(msg.Chat.ID, text)
		replyMsg.ReplyToMessageID = msg.MessageID

		_, err = bot.Send(replyMsg)
		if err != nil {
			log.Error("Error send msg to channel", "error", err.Error())
		}
	}
}
