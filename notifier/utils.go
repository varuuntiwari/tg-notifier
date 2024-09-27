package notifier

import (
	telBot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func sendTelegramText(config TelegramNotifier, msg string) (status error) {
	bot, status := telBot.NewBotAPI(config.BotToken)
	if status != nil {
		return
	}

	botMsg := telBot.NewMessage(config.UserID, msg)
	_, status = bot.Send(botMsg)

	if status != nil {
		return
	}

	return nil
}