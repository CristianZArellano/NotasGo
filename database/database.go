package database

import (
	"notasGo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("notas.db"), &gorm.Config{})
	if err != nil {
		panic("No se pudo conectar a la base de datos: " + err.Error())
	}

	db.AutoMigrate(&models.Note{}, &models.User{})

	DB = db
}
