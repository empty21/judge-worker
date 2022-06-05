package logger

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

var Logger *zap.SugaredLogger
var AMQPLogger *zap.SugaredLogger

func init() {
	Logger = createLogger("CommonLogger", zap.DebugLevel, fmt.Sprintf("logs/log-%s.log", time.Now().Format("20060102")))
	AMQPLogger = createLogger("AMQPLogger", zap.DebugLevel, fmt.Sprintf("logs/amqp-%s.log", time.Now().Format("20060102")))
}
