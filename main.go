package main

import (
	"firebot/internal/config"
	"firebot/internal/firefly"
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	configPath := flag.String("config", "./firebot.yml", "path to config file")
	flag.Parse()

	c, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	fclient := firefly.NewFireflyClient(c.Firefly.ApiUrl, c.Firefly.Token)

	bot, err := tgbotapi.NewBotAPI(c.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		cmdCfg := tgbotapi.NewSetMyCommands(
			tgbotapi.BotCommand{
				Command:     "balance",
				Description: "Баланс",
			},
		)
		bot.Send(cmdCfg)

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "balance":
			data, err := fclient.GetAccounts()
			if err != nil {
				log.Fatalln("Error on getting accounts: ", err)
			}
			msg.Text = data
			msg.ParseMode = "markdown"
		case "budgets":
			data, err := fclient.GetBudget()
			if err != nil {
				log.Fatalln("Error on getting budgets: ", err)
			}
			log.Println(data)
			msg.Text = data
			msg.ParseMode = "markdown"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
