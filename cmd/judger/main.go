package main

import (
	"judger/pkg/amqp"
	_ "judger/pkg/config"
	_ "judger/pkg/logger"
	_ "judger/pkg/redis"
	_ "judger/pkg/runner"
)

func main() {
	amqp.Init()
}
