package crons

import "github.com/robfig/cron/v3"

var Engine *cron.Cron

func Initialize() {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	Engine = cron.New(cron.WithParser(secondParser), cron.WithChain())
}
