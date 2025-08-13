package response

type MetaResponse struct {
	ID        uint   `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
	Prazo     string `json:"prazo"` // pode ser string formatada
}
