package request

import "time"

type MetaRequest struct {
	Titulo    string    `json:"titulo" validate:"required,min=3,max=100"`
	Descricao string    `json:"descricao" validate:"max=500"`
	Status    string    `json:"status" validate:"required,oneof=pendente em_andamento concluida"`
	Prazo     time.Time `json:"prazo" validate:"required"`
}

// type Meta struct {
// 	gorm.Model
// 	Titulo    string    `json:"titulo" gorm:"not null"`
// 	Descricao string    `json:"descricao"`
// 	Status    string    `json:"status" gorm:"default:'pendente'"`
// 	Prazo     time.Time `json:"prazo" gorm:"not null"`
// 	UserID    uint      `json:"user_id"`
// }
