package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"judger/pkg/logger"
)

type config struct {
	Runtime       string `env:"RUNTIME"`
	RedisUri      string `env:"REDIS_URI,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDatabase string `env:"REDIS_DATABASE,default=0"`
	AMQPUri       string `env:"AMQP_URI,required=true"`
	TGBotToken    string `env:"TG_BOT_TOKEN"`
	TGChatId      int64  `env:"TG_CHAT_ID"`
}

var Config config

func init() {
	_ = godotenv.Load()

	_, err := env.UnmarshalFromEnviron(&Config)
	if err != nil {
		logger.Logger.Fatal("Reading configuration from environment failed:", err)
	}
}
