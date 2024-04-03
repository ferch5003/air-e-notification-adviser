package caribesol

import (
	"air-e-notification-adviser/config"
	"air-e-notification-adviser/internal/caribesol/caribesoltest"
	"air-e-notification-adviser/internal/caribesol/dto"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetNICSuccessfulRequest(t *testing.T) {
	server := caribesoltest.NewServer()
	defer server.Close()

	// Given
	cfg := &config.EnvVars{
		CaribeSolBaseURL: server.URL,
	}

	caribeSolBody := dto.ConsultarNICDTORequest{}
	caribeSolClient := NewClient(context.Background(), cfg)

	// When
	response, err := caribeSolClient.GetNIC(caribeSolBody)

	// Then
	require.NoError(t, err)
	require.Equal(t, dto.NoneEstado, response.Estado)
	require.Equal(t, "no hay notificaciones con este nic", response.Msg)
}

func TestClient_GetNICErrorWrongURL(t *testing.T) {
	server := caribesoltest.NewServer()
	defer server.Close()

	// Given
	cfg := &config.EnvVars{
		CaribeSolBaseURL: "not_exist",
	}

	caribeSolBody := dto.ConsultarNICDTORequest{}
	caribeSolClient := NewClient(context.Background(), cfg)

	// When
	_, err := caribeSolClient.GetNIC(caribeSolBody)

	// Then
	require.ErrorContains(t, err, "unsupported protocol scheme")
}
