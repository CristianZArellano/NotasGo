package controllers

import (
	"net/http"
	"notasGo/database"
	"notasGo/models"

	"github.com/gin-gonic/gin"
)

// GetNotes godoc
// @Summary Obtiene todas las notas
// @Description Devuelve una lista de todas las notas almacenadas
// @Tags notas
// @Produce json
// @Success 200 {array} models.Note
// @Failure 500 {object} map[string]string
// @Router /notes [get]
func GetNotes(c *gin.Context) {
	var notes []models.Note
	result := database.DB.Find(&notes) // SELECT * FROM notes;
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, notes)
}

// GetNoteByID godoc
// @Summary Obtiene una nota por ID
// @Description Devuelve una nota específica según su ID
// @Tags notas
// @Produce json
// @Param id path int true "ID de la nota"
// @Success 200 {object} models.Note
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [get]
func GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	result := database.DB.First(&note, id) // SELECT * FROM notes WHERE id = ?
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota no encontrada"})
		return
	}
	c.JSON(http.StatusOK, note)
}

// CreateNote godoc
// @Summary Crea una nueva nota
// @Description Crea una nueva nota con título y contenido
// @Tags notas
// @Accept json
// @Produce json
// @Param note body models.Note true "Nota a crear"
// @Success 200 {object} models.Note
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes [post]
func CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&note) // INSERT INTO notes ...
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

// UpdateNote godoc
// @Summary Actualiza completamente una nota
// @Description Reemplaza todos los campos de una nota según el ID
// @Tags notas
// @Accept json
// @Produce json
// @Param id path int true "ID de la nota"
// @Param note body models.Note true "Nota con los datos actualizados"
// @Success 200 {object} models.Note
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [put]
func UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note

	if err := database.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota no encontrada"})
		return
	}

	var updated models.Note
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Title = updated.Title
	note.Content = updated.Content

	database.DB.Save(&note) // Actualiza en la base de datos
	c.JSON(http.StatusOK, note)
}

// PatchNote godoc
// @Summary Actualiza parcialmente una nota
// @Description Modifica solo los campos enviados de la nota
// @Tags notas
// @Accept json
// @Produce json
// @Param id path int true "ID de la nota"
// @Param note body map[string]interface{} true "Campos a actualizar"
// @Success 200 {object} models.Note
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /notes/{id} [patch]
func PatchNote(c *gin.Context) {
	id := c.Param("id")
	var note models.Note
	if err := database.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nota no encontrada"})
		return
	}

	var partial map[string]interface{}
	if err := c.BindJSON(&partial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&note).Updates(partial) // Actualiza solo los campos presentes
	c.JSON(http.StatusOK, note)
}

// DeleteNote godoc
// @Summary Elimina una nota
// @Description Elimina una nota según su ID
// @Tags notas
// @Produce json
// @Param id path int true "ID de la nota"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes/{id} [delete]
func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Note{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Nota eliminada"})
}

// Home maneja GET /
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, models.Mensaje{Message: "Bienvenido a la API de Notas"})
}
