package module

import (
	"bookstore/internal/app/constants/errors"
	"bookstore/internal/app/database"
	"bookstore/internal/app/domain/orders/items/model"
	"bookstore/internal/app/domain/orders/items/repository/interface"
	"bookstore/internal/app/domain/server"
	"bookstore/internal/app/logger"
	"context"
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

func (m ItemModule) Get() ([]model.Item, *errors.ApiError) {
	return m.Repository.Get()
}

func (m ItemModule) Find(id string) (model.Item, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ItemModule) Create(item model.Item) (uint, *errors.ApiError) {
	client := server.NewOrdersServiceClient(m.GRPCConn)
	req := &server.Request{
		BookId: strconv.FormatUint(uint64(item.BookId), 10),
		Amount: int64(item.Amount),
	}

	resp, err := client.ChangeAmount(context.Background(), req)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToUpdateInventoryAmount,
			Error:   err.Error(),
		}
	}

	if !resp.Success {
		return 0, &errors.ApiError{
			Status:  int(resp.Status),
			Message: errors.FailedToUpdateInventoryAmount,
			Error:   err.Error(),
		}
	}

	return m.Repository.Create(item)
}

func (m ItemModule) Update(id string, upItem model.Item) *errors.ApiError {
	return m.Repository.Update(id, upItem)
}

func (m ItemModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
