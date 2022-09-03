package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Update uint16
	}
	Telegram struct {
		Token string
		Group int64
	}
	Http struct {
		Repeat  uint8
		Timeout uint8
		Delay   float64
		Sites   []struct {
			Url      string
			Elements []string
		}
	}
}

func main() {
	// Read config from yaml
	config := Config{}
	filename, _ := filepath.Abs("./conf/config.yaml")
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Parse yaml
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Telegram bot
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	// Running HTTP checker
	for _, site := range config.Http.Sites {
		go httpCheck(config.App.Update, bot, config.Telegram.Group, site, config.Http.Timeout, config.Http.Repeat, config.Http.Delay)
	}

	botUpdate(bot, config.Http.Sites)
}

// Telegram bot for listening to incoming commands
func botUpdate(bot *tgbotapi.BotAPI, sites []struct {
	Url      string
	Elements []string
}) {

	// Create string for HTTP(s) monitoring sites
	sitesString := ""
	for _, site := range sites {
		sitesString += site.Url + "\n"
	}

	// Telegram bot listener
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 300
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = "Hi, I am a monitoring bot! Your (group) ID = " + strconv.FormatInt(update.Message.Chat.ID, 10)
		case "list":
			msg.Text = "HTTP(s) monitoring sites:\n" + sitesString
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
