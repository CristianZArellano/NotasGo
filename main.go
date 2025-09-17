package main

import (
	"notasGo/database"
	"notasGo/routes"

	_ "notasGo/docs" // documentación generada por swag
)

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
