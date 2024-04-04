package dto

type Estado string

const NoneEstado Estado = "0"

type Historial struct {
	ID               int    `json:"id"`
	NuevaFecha       string `json:"nuevafecha"`
	NIC              string `json:"nic"`
	NoDocumento      string `json:"no_documento"`
	NombreDocumento  string `json:"nombre_documento"`
	Municipio        string `json:"municipio"`
	Barrio           string `json:"barrio"`
	Via              string `json:"via"`
	Crucero          string `json:"crucero"`
	Placa            string `json:"placa"`
	NombreReclamante string `json:"nombre_reclamante"`
	Estado           string `json:"estado"`
}

type ConsultarNICDTOResponse struct {
	Estado    Estado      `json:"estado"`
	Msg       string      `json:"msg"`
	Historial []Historial `json:"historial"`
}
