package model

import (
	"time"
)

type Item struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint      `json:"bookId" binding:"required"`
	UserId    uint      `json:"userId" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
