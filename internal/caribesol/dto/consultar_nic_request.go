package dto

type Tipo string

const (
	NICTipo               Tipo = "1"
	NumeroDeDocumentoTipo Tipo = "2"
)

type ConsultarNICDTORequest struct {
	NIC  string `json:"nic"`
	Tipo Tipo   `json:"tipo"`
}
