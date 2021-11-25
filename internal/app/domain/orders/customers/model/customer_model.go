package model

import "fmt"

type Customer struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (c Customer) GetName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}
