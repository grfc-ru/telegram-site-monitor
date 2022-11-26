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
}, timeout uint8, repeat uint8, delay float64) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	for {
		errorHTML := 0
		deface := false
		for i := 0; i < int(repeat); i++ {

			start := time.Now()
			resp, err := client.Get(site.Url)
			elapsed := time.Since(start).Seconds()

			if err != nil {
				errorHTML++
			} else {
				if resp.StatusCode != 200 {
					msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP error. Code "+strconv.Itoa(resp.StatusCode))
					bot.Send(msg)
					break
				}

				if elapsed >= delay {
					msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP delay "+strconv.FormatFloat(elapsed, 'f', 3, 32)+" sec.")
					bot.Send(msg)
					break
				}

				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					errorHTML++
				} else {
					body := string(bodyBytes)

					for _, element := range site.Elements {
						if !strings.Contains(body, element) {
							msg := tgbotapi.NewMessage(group, "Site "+site.Url+" defaced. Element '"+element+"' not found.")
							bot.Send(msg)
							deface = true
							break
						}
					}
					if deface {
						break
					}
				}
			}
		}
		if errorHTML >= int(repeat) {
			msg := tgbotapi.NewMessage(group, "Site "+site.Url+" HTTP get error")
			bot.Send(msg)
			errorHTML = 0
		}
		time.Sleep(time.Duration(update) * time.Second)
	}
}
