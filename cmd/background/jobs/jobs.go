package jobs

import (
	"air-e-notification-adviser/cmd/background/notifications"
	"air-e-notification-adviser/config"
	"air-e-notification-adviser/internal/caribesol/dto"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	Mailer        *Mail
	Wait          *sync.WaitGroup
	ErrorChan     chan error
	ErrorChanDone chan bool
}

func NewWorker(mailer *Mail, wg *sync.WaitGroup) *Worker {
	return &Worker{
		Mailer:        mailer,
		Wait:          wg,
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}
}

// listenForShutdown listens for SIGINT and SIGTERM signals and calls the shutdown
// function when they are received. This prevents goroutines to stop immediately and
// let them finish their jobs first, like sending and email, etc.
// IMPORTANT NOTE: One of main part of concurrency.
func (w *Worker) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	w.shutdown()
}

// listenForNotifications listens for incoming notifications from CaribeSol API.
func (w *Worker) listenForNotifications(
	cfg *config.EnvVars,
	notificationsWorker *notifications.Worker,
	logger *zap.Logger) {
	for {
		select {
		case response := <-notificationsWorker.ResponseChan:
			t := time.Now()

			// print location and local time
			location, err := time.LoadLocation("America/Bogota")
			if err != nil {
				fmt.Println(err)
			}

			subject := "Notificación de Air-E aun sin responder"
			if response.Estado != "0" {
				subject = "Notificación de Air-E con posible respuesta"
			}

			message := Message{
				From:    cfg.FromMail,
				To:      cfg.ToMail,
				Subject: subject,
				Data: struct {
					ConsultarNICDTOResponse dto.ConsultarNICDTOResponse
					Datetime                string
				}{
					ConsultarNICDTOResponse: response,
					Datetime:                fmt.Sprintf("%v", t.In(location)),
				},
			}

			w.sendEmail(message)
		case err := <-w.ErrorChan:
			logger.Error(err.Error())
		case <-w.ErrorChanDone:
			return
		}
	}
}

func (w *Worker) shutdown() {
	// perform any cleanup tasks.
	fmt.Println("would run cleanup tasks...")

	// block until WaitGroup is empty.
	w.Wait.Wait()

	w.Mailer.DoneChan <- true
	w.ErrorChanDone <- true

	fmt.Println("closing channels and shutting down application...")

	close(w.Mailer.MailerChan)
	close(w.Mailer.ErrorChan)
	close(w.Mailer.DoneChan)
	close(w.ErrorChan)
	close(w.ErrorChanDone)
}

func Start(
	lc fx.Lifecycle,
	cfg *config.EnvVars,
	c *cron.Cron,
	mainWorker *Worker,
	notificationsWorker *notifications.Worker,
	logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("Starting job worker"))

			go mainWorker.listenForMail(logger)

			go mainWorker.listenForNotifications(cfg, notificationsWorker, logger)

			go func() {
				logger.Info("Starting...")

				id, err := notificationsWorker.Start()
				if err != nil {
					logger.Info("err: ", zap.Error(err))
				}

				logger.Info(fmt.Sprintf("Notification worker ID: %d", id))

				c.Start()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing job worker...")

			go mainWorker.listenForShutdown()

			return nil
		},
	})
}
