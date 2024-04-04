package main

import (
	"air-e-notification-adviser/cmd/background/jobs"
	"air-e-notification-adviser/cmd/background/notifications"
	"air-e-notification-adviser/config"
	"air-e-notification-adviser/internal/caribesol"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"sync"
)

func main() {
	configurations, err := config.NewConfigurations()
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	app := fx.New(
		// creates: config.EnvVars
		fx.Supply(configurations),
		// creates: *zap.Logger
		fx.Supply(logger),
		// creates: *sync.WaitGroup
		fx.Supply(&sync.WaitGroup{}),

		// creates: context.Context
		fx.Provide(context.Background),
		// creates: *cron.Cron
		fx.Provide(cron.New),
		// creates: *caribesol.Client
		fx.Provide(caribesol.NewClient),
		// creates: *notifications.Worker
		fx.Provide(notifications.NewWorker),
		// creates: *jobs.Mail
		fx.Provide(jobs.NewMail),
		// creates: *jobs.Worker
		fx.Provide(jobs.NewWorker),

		// Start job worker.
		fx.Invoke(jobs.Start),
	)

	app.Run()
}
