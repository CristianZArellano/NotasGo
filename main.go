package main

import (
	"notasGo/database"
	"notasGo/routes"

	_ "notasGo/docs" // documentación generada por swag
)

func main() {
	database.Connect() // Inicializa SQLite y crea tablas si no existen
	r := routes.SetupRouter()
	r.Run(":8080")
}
