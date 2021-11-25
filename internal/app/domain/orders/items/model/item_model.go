package model

import (
	"bookstore/internal/app/domain/orders/customers/model"
	"time"
)

type Item struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint           `json:"bookId" binding:"required"`
	Customer  model.Customer `json:"customers" binding:"required"`
	Amount    int            `json:"amount" binding:"required"`
	CreatedAt time.Time      `json:"orderDate" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time      `time_format:"2006-01-02 15:04:05"`
}
