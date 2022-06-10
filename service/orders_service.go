package service

import (
	"context"
	"go-rest-api/model/web"
)

type OrdersService interface {
	Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse
	Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse
	Delete(ctx context.Context, ordersId int)
	FindById(ctx context.Context, ordersId int) web.OrdersResponse
	FindAll(ctx context.Context) []web.OrdersResponse
}
