package services

import (
	"errors"
	"fmt"
	"notasGo/database"
	"notasGo/models"

	"gorm.io/gorm"
)

type NoteService struct {
	userService *UserService
}

func NewNoteService() *NoteService {
	return &NoteService{
		userService: NewUserService(),
	}
}

// GetAllNotes retrieves all notes with user information
func (s *NoteService) GetAllNotes() ([]models.Note, int64, error) {
	var notes []models.Note
	var count int64
	
	if err := database.DB.Preload("User").Find(&notes).Error; err != nil {
		return nil, 0, err
	}
	
	database.DB.Model(&models.Note{}).Count(&count)
	return notes, count, nil
}

// GetNoteByID retrieves a note by ID with user information
func (s *NoteService) GetNoteByID(id string) (*models.Note, error) {
	var note models.Note
	if err := database.DB.Preload("User").First(&note, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("nota no encontrada")
		}
		return nil, err
	}
	return &note, nil
}

// CreateNote creates a new note
func (s *NoteService) CreateNote(req *models.CreateNoteRequest) (*models.Note, error) {
	// Verify user exists
	_, err := s.userService.GetUserByID(fmt.Sprintf("%d", req.UserID))
	if err != nil {
		return nil, errors.New("usuario no encontrado")
	}

	note := models.Note{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := database.DB.Create(&note).Error; err != nil {
		return nil, err
	}

	// Load user information for response
	createdNote, err := s.GetNoteByID(fmt.Sprintf("%d", note.ID))
	if err != nil {
		return nil, err
	}

	return createdNote, nil
}

// UpdateNote updates an existing note
func (s *NoteService) UpdateNote(id string, req *models.UpdateNoteRequest) (*models.Note, error) {
	note, err := s.GetNoteByID(id)
	if err != nil {
		return nil, err
	}

	// If UserID is being changed, verify the new user exists
	if req.UserID != 0 && req.UserID != note.UserID {
		_, err := s.userService.GetUserByID(fmt.Sprintf("%d", req.UserID))
		if err != nil {
			return nil, errors.New("usuario no encontrado")
		}
	}

	// Update fields
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.UserID != 0 {
		updates["user_id"] = req.UserID
	}

	if err := database.DB.Model(note).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Return updated note with user information
	updatedNote, err := s.GetNoteByID(id)
	if err != nil {
		return nil, err
	}

	return updatedNote, nil
}

// PatchNote partially updates a note
func (s *NoteService) PatchNote(id string, updates map[string]interface{}) (*models.Note, error) {
	note, err := s.GetNoteByID(id)
	if err != nil {
		return nil, err
	}

	// Validate UserID if present
	if userID, exists := updates["user_id"]; exists {
		if uid, ok := userID.(float64); ok {
			_, err := s.userService.GetUserByID(fmt.Sprintf("%d", int(uid)))
			if err != nil {
				return nil, errors.New("usuario no encontrado")
			}
		}
	}

	if err := database.DB.Model(note).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Return updated note with user information
	updatedNote, err := s.GetNoteByID(id)
	if err != nil {
		return nil, err
	}

	return updatedNote, nil
}

// DeleteNote deletes a note by ID
func (s *NoteService) DeleteNote(id string) error {
	note, err := s.GetNoteByID(id)
	if err != nil {
		return err
	}

	if err := database.DB.Delete(note).Error; err != nil {
		return errors.New("error al eliminar nota")
	}

	return nil
}

// GetNotesByUser retrieves all notes for a specific user
func (s *NoteService) GetNotesByUser(userID string) (*models.User, []models.Note, int64, error) {
	// Verify user exists
	user, err := s.userService.GetUserByID(userID)
	if err != nil {
		return nil, nil, 0, err
	}

	var notes []models.Note
	var count int64
	
	if err := database.DB.Preload("User").Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return nil, nil, 0, err
	}
	
	database.DB.Model(&models.Note{}).Where("user_id = ?", userID).Count(&count)
	
	return user, notes, count, nil
}