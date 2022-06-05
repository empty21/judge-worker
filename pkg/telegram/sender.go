package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"judger/pkg/config"
	"judger/pkg/logger"
)

var bot *tgbotapi.BotAPI

func SendFile(filePath string) {
	_, err := bot.Send(tgbotapi.NewDocument(config.Config.TGChatId, tgbotapi.FilePath(filePath)))
	if err != nil {
		logger.Logger.Error(err)
	}
}

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(config.Config.TGBotToken)
	if err != nil {
		logger.Logger.Error(err)
	}
}
