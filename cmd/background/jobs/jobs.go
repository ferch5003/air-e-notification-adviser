package jobs

import (
	"air-e-notification-adviser/cmd/background/notifications"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Worker struct {
	ErrorChan chan error
	DoneChan  chan bool
}

func NewWorker() *Worker {
	jobWorker := &Worker{
		ErrorChan: make(chan error),
		DoneChan:  make(chan bool),
	}

	return jobWorker
}

func Start(
	lc fx.Lifecycle,
	c *cron.Cron,
	mainWorker *Worker,
	notificationsWorker *notifications.Worker,
	logger *zap.Logger) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Starting job worker"))

			go func() {
				logger.Info("Starting...")

				id, err := notificationsWorker.Start()
				if err != nil {
					logger.Info("err: ", zap.Error(err))
				}

				logger.Info(fmt.Sprintf("Notification worker ID: %d", id))

				c.Start()

				select {
				case <-mainWorker.DoneChan:
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing job worker...")

			return nil
		},
	})
}
