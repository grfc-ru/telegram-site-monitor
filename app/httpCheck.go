package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Checking the availability of the site via the HTTP protocol
func httpCheck(update uint16, bot *tgbotapi.BotAPI, group int64, site struct {
	Url      string
	Elements []string
}, timeout uint8, repeat uint8) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	for {
		errorHTML := 0
		deface := false
		for i := 0; i < int(repeat); i++ {

			resp, err := client.Get(site.Url)
			if err != nil {
				msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP get error")
				bot.Send(msg)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP get error")
				bot.Send(msg)
			}
			body := string(bodyBytes)

			if resp.StatusCode != 200 {
				msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP error. Code "+strconv.Itoa(resp.StatusCode))
				bot.Send(msg)
				break
			}

			for _, element := range site.Elements {
				if !strings.Contains(body, element) {
					msg := tgbotapi.NewMessage(group, "Site "+site.Url+" defaced. Element '"+element+"' not found.")
					bot.Send(msg)
					deface = true
				}
			}
			if deface {
				break
			}
		}
		if errorHTML >= int(repeat-1) {
			msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP get error")
			bot.Send(msg)
		}
		time.Sleep(time.Duration(update) * time.Second)
	}
}
