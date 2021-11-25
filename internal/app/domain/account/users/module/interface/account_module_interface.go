package _interface

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/constants/login"
	"bookstore/internal/app/domain/account/users/model"
)

type AccountModuleInterface interface {
	Login(credentials model.Account) login.Message
	Logout(session model.UserSession) (message login.Message)
	Get() ([]model.Account, *errors.ApiError)
	Find(id string) (model.Account, *errors.ApiError)
	Create(account model.Account) (uint, *errors.ApiError)
	Update(id string, upAccount model.Account) *errors.ApiError
	Delete(id string) *errors.ApiError
}
