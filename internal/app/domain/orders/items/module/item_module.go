package module

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/items/model"
	"bookstore/internal/app/domain/orders/items/model/converter"
	"bookstore/internal/app/domain/orders/items/repository/interface"
	"bookstore/internal/app/domain/server"
	"bookstore/internal/app/logger"
	"context"
	"encoding/json"
	_errors "github.com/pkg/errors"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type ItemModule struct {
	Repository _interface.ItemRepositoryInterface
	Cache      database.CacheInterface
	GRPCConn   *grpc.ClientConn
	Log        logger.LogInterface
}

func NewItemModule(
	repo _interface.ItemRepositoryInterface,
	grpcConn *grpc.ClientConn,
	cache database.CacheInterface,
	log logger.LogInterface) ItemModule {
	return ItemModule{
		Repository: repo,
		Cache:      cache,
		GRPCConn:   grpcConn,
		Log:        log,
	}
}

const (
	AllData = "items"
)

func (m ItemModule) Get() ([]model.Item, *errors.ApiError) {
	b, err := m.Cache.Get(AllData)
	if err != nil {
		m.Log.Error(err)
	}
	if b != nil {
		return converter.ConvertToSliceItemObjFromCache(b)
	}

	item, apiError := m.Repository.Get()
	if apiError != nil {
		return nil, apiError
	}

	m.setAllItemsCache(item)
	return item, nil
}

func (m ItemModule) Find(id string) (model.Item, *errors.ApiError) {
	b, err := m.Cache.Get(id)
	if err != nil {
		m.Log.Error(err)
	}
	if b != nil {
		return converter.ConvertToItemObjFromCache(b)
	}

	item, apiError := m.Repository.Find(id)
	if apiError != nil {
		return model.Item{}, apiError
	}

	m.setItemCache(item)
	return item, nil
}

func (m ItemModule) Create(item model.Item) (model.Item, *errors.ApiError) {
	client := server.NewOrdersServiceClient(m.GRPCConn)
	req := &server.Request{
		BookId: strconv.FormatUint(uint64(item.BookId), 10),
		Amount: int64(item.Amount),
	}

	resp, err := client.ChangeAmount(context.Background(), req)
	if err != nil {
		return model.Item{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToUpdateInventoryAmount,
			Error:   err.Error(),
		}
	}

	if !resp.Success {
		return model.Item{}, &errors.ApiError{
			Status:  int(resp.Status),
			Message: errors.FailedToUpdateInventoryAmount,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Flush(AllData)
	if err != nil {
		m.Log.Error(err)
	}

	item, apiError := m.Repository.Create(item)
	if apiError != nil {
		return model.Item{}, apiError
	}
	m.setItemCache(item)
	return item, nil
}

func (m ItemModule) Update(id string, upItem model.Item) *errors.ApiError {
	apiError := m.Repository.Update(id, upItem)
	if apiError != nil {
		return apiError
	}

	err := m.Cache.Flush(id)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Flush(AllData)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToFlushAllCache)
		m.Log.Error(err)
	}
	return nil
}

func (m ItemModule) Delete(id string) *errors.ApiError {
	apiError := m.Repository.Delete(id)
	if apiError != nil {
		return apiError
	}

	err := m.Cache.Flush(id)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Flush(AllData)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToFlushAllCache)
		m.Log.Error(err)
	}
	return nil
}

func (m ItemModule) setItemCache(item model.Item) {
	b, err := json.Marshal(item)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Set(strconv.FormatUint(uint64(item.ID), 10), b)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToSetCache)
	}
}

func (m ItemModule) setAllItemsCache(items []model.Item) {
	b, err := json.Marshal(items)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Set(AllData, b)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToSetCache)
	}
}
