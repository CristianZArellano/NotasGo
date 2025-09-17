package controllers

import (
	"net/http"
	"notasGo/models"
	"notasGo/services"
	"notasGo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	noteService *services.NoteService
}

func NewNoteController() *NoteController {
	return &NoteController{
		noteService: services.NewNoteService(),
	}
}

// GetNotes godoc
// @Summary Obtiene todas las notas
// @Description Devuelve una lista de todas las notas con información del usuario
// @Tags notas
// @Produce json
// @Success 200 {object} models.NotesListResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /notes [get]
func (ctrl *NoteController) GetNotes(c *gin.Context) {
	notes, total, err := ctrl.noteService.GetAllNotes()
	if err != nil {
		utils.InternalServerError(c, "Error al obtener notas", err)
		return
	}

	// Convert to response format
	var noteResponses []models.NoteResponse
	for _, note := range notes {
		userResponse := models.UserResponse{
			ID:        note.User.ID,
			Username:  note.User.Username,
			Email:     note.User.Email,
			Role:      note.User.Role,
			Status:    note.User.Status,
			CreatedAt: note.User.CreatedAt,
			UpdatedAt: note.User.UpdatedAt,
		}

		noteResponses = append(noteResponses, models.NoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
			UserID:    note.UserID,
			User:      userResponse,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		})
	}

	response := models.NotesListResponse{
		Success: true,
		Message: "Notas obtenidas exitosamente",
		Notes:   noteResponses,
		Total:   total,
	}

	c.JSON(http.StatusOK, response)
}

// GetNoteByID godoc
// @Summary Obtiene una nota por ID
// @Description Devuelve una nota específica con información del usuario
// @Tags notas
// @Produce json
// @Param id path int true "ID de la nota"
// @Success 200 {object} models.APIResponse{data=models.NoteResponse}
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /notes/{id} [get]
func (ctrl *NoteController) GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	
	note, err := ctrl.noteService.GetNoteByID(id)
	if err != nil {
		if err.Error() == "nota no encontrada" {
			utils.NotFoundError(c, "Nota no encontrada")
			return
		}
		utils.InternalServerError(c, "Error al obtener nota", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        note.User.ID,
		Username:  note.User.Username,
		Email:     note.User.Email,
		Role:      note.User.Role,
		Status:    note.User.Status,
		CreatedAt: note.User.CreatedAt,
		UpdatedAt: note.User.UpdatedAt,
	}

	noteResponse := models.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		UserID:    note.UserID,
		User:      userResponse,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Nota obtenida exitosamente", noteResponse)
}

// CreateNote godoc
// @Summary Crea una nueva nota
// @Description Crea una nueva nota asociada a un usuario
// @Tags notas
// @Accept json
// @Produce json
// @Param note body models.CreateNoteRequest true "Datos de la nota"
// @Success 201 {object} models.APIResponse{data=models.NoteResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /notes [post]
func (ctrl *NoteController) CreateNote(c *gin.Context) {
	var req models.CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(c, "Datos inválidos", err)
		return
	}

	note, err := ctrl.noteService.CreateNote(&req)
	if err != nil {
		if err.Error() == "usuario no encontrado" {
			utils.BadRequestError(c, "Usuario no encontrado", nil)
			return
		}
		utils.InternalServerError(c, "Error al crear nota", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        note.User.ID,
		Username:  note.User.Username,
		Email:     note.User.Email,
		Role:      note.User.Role,
		Status:    note.User.Status,
		CreatedAt: note.User.CreatedAt,
		UpdatedAt: note.User.UpdatedAt,
	}

	noteResponse := models.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		UserID:    note.UserID,
		User:      userResponse,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusCreated, "Nota creada exitosamente", noteResponse)
}

// UpdateNote godoc
// @Summary Actualiza una nota completa
// @Description Reemplaza todos los campos de una nota
// @Tags notas
// @Accept json
// @Produce json
// @Param id path int true "ID de la nota"
// @Param note body models.UpdateNoteRequest true "Datos de la nota"
// @Success 200 {object} models.APIResponse{data=models.NoteResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /notes/{id} [put]
func (ctrl *NoteController) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	
	var req models.UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestError(c, "Datos inválidos", err)
		return
	}

	note, err := ctrl.noteService.UpdateNote(id, &req)
	if err != nil {
		if err.Error() == "nota no encontrada" {
			utils.NotFoundError(c, "Nota no encontrada")
			return
		}
		if err.Error() == "usuario no encontrado" {
			utils.BadRequestError(c, "Usuario no encontrado", nil)
			return
		}
		utils.InternalServerError(c, "Error al actualizar nota", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        note.User.ID,
		Username:  note.User.Username,
		Email:     note.User.Email,
		Role:      note.User.Role,
		Status:    note.User.Status,
		CreatedAt: note.User.CreatedAt,
		UpdatedAt: note.User.UpdatedAt,
	}

	noteResponse := models.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		UserID:    note.UserID,
		User:      userResponse,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Nota actualizada exitosamente", noteResponse)
}

