package model

import (
	"fmt"
	"time"
)

// Book contains all Book's table properties.
type Book struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title           string    `json:"title"`
	AuthorFirstName string    `json:"authorFirstName"`
	AuthorLastName  string    `json:"authorLastName"`
	Amount          int64     `json:"amount"`
	CreatedAt       time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt       time.Time `time_format:"2006-01-02 15:04:05"`
}

func (b Book) GetAuthorName() string {
	return fmt.Sprintf("%s %s", b.AuthorFirstName, b.AuthorLastName)
}
