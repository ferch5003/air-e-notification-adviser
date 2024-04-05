package config

import (
	"air-e-notification-adviser/internal/platform/files"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type EnvVars struct {
	// App Data.
	AppName          string
	Port             string
	CaribeSolBaseURL string
	SearchNICCron    string

	// Data.
	NIC  string
	Tipo string

	// SMTP.
	SMTPHost        string
	SMTPPort        int
	SMTPUsername    string
	SMTPPassword    string
	SMTPFromAddress string

	// Mail.
	FromMail string
	ToMail   string
}

func NewConfigurations() (*EnvVars, error) {
	area := os.Getenv("AREA")

	if area == "" {
		envFilepath, err := files.GetFile(".env")
		if err != nil {
			return nil, err
		}

		if err := godotenv.Load(envFilepath); err != nil {
			return nil, err
		}
	}

	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	caribeSolBaseURL := os.Getenv("CARIBE_SOL_BASE_URL")
	searchNICCron := os.Getenv("SEARCH_NIC_CRON")

	nic := os.Getenv("NIC")
	tipo := os.Getenv("TIPO")

	smtpHost := os.Getenv("SMTP_HOST")
	stmStrPort := os.Getenv("SMTP_PORT")

	smtpPort, err := strconv.Atoi(stmStrPort)
	if err != nil {
		return nil, err
	}

	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFromAddress := os.Getenv("SMTP_FROM_ADDRESS")

	fromMail := os.Getenv("FROM_MAIL")
	toMail := os.Getenv("TO_MAIL")

	environment := &EnvVars{
		AppName:          appName,
		Port:             port,
		CaribeSolBaseURL: caribeSolBaseURL,
		SearchNICCron:    searchNICCron,

		NIC:  nic,
		Tipo: tipo,

		SMTPHost:        smtpHost,
		SMTPPort:        smtpPort,
		SMTPUsername:    smtpUsername,
		SMTPPassword:    smtpPassword,
		SMTPFromAddress: smtpFromAddress,

		FromMail: fromMail,
		ToMail:   toMail,
	}

	return environment, nil
}
