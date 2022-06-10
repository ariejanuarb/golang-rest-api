package app

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/controller"
	"go-rest-api/exception"
)

func NewRouter(categoryController controller.CategoryController, customerController controller.ProductController, orderProductController controller.OrderProductController, ordersController controller.OrdersController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customerController.FindAll)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.GET("/api/orderproducts", orderProductController.FindAll)
	router.GET("/api/orderproducts/:orderProductId", orderProductController.FindById)
	router.POST("/api/orderproducts", orderProductController.Create)
	router.PUT("/api/orderproducts/:orderProductId", orderProductController.Update)
	router.DELETE("/api/orderproducts/:orderProductId", orderProductController.Delete)

	router.GET("/api/orders", ordersController.FindAll)
	router.GET("/api/orders/:ordersId", ordersController.FindById)
	router.POST("/api/orders", ordersController.Create)
	router.PUT("/api/orders/:ordersId", ordersController.Update)
	router.DELETE("/api/orders/:ordersId", ordersController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
