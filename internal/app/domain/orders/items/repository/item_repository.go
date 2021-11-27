package repository

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/items/model"
	"bookstore/internal/app/domain/orders/items/model/converter"
)

type ItemRepository struct {
	DB database.GORMServiceInterface
}

func NewItemRepository(db database.GORMServiceInterface) ItemRepository {
	return ItemRepository{
		DB: db,
	}
}

func (r ItemRepository) Get() ([]model.Item, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]model.Item{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	items, apiError := converter.ConvertToSliceItemObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return items, nil
}

func (r ItemRepository) Find(id string) (model.Item, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Item{}, id)
	if err != nil {
		return model.Item{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	item, apiError := converter.ConvertToItemObj(result)
	if apiError != nil {
		return model.Item{}, apiError
	}
	return item, nil
}

func (r ItemRepository) Create(item model.Item) (uint, *errors.ApiError) {
	err := r.DB.Persist(&item)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return item.ID, nil
}

func (r ItemRepository) Update(id string, upItem model.Item) *errors.ApiError {
	err := r.DB.Refresh(&upItem, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r ItemRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Item{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
