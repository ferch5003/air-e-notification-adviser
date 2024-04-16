package dto

type Estado string

const NoneEstado Estado = "0"

type Historial struct {
	ID               string `json:"id"`
	NuevaFecha       string `json:"nuevafecha"`
	FechaSystem      string `json:"fecha_system"`
	Fecha            string `json:"fecha"`
	FechaFestivo     string `json:"fecha_festivo"`
	NIC              string `json:"nic"`
	NoDocumento      string `json:"no_documento"`
	NombreDocumento  string `json:"nombre_documento"`
	Municipio        string `json:"municipio"`
	Barrio           string `json:"barrio"`
	Via              string `json:"via"`
	Crucero          string `json:"crucero"`
	Placa            string `json:"placa"`
	NombreReclamante string `json:"nombre_reclamante"`
	StartTS          string `json:"start_ts"`
	TestTS           string `json:"test_ts"`
	Estado           string `json:"estado"`
}

type ConsultarNICDTOResponse struct {
	Estado    Estado      `json:"estado"`
	Msg       string      `json:"msg"`
	Historial []Historial `json:"historial"`
}
