package models

import "time"

// Standard API responses
type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operación exitosa"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error en la operación"`
	Error   string `json:"error,omitempty" example:"Detalles del error"`
}

// User responses
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Username  string    `json:"username" example:"johndoe"`
	Email     string    `json:"email" example:"john@example.com"`
	Role      string    `json:"role" example:"user"`
	Status    string    `json:"status" example:"activo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	Success bool         `json:"success" example:"true"`
	Message string       `json:"message" example:"Login exitoso"`
	User    UserResponse `json:"user"`
}

type UsersListResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Usuarios obtenidos exitosamente"`
	Users   []UserResponse `json:"users"`
	Total   int64          `json:"total" example:"10"`
}

// Note responses
type NoteResponse struct {
	ID        int          `json:"id" example:"1"`
	Title     string       `json:"title" example:"Mi nota"`
	Content   string       `json:"content" example:"Contenido de la nota"`
	UserID    uint         `json:"user_id" example:"1"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type NotesListResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Notas obtenidas exitosamente"`
	Notes   []NoteResponse `json:"notes"`
	Total   int64          `json:"total" example:"5"`
}

type UserNotesResponse struct {
	Success  bool           `json:"success" example:"true"`
	Message  string         `json:"message" example:"Notas del usuario obtenidas exitosamente"`
	User     UserResponse   `json:"user"`
	Notes    []NoteResponse `json:"notes"`
	Total    int64          `json:"total" example:"3"`
}