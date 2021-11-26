package model

import (
	"fmt"
	"time"
)

type Item struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId          uint      `json:"bookId" binding:"required"`
	ClientFirstName string    `json:"clientFirstName" binding:"required"`
	ClientLastName  string    `json:"clientLastName" binding:"required"`
	Amount          int       `json:"amount" binding:"required"`
	CreatedAt       time.Time `json:"orderDate" time_format:"2006-01-02 15:04:05"`
	UpdatedAt       time.Time `time_format:"2006-01-02 15:04:05"`
}

func (i Item) GetConsumerName() string {
	return fmt.Sprintf("%s %s", i.ClientFirstName, i.ClientLastName)
}
