package config

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"judger/pkg/logger"
)

type config struct {
	Runtime     string `env:"RUNTIME"`
	FileAPIKey  string `env:"FILE_API_KEY,required=true"`
	AMQPUri     string `env:"AMQP_URI,required=true"`
	TGBotToken  string `env:"TG_BOT_TOKEN"`
	TGChatId    int64  `env:"TG_CHAT_ID"`
	JobQueue    string `env:"JOB_QUEUE,required=true"`
	ResultQueue string `env:"RESULT_QUEUE,required=true"`
}

var Config config

func init() {
	_ = godotenv.Load()
	_, err := env.UnmarshalFromEnviron(&Config)
	logger.Logger.Info("Runtime environment: ", Config.Runtime)

	if err != nil {
		logger.Logger.Fatal("Reading configuration from environment failed:", err)
	}
}
