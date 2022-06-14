package main

import (
	"time"

	"github.com/go-ping/ping"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Checking resource availability via ICMP
func icmpChecker(update uint16, bot *tgbotapi.BotAPI, group int64, host string, count uint8, timeout uint8, timedelay uint16) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}
	pinger.Count = int(count)
	pinger.Timeout = 60 * time.Second
	for {
		err = pinger.Run()
		if err != nil {
			panic(err)
		}
		stats := pinger.Statistics()

		if stats.MaxRtt.Seconds() > float64(timeout) {
			msg := tgbotapi.NewMessage(group, "Host "+host+" ICMP error")
			bot.Send(msg)
		} else if stats.MaxRtt.Milliseconds() > int64(timedelay) {
			msg := tgbotapi.NewMessage(group, "Host "+host+" ICMP delay is "+stats.MaxRtt.String())
			bot.Send(msg)
		}
		time.Sleep(time.Duration(update) * time.Second)
	}
}
