package models

import "time"

// Note define la estructura mínima de una nota
type Note struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id" gorm:"default:1"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Mensaje struct {
	Message string `json:"message"`
}
