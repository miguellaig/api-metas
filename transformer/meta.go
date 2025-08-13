package transformer

import (
	"projeto-metas/dto/request"
	"projeto-metas/dto/response"
	"projeto-metas/models"
)

func ToMetaModel(meta *request.MetaRequest) *models.Meta {
	return &models.Meta{
		Titulo:    meta.Titulo,
		Descricao: meta.Descricao,
		Status:    meta.Status,
		Prazo:     meta.Prazo,
	}
}

func ToMetaResponse(meta *models.Meta) *response.MetaResponse {
	return &response.MetaResponse{
		ID:        meta.ID,
		Titulo:    meta.Titulo,
		Descricao: meta.Descricao,
		Status:    meta.Status,
		Prazo:     meta.Prazo.Format("2006-01-02"), // formatando a data só como yyyy-mm-dd
	}
}

// type Meta struct {
// 	gorm.Model
// 	Titulo    string    `json:"titulo" gorm:"not null"`
// 	Descricao string    `json:"descricao"`
// 	Status    string    `json:"status" gorm:"default:'pendente'"`
// 	Prazo     time.Time `json:"prazo" gorm:"not null"`
// 	UserID    uint      `json:"user_id"`
// }
