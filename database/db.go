package database

import (
	"log"
	"projeto-metas/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conectar() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Meta{})
}
