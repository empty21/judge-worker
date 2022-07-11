package config

import (
	"github.com/Netflix/go-env"
	"github.com/denisbrodbeck/machineid"
	"github.com/joho/godotenv"
	"judger/pkg/logger"
)

type config struct {
	MachineId     string `env:"MACHINE_ID"`
	Runtime       string `env:"RUNTIME"`
	FileAPIKey    string `env:"FILE_API_KEY,required=true"`
	RedisUri      string `env:"REDIS_URI,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDatabase int    `env:"REDIS_DATABASE,default=10"`
	AMQPUri       string `env:"AMQP_URI,required=true"`
	TGBotToken    string `env:"TG_BOT_TOKEN"`
	TGChatId      int64  `env:"TG_CHAT_ID"`
}

var Config config

func init() {
	_ = godotenv.Load()
	_, err := env.UnmarshalFromEnviron(&Config)
	logger.Logger.Info("Runtime environment: ", Config.Runtime)

	if err != nil {
		logger.Logger.Fatal("Reading configuration from environment failed:", err)
	}
	if Config.MachineId == "" {
		id, err := machineid.ID()
		if err != nil {
			logger.Logger.Fatal("Can not read id of the machine:", err)
		}
		Config.MachineId = id
	}
}
