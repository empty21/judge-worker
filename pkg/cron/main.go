package cron

import (
	"github.com/go-co-op/gocron"
	"judger/pkg/file"
	"time"
)

func init() {
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(1).Hour().Do(file.CleanUpCache)
	s.StartAsync()
}
