package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/repository"
	"go-rest-api/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	orderProductRepository := repository.NewOrderProductRepository()
	orderProductService := service.NewOrderProductService(orderProductRepository, db, validate)
	orderProductController := controller.NewOrderProductController(orderProductService)

	ordersRepository := repository.NewOrdersRepository()
	ordersService := service.NewOrdersService(ordersRepository, db, validate)
	ordersController := controller.NewOrdersController(ordersService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(categoryController, customerController, orderProductController, ordersController, productController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
