package service

import (
	"context"
	"go-rest-api/model/web"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, customerId int)
	FindById(ctx context.Context, customerId int) web.CustomerResponse
	FindAll(ctx context.Context) []web.CustomerResponse
}
