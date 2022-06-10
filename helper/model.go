package helper

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}
func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		Address:     customer.Address,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		//CreatedAt:  customer.CreatedAt,
		//UpdatedAt:  customer.UpdatedAt,
	}
}

func ToOrderProductResponse(orderProduct domain.OrderProduct) web.OrderProductResponse {
	return web.OrderProductResponse{
		Id:        orderProduct.Id,
		OrderId:   orderProduct.OrderId,
		ProductId: orderProduct.ProductId,
		Qty:       orderProduct.Qty,
		Price:     orderProduct.Price,
		Amount:    orderProduct.Amount,
		//CreatedAt:  product.CreatedAt,
		//UpdatedAt:  product.UpdatedAt,
	}
}
func ToOrderProductResponses(orderProduct []domain.OrderProduct) []web.OrderProductResponse {
	var orderProductResponses []web.OrderProductResponse
	for _, orderProduct := range orderProduct {
		orderProductResponses = append(orderProductResponses, ToOrderProductResponse(orderProduct))
	}
	return orderProductResponses
}

func ToOrdersResponse(orders domain.Orders) web.OrdersResponse {
	return web.OrdersResponse{
		Id: orders.Id,
		// OrderDate:   orders.OrderDate,
		CustomerId:  orders.CustomerId,
		TotalAmount: orders.TotalAmount,
	}
}
func ToOrdersResponses(orderst []domain.Orders) []web.OrdersResponse {
	var ordersResponses []web.OrdersResponse
	for _, orders := range orderst {
		ordersResponses = append(ordersResponses, ToOrdersResponse(orders))
	}
	return ordersResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
		//CreatedAt:  product.CreatedAt,
		//UpdatedAt:  product.UpdatedAt,
	}
}
func ToProduct(product web.ProductResponse) domain.Product {
	return domain.Product{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
		//CreatedAt:  product.CreatedAt,
		//UpdatedAt:  product.UpdatedAt,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
