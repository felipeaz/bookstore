package model

import (
	"fmt"
	"strings"
	"time"
)

// Book contains all Book's table properties.
type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title     string    `json:"title" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

func (b Book) GetAuthorName() string {
	s := fmt.Sprintf("%s %s", b.FirstName, b.LastName)
	return strings.TrimSpace(s)
}
