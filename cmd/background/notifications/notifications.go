package notifications

import (
	"air-e-notification-adviser/config"
	"air-e-notification-adviser/internal/caribesol"
	"air-e-notification-adviser/internal/caribesol/dto"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// defaultSearchNICCronSpec is the time to run the job to search NIC at 12:00.
const defaultSearchNICCronSpec = "0 12 * * *"

type Worker struct {
	ctx             context.Context
	cfg             *config.EnvVars
	cron            *cron.Cron
	caribeSolClient caribesol.Client
	ResponseChan    chan dto.ConsultarNICDTOResponse
	ErrorChan       chan error
	DoneChan        chan bool
	logger          *zap.Logger
}

func NewWorker(
	ctx context.Context,
	cfg *config.EnvVars,
	cron *cron.Cron,
	caribeSolClient caribesol.Client,
	logger *zap.Logger) *Worker {
	return &Worker{
		ctx:             ctx,
		cfg:             cfg,
		cron:            cron,
		caribeSolClient: caribeSolClient,
		ResponseChan:    make(chan dto.ConsultarNICDTOResponse, 100),
		ErrorChan:       make(chan error),
		DoneChan:        make(chan bool),
		logger:          logger,
	}
}

func (w Worker) Start() (int, error) {
	cronSpec := defaultSearchNICCronSpec
	if w.cfg.SearchNICCron != "" {
		cronSpec = w.cfg.SearchNICCron
	}

	entryID, err := w.cron.AddFunc(cronSpec, w.searchCaribeSolAPIJob)
	if err != nil {
		return 0, err
	}

	return int(entryID), nil
}

func (w Worker) searchCaribeSolAPIJob() {
	caribesolReq := dto.ConsultarNICDTORequest{
		NIC:  w.cfg.NIC,
		Tipo: dto.Tipo(w.cfg.Tipo),
	}

	response, err := w.caribeSolClient.GetNIC(caribesolReq)
	if err != nil {
		w.ErrorChan <- err
	}

	w.ResponseChan <- response
}
