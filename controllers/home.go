package controllers

import (
	"net/http"
	"notasGo/database"
	"notasGo/models"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	var notes []models.Note
	if err := database.DB.Find(&notes).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"Title": "NotasGo",
			"Error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": "NotasGo",
		"Notes": notes,
	})
}
