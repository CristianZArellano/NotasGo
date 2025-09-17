package models

// Note define la estructura mínima de una nota
type Note struct {
	ID      int    `json:"id"`      // identificador único
	Title   string `json:"title"`   // título de la nota
	Content string `json:"content"` // contenido de la nota
}

type Mensaje struct {
	Message string `json:"message"`
}
