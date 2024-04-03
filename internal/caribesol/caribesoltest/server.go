package caribesoltest

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

type Server struct {
	URL string
}

func NewServer() *httptest.Server {
	return httptest.NewServer(&Server{})
}

func (s Server) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	if strings.Contains(rq.URL.String(), "/service/api.php?rquest=consultar_nic") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		caribeSolResponse := getCaribeSolSuccessfulResponse()

		_, err := w.Write([]byte(caribeSolResponse))
		if err != nil {
			return
		}
	}
}

func getCaribeSolSuccessfulResponse() string {
	rp := `{
				"estado": "0",
				"msg": "no hay notificaciones con este nic"
			}`
	return rp
}
