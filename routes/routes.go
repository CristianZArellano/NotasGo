package routes

import (
	"notasGo/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Inicializa Gin con logger y recovery

	// Ruta de inicio
	r.GET("/", controllers.Home)

	// Rutas de notas
	r.GET("/notes", controllers.GetNotes)          // Obtener todas las notas
	r.GET("/notes/:id", controllers.GetNoteByID)   // Obtener una nota por ID
	r.POST("/notes", controllers.CreateNote)       // Crear una nueva nota
	r.PUT("/notes/:id", controllers.UpdateNote)    // Actualizar una nota completa (PUT)
	r.PATCH("/notes/:id", controllers.PatchNote)   // Actualizar parcialmente (PATCH)
	r.DELETE("/notes/:id", controllers.DeleteNote) // Eliminar una nota

	// Ruta de Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
