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

const (
	Enter = `
`
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
		switch {
		case u.InlineQuery != nil:
			msg := u.InlineQuery

			// got message
			log.Debug("MSG", "chat", "inline", "from", msg.From.UserName, "_", msg.Query)

			out := tgbotapi.InlineConfig{
				InlineQueryID: u.InlineQuery.ID,
				IsPersonal:    true,
			}

			t, links, err := inlineQuery(tunner, ms, msg.Query)
			if t == nil {
				if err != nil {
					log.Error("Error inline query", "query", msg.Query, "error", err.Error())
				}

				_, err = bot.AnswerInlineQuery(out)
				if err != nil {
					log.Error("Error send inline query response", "error", err.Error())
				}

				continue
			}

			text := fmt.Sprintf("Found in *%s*  "+Enter+"%s (%s) - %s  "+Enter,
				t.StreamerType(), t.Track(), t.Album(), t.Artist())
			if len(links) > 0 {
				for _, link := range links {
					text += fmt.Sprintf(" - [%s](%s)  "+Enter, link.StreamerType.String(), link.Link)
				}
			}

			results := []interface{}{
				tgbotapi.InlineQueryResultArticle{
					Type:  "article",
					ID:    "1",
					Title: t.Track(),
					InputMessageContent: tgbotapi.InputTextMessageContent{
						Text:                  text,
						ParseMode:             "Markdown",
						DisableWebPagePreview: false,
					},
					ThumbURL:    t.AlbumPic(),
					ThumbHeight: 50,
					ThumbWidth:  50,
				},
			}

			out.Results = results
			_, err = bot.AnswerInlineQuery(out)
			if err != nil {
				log.Error("Error send inline query response", "error", err.Error())

				continue
			}
		case u.Message != nil:
			msg := u.Message

			if !msg.IsCommand() {
				continue
			}

			switch msg.Command() {
			case "start":
				if _, err := bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Hello!")); err != nil {
					log.Error("Error send reply", "error", err.Error())
				}
			case "stop":
				if _, err := bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ok. Bye!")); err != nil {
					log.Error("Error send reply", "error", err.Error())
				}
			case "help":
				text := "Use: @unitunebot [link]\nFor now, bot supports Spotify, YandexMusic, AppleMusic(soon)\n\n" +
					"Commands:\nhelp - how to use this bot\ncontact - developer info"
				if _, err := bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text)); err != nil {
					log.Error("Error send reply", "error", err.Error())
				}
			case "contact":
				text := "Creator: @truewebber\nGithub: github.com/truewebber/unitune"
				if _, err := bot.Send(tgbotapi.NewMessage(msg.Chat.ID, text)); err != nil {
					log.Error("Error send reply", "error", err.Error())
				}
			default:
			}
		default:
			log.Debug("NOT MSG")
		}
	}
}

func inlineQuery(tunner *tune.Tunner, ms *seeker.MasterSeeker, q string) (tune.Tune, []seeker.Tune, error) {
	// parse url
	links := xurls.Strict.FindAllString(q, -1)
	if len(links) == 0 {
		return nil, nil, nil
	}

	// got links from msg
	log.Debug("LINKS", "_", links)

	// check link
	t, err := tunner.Tune(links[0])
	if err != nil {
		if err == tune.UnknownType {
			// it is not music link
			log.Debug("No music links")

			return nil, nil, nil
		}

		// unknown error
		log.Error("Error get tune", "error", err.Error())

		return nil, nil, err
	}

	if t == nil {
		return nil, nil, nil
	}

	// found music track
	log.Debug("FOUND TRACK LINK", "Actor", t.Artist(), "Album", t.Album(), "Title", t.Track())

	out, errs := ms.LookUpTune(t)
	if len(errs) > 0 {
		for _, err := range errs {
			log.Error("Error lookup", "error", err.Error())
		}
	}

	return t, out, nil
}
