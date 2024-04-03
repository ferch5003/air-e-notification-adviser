package caribesol

import (
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
	t.Setenv("CARIBE_SOL_BASE_URL", server.URL)

	caribeSolBody := dto.ConsultarNICDTORequest{}
	caribeSolClient := NewClient()

	// When
	response, err := caribeSolClient.GetNIC(context.Background(), caribeSolBody)

	// Then
	require.NoError(t, err)
	require.Equal(t, dto.NoneEstado, response.Estado)
	require.Equal(t, "no hay notificaciones con este nic", response.Msg)
}

func TestClient_GetNICErrorWrongURL(t *testing.T) {
	server := caribesoltest.NewServer()
	defer server.Close()

	// Given
	t.Setenv("CARIBE_SOL_BASE_URL", "not_exist")

	caribeSolBody := dto.ConsultarNICDTORequest{}
	caribeSolClient := NewClient()

	// When
	_, err := caribeSolClient.GetNIC(context.Background(), caribeSolBody)

	// Then
	require.ErrorContains(t, err, "unsupported protocol scheme")
}
