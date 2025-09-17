package main

import (
	"notasGo/database"
	"notasGo/routes"

	_ "notasGo/docs" // documentación generada por swag
)

// @title NotasGo API
// @version 2.0
// @description API REST refactorizada para gestión de notas y usuarios con Go, Gin y GORM
// @host localhost:8080
// @BasePath /
// @schemes http
// @contact.name Soporte API NotasGo
// @contact.email soporte@notasgo.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

func main() {
	// Inicializar la base de datos
	database.Connect()

	// Inicializar Gin con rutas
	r := routes.SetupRouter()

	// Servir archivos estáticos (CSS, imágenes, JS, etc.)
	r.Static("/static", "./static")

	// Cargar templates HTML
	r.LoadHTMLGlob("templates/*")

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