// PatchNote godoc
// @Summary Actualiza parcialmente una nota
// @Description Modifica solo los campos enviados de la nota
// @Tags notas
// @Accept json
// @Produce json
// @Param id path int true "ID de la nota"
// @Param updates body map[string]interface{} true "Campos a actualizar"
// @Success 200 {object} models.APIResponse{data=models.NoteResponse}
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /notes/{id} [patch]
func (ctrl *NoteController) PatchNote(c *gin.Context) {
	id := c.Param("id")
	
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		utils.BadRequestError(c, "Datos inválidos", err)
		return
	}

	note, err := ctrl.noteService.PatchNote(id, updates)
	if err != nil {
		if err.Error() == "nota no encontrada" {
			utils.NotFoundError(c, "Nota no encontrada")
			return
		}
		if err.Error() == "usuario no encontrado" {
			utils.BadRequestError(c, "Usuario no encontrado", nil)
			return
		}
		utils.InternalServerError(c, "Error al actualizar nota", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        note.User.ID,
		Username:  note.User.Username,
		Email:     note.User.Email,
		Role:      note.User.Role,
		Status:    note.User.Status,
		CreatedAt: note.User.CreatedAt,
		UpdatedAt: note.User.UpdatedAt,
	}

	noteResponse := models.NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		UserID:    note.UserID,
		User:      userResponse,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	utils.SuccessResponse(c, http.StatusOK, "Nota actualizada exitosamente", noteResponse)
}

// DeleteNote godoc
// @Summary Elimina una nota
// @Description Elimina una nota por su ID
// @Tags notas
// @Produce json
// @Param id path int true "ID de la nota"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /notes/{id} [delete]
func (ctrl *NoteController) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	
	err := ctrl.noteService.DeleteNote(id)
	if err != nil {
		if err.Error() == "nota no encontrada" {
			utils.NotFoundError(c, "Nota no encontrada")
			return
		}
		utils.InternalServerError(c, "Error al eliminar nota", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Nota eliminada exitosamente", nil)
}

// GetNotesByUser godoc
// @Summary Obtiene todas las notas de un usuario
// @Description Devuelve todas las notas que pertenecen a un usuario específico
// @Tags notas
// @Produce json
// @Param user_id path int true "ID del usuario"
// @Success 200 {object} models.UserNotesResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/user/{user_id}/notes [get]
func (ctrl *NoteController) GetNotesByUser(c *gin.Context) {
	userID := c.Param("user_id")
	
	user, notes, total, err := ctrl.noteService.GetNotesByUser(userID)
	if err != nil {
		if err.Error() == "usuario no encontrado" {
			utils.NotFoundError(c, "Usuario no encontrado")
			return
		}
		utils.InternalServerError(c, "Error al obtener notas del usuario", err)
		return
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	var noteResponses []models.NoteResponse
	for _, note := range notes {
		noteResponses = append(noteResponses, models.NoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
			UserID:    note.UserID,
			User:      userResponse,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		})
	}

	response := models.UserNotesResponse{
		Success: true,
		Message: "Notas del usuario obtenidas exitosamente",
		User:    userResponse,
		Notes:   noteResponses,
		Total:   total,
	}

	c.JSON(http.StatusOK, response)
}

// Legacy form handlers (keep for backward compatibility with HTML forms)
func (ctrl *NoteController) CreateNoteForm(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	userIDStr := c.PostForm("user_id")

	var userID uint = 1 // Default user
	if userIDStr != "" {
		if id, err := strconv.ParseUint(userIDStr, 10, 32); err == nil {
			userID = uint(id)
		}
	}

	req := models.CreateNoteRequest{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	_, err := ctrl.noteService.CreateNote(&req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{"Error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func (ctrl *NoteController) UpdateNoteForm(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	content := c.PostForm("content")

	req := models.UpdateNoteRequest{
		Title:   title,
		Content: content,
	}

	_, err := ctrl.noteService.UpdateNote(id, &req)
	if err != nil {
		c.HTML(http.StatusNotFound, "index.html", gin.H{"Error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func (ctrl *NoteController) DeleteNoteForm(c *gin.Context) {
	id := c.PostForm("id")
	
	err := ctrl.noteService.DeleteNote(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "index.html", gin.H{"Error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}