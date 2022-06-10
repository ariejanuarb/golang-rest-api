package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
	"go-rest-api/repository"
)

type OrderProductServiceImpl struct {
	OrderProductRepository repository.OrderProductRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewOrderProductService(orderProductRepository repository.OrderProductRepository, DB *sql.DB, validate *validator.Validate) OrderProductService {
	return &OrderProductServiceImpl{
		OrderProductRepository: orderProductRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *OrderProductServiceImpl) Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderProduct := domain.OrderProduct{
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Qty:       request.Qty,
		Price:     request.Price,
		Amount:    request.Amount,
	}

	orderProduct = service.OrderProductRepository.Save(ctx, tx, orderProduct)

	return helper.ToOrderProductResponse(orderProduct)
}

func (service *OrderProductServiceImpl) Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderProduct, err := service.OrderProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	orderProduct.OrderId = request.OrderId
	orderProduct.ProductId = request.ProductId
	orderProduct.Qty = request.Qty
	orderProduct.Price = request.Price
	orderProduct.Amount = request.Amount

	orderProduct = service.OrderProductRepository.Update(ctx, tx, orderProduct)

	return helper.ToOrderProductResponse(orderProduct)
}

func (service *OrderProductServiceImpl) Delete(ctx context.Context, orderProductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderProduct, err := service.OrderProductRepository.FindById(ctx, tx, orderProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderProductRepository.Delete(ctx, tx, orderProduct)
}

func (service *OrderProductServiceImpl) FindById(ctx context.Context, orderProductId int) web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderProduct, err := service.OrderProductRepository.FindById(ctx, tx, orderProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderProductResponse(orderProduct)
}

func (service *OrderProductServiceImpl) FindAll(ctx context.Context) []web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderProducts := service.OrderProductRepository.FindAll(ctx, tx)

	return helper.ToOrderProductResponses(orderProducts)
}
