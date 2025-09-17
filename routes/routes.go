package routes

import (
	"notasGo/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Initialize controllers
	userController := controllers.NewUserController()
	noteController := controllers.NewNoteController()

	// API v1 routes group
	v1 := r.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.GET("", userController.GetUsers)
			users.GET("/:id", userController.GetUserByID)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

		// Authentication routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userController.RegisterUser)
			auth.POST("/login", userController.LoginUser)
		}

		// Note routes
		notes := v1.Group("/notes")
		{
			notes.GET("", noteController.GetNotes)
			notes.GET("/:id", noteController.GetNoteByID)
			notes.POST("", noteController.CreateNote)
			notes.PUT("/:id", noteController.UpdateNote)
			notes.PATCH("/:id", noteController.PatchNote)
			notes.DELETE("/:id", noteController.DeleteNote)
		}

		// User notes routes (moved outside users group to avoid conflicts)
		v1.GET("/user/:user_id/notes", noteController.GetNotesByUser)
	}

	// Legacy routes for backward compatibility
	legacy := r.Group("")
	{
		// Dashboard route
		legacy.GET("/", controllers.Dashboard)

		// Legacy user routes
		legacy.GET("/users", userController.GetUsers)
		legacy.GET("/users/:id", userController.GetUserByID)
		legacy.PUT("/users/:id", userController.UpdateUser)
		legacy.DELETE("/users/:id", userController.DeleteUser)
		legacy.POST("/register", userController.RegisterUser)
		legacy.POST("/login", userController.LoginUser)

		// Legacy note routes
		legacy.GET("/notes", noteController.GetNotes)
		legacy.GET("/notes/:id", noteController.GetNoteByID)
		legacy.POST("/notes", noteController.CreateNote)
		legacy.PUT("/notes/:id", noteController.UpdateNote)
		legacy.PATCH("/notes/:id", noteController.PatchNote)
		legacy.DELETE("/notes/:id", noteController.DeleteNote)

		// User notes route
		legacy.GET("/user/:user_id/notes", noteController.GetNotesByUser)

		// HTML form routes
		legacy.POST("/notes/create", noteController.CreateNoteForm)
		legacy.POST("/notes/delete", noteController.DeleteNoteForm)
		legacy.POST("/notes/update", noteController.UpdateNoteForm)
	}

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}