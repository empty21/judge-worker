package config

import "github.com/jinzhu/configor"

type _Config struct {
	AMQPUri          string `env:"AMQP_URI" yaml:"AMQPUri" required:"true"`
	Debug            bool   `env:"DEBUG" yaml:"Debug" default:"false"`
	LogLevel         string `env:"LOG_LEVEL" yaml:"LogLevel" default:"info"`
	Sandbox          string `env:"SANDBOX" yaml:"Sandbox" default:"docker"`
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN" yaml:"TelegramBotToken"`
	TelegramChatID   int64  `env:"TELEGRAM_CHAT_ID" yaml:"TelegramChatID"`
}

var Config _Config

func init() {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		panic(err)
	}
}
