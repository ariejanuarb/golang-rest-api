package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
	"go-rest-api/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	logrus.Info("product Create service_impl start")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	logrus.Info("product Create service_impl end")
	return helper.ToProductResponse(product)

}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	logrus.Info("product Update service_impl start")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	product := helper.ToProduct(productResponse)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Price = request.Price
	product.CategoryId = request.CategoryId

	product = service.ProductRepository.Update(ctx, tx, product)

	logrus.Info("product Update service_impl start")
	return helper.ToProductResponse(product)

}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	logrus.Info("product Delete service_impl start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product := helper.ToProduct(productResponse)
	service.ProductRepository.Delete(ctx, tx, product)
	logrus.Info("product Delete service_impl end")
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	logrus.Info("product FindById service_impl start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	logrus.Info("product FindById service_impl end")
	return product
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	logrus.Info("product FindAll service_impl start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	//return helper.ToProductResponses(products) // return for FindAll without INNER Join
	logrus.Info("product FindAll service_impl end")
	return products
}
