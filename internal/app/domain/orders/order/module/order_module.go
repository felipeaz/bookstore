package module

import (
	"bookstore/build/server/amqp/sender"
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/order/model"
	"bookstore/internal/app/domain/orders/order/model/converter"
	"bookstore/internal/app/domain/orders/order/repository/interface"
	"bookstore/internal/app/domain/server"
	"bookstore/internal/app/logger"
	"context"
	"encoding/json"
	_errors "github.com/pkg/errors"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type OrderModule struct {
	Repository _interface.OrderRepositoryInterface
	Queue      *sender.RabbitMQ
	Cache      database.CacheInterface
	GRPCConn   *grpc.ClientConn
	Log        logger.LogInterface
}

func NewOrderModule(
	repo _interface.OrderRepositoryInterface,
	queue *sender.RabbitMQ,
	grpcConn *grpc.ClientConn,
	cache database.CacheInterface,
	log logger.LogInterface) OrderModule {
	return OrderModule{
		Repository: repo,
		Cache:      cache,
		Queue:      queue,
		GRPCConn:   grpcConn,
		Log:        log,
	}
}

const (
	AllData = "orders"
)

func (m OrderModule) Get() ([]model.Order, *errors.ApiError) {
	b, err := m.Cache.Get(AllData)
	if err != nil {
		m.Log.Error(err)
	}
	if b != nil {
		return converter.ConvertToSliceOrderObjFromCache(b)
	}

	order, apiError := m.Repository.Get()
	if apiError != nil {
		return nil, apiError
	}

	m.setAllOrdersCache(order)
	return order, nil
}

func (m OrderModule) Find(id string) (model.Order, *errors.ApiError) {
	b, err := m.Cache.Get(id)
	if err != nil {
		m.Log.Error(err)
	}
	if b != nil {
		return converter.ConvertToOrderObjFromCache(b)
	}

	order, apiError := m.Repository.Find(id)
	if apiError != nil {
		return model.Order{}, apiError
	}

	m.setOrderCache(order)
	return order, nil
}

func (m OrderModule) Create(orderObj model.Order) (model.Order, *errors.ApiError) {
	apiError := m.updateInventory(orderObj)
	if apiError != nil {
		return model.Order{}, apiError
	}

	order, apiError := m.Repository.Create(orderObj)
	if apiError != nil {
		return model.Order{}, apiError
	}

	m.setOrderCache(order)

	apiError = m.pushOrderToShippingQueue(order)
	if apiError != nil {
		return model.Order{}, apiError
	}

	return order, nil
}

func (m OrderModule) Update(id string, upOrder model.Order) *errors.ApiError {
	apiError := m.Repository.Update(id, upOrder)
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

func (m OrderModule) Delete(id string) *errors.ApiError {
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

func (m OrderModule) setOrderCache(order model.Order) {
	b, err := json.Marshal(order)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Set(strconv.FormatUint(uint64(order.ID), 10), b)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToSetCache)
	}
}

func (m OrderModule) setAllOrdersCache(orders []model.Order) {
	b, err := json.Marshal(orders)
	if err != nil {
		m.Log.Error(err)
	}
	err = m.Cache.Set(AllData, b)
	if err != nil {
		err = _errors.Wrap(err, errors.FailedToSetCache)
	}
}

func (m OrderModule) updateInventory(order model.Order) *errors.ApiError {
	client := server.NewOrdersServiceClient(m.GRPCConn)
	req := &server.Request{
		BookId: strconv.FormatUint(uint64(order.BookId), 10),
		Amount: int64(order.Amount),
	}

	resp, err := client.ChangeAmount(context.Background(), req)
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToUpdateInventoryAmount,
			Error:   err.Error(),
		}
	}

	if !resp.Success {
		return &errors.ApiError{
			Status:  int(resp.Status),
			Message: errors.FailedToUpdateInventoryAmount,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Flush(AllData)
	if err != nil {
		m.Log.Error(err)
	}
	return nil
}

func (m OrderModule) pushOrderToShippingQueue(order model.Order) *errors.ApiError {
	b, err := json.Marshal(order)
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToUnmarshalObj,
			Error:   err.Error(),
		}
	}

	err = m.Queue.PushMessage(b)
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToPushOrderToQueue,
			Error:   err.Error(),
		}
	}

	return nil
}
