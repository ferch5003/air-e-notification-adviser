package dto

type Estado string

const NoneEstado Estado = "0"

type ConsultarNICDTOResponse struct {
	Estado Estado `json:"estado"`
	Msg    string `json:"Msg"`
}
