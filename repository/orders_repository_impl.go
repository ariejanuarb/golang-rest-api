package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() OrdersRepository {
	return &OrdersRepositoryImpl{}
}

func (repository *OrdersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "insert into orders(customer_id, total_amount) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.TotalAmount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orders.Id = int(id)
	return orders
}

func (repository *OrdersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "update orders set customer_id = ?, total_amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.TotalAmount, orders.Id)
	helper.PanicIfError(err)

	return orders
}

func (repository *OrdersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orders domain.Orders) {
	SQL := "delete from orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.Id)
	helper.PanicIfError(err)
}

func (repository *OrdersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Orders, error) {
	SQL := "select id, customer_id, total_amount from orders where id = ?" // , order_date
	rows, err := tx.QueryContext(ctx, SQL, ordersId)
	helper.PanicIfError(err)
	defer rows.Close()

	orders := domain.Orders{}
	if rows.Next() {
		err := rows.Scan(&orders.Id, &orders.CustomerId, &orders.TotalAmount) // , &orders.OrderDate
		helper.PanicIfError(err)
		return orders, nil
	} else {
		return orders, errors.New("orders is not found")
	}
}

func (repository *OrdersRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Orders {
	SQL := "select id, customer_id, total_amount from orders" // , order_date
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderest []domain.Orders
	for rows.Next() {
		orders := domain.Orders{}
		err := rows.Scan(&orders.Id, &orders.CustomerId, &orders.TotalAmount) // , &orders.OrderDate
		helper.PanicIfError(err)
		orderest = append(orderest, orders)
	}
	return orderest
}
