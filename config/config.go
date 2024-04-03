package config

import (
	"air-e-notification-adviser/internal/platform/files"
	"github.com/joho/godotenv"
	"os"
)

type EnvVars struct {
	// App Data.
	AppName          string
	CaribeSolBaseURL string

	// Data
	NIC string
}

func NewConfigurations() (*EnvVars, error) {
	envFilepath, err := files.GetFile(".env")
	if err != nil {
		return nil, err
	}

	if err := godotenv.Load(envFilepath); err != nil {
		return nil, err
	}

	appName := os.Getenv("APP_NAME")
	caribeSolBaseURL := os.Getenv("CARIBE_SOL_BASE_URL")
	nic := os.Getenv("NIC")

	environment := &EnvVars{
		AppName:          appName,
		CaribeSolBaseURL: caribeSolBaseURL,
		NIC:              nic,
	}

	return environment, nil
}
