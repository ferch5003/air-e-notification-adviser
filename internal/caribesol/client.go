package caribesol

import (
	"air-e-notification-adviser/internal/caribesol/dto"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const _serviceAPIPath = "service/api.php"

type Client interface {
	// GetNIC returns a response for the notifications provided by Air-E.
	GetNIC(ctx context.Context, body dto.ConsultarNICDTORequest) (dto.ConsultarNICDTOResponse, error)
}

type client struct {
	baseUrl    string
	httpClient http.Client
}

func NewClient() Client {
	return client{
		baseUrl: os.Getenv("CARIBE_SOL_BASE_URL"),
		httpClient: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c client) GetNIC(ctx context.Context, body dto.ConsultarNICDTORequest) (dto.ConsultarNICDTOResponse, error) {
	var jsonData []byte
	if data, err := json.Marshal(body); err != nil {

	} else {
		jsonData = data
	}

	caribeSolEndpoint := fmt.Sprintf("%s/%s?rquest=consultar_nic", c.baseUrl, _serviceAPIPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, caribeSolEndpoint, bytes.NewReader(jsonData))
	if err != nil {
		return dto.ConsultarNICDTOResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(req.Body)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return dto.ConsultarNICDTOResponse{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return dto.ConsultarNICDTOResponse{}, errors.New("error on Caribe Sol API")
	}

	apiResponseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var caribeSolResponse dto.ConsultarNICDTOResponse
	err = json.Unmarshal(apiResponseBody, &caribeSolResponse)
	if err != nil {
		return dto.ConsultarNICDTOResponse{}, err
	}

	return caribeSolResponse, nil
}
