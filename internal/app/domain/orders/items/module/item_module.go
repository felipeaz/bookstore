package module

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/items/model"
	"bookstore/internal/app/domain/orders/items/repository/interface"
	"bookstore/internal/app/logger"
)

type ItemModule struct {
	Repository _interface.ItemRepositoryInterface
	Cache      database.CacheInterface
	Log        logger.LogInterface
}

func NewItemModule(repo _interface.ItemRepositoryInterface, cache database.CacheInterface, log logger.LogInterface) ItemModule {
	return ItemModule{
		Repository: repo,
		Cache:      cache,
		Log:        log,
	}
}

func (m ItemModule) Get(bookId string) ([]model.Item, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ItemModule) Find(id string) (model.Item, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ItemModule) Create(item model.Item) (uint, *errors.ApiError) {
	return m.Repository.Create(item)
}

func (m ItemModule) Update(id string, upItem model.Item) *errors.ApiError {
	return m.Repository.Update(id, upItem)
}

func (m ItemModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
