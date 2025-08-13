package models

import (
	"time"

	"gorm.io/gorm"
)

type Meta struct {
	gorm.Model
	Titulo    string    `json:"titulo" gorm:"not null"`
	Descricao string    `json:"descricao"`
	Status    string    `json:"status" gorm:"default:'pendente'"`
	Prazo     time.Time `json:"prazo" gorm:"not null"`
	UserID    uint      `json:"user_id"`
}
