package models

// User request/response structures
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

type UpdateUserRequest struct {
	Username string `json:"username,omitempty" binding:"omitempty,min=3,max=50" example:"johndoe_updated"`
	Email    string `json:"email,omitempty" binding:"omitempty,email" example:"john.doe@example.com"`
	Role     string `json:"role,omitempty" binding:"omitempty,oneof=user admin" example:"admin"`
	Status   string `json:"status,omitempty" binding:"omitempty,oneof=activo inactivo" example:"activo"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// Note request structures
type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200" example:"Mi nota importante"`
	Content string `json:"content" binding:"required" example:"Esta es el contenido de mi nota"`
	UserID  uint   `json:"user_id" binding:"required" example:"1"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title,omitempty" binding:"omitempty,min=1,max=200" example:"Nota actualizada"`
	Content string `json:"content,omitempty" example:"Contenido actualizado"`
	UserID  uint   `json:"user_id,omitempty" example:"1"`
}