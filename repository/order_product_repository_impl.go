package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrderProductRepositoryImpl struct {
}

func NewOrderProductRepository() OrderProductRepository {
	return &OrderProductRepositoryImpl{}
}

func (repository *OrderProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orderProduct domain.OrderProduct) domain.OrderProduct {
	SQL := "insert into order_product(order_id, product_id, qty, price, amount) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, orderProduct.OrderId, orderProduct.ProductId, orderProduct.Qty, orderProduct.Price, orderProduct.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orderProduct.Id = int(id)
	return orderProduct
}

func (repository *OrderProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orderProduct domain.OrderProduct) domain.OrderProduct {
	SQL := "update order_product set order_id = ?, product_id = ?, qty = ?, price = ?, amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orderProduct.OrderId, orderProduct.ProductId, orderProduct.Qty, orderProduct.Price, orderProduct.Amount, orderProduct.Id)
	helper.PanicIfError(err)

	return orderProduct
}

func (repository *OrderProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderProduct domain.OrderProduct) {
	SQL := "delete from order_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orderProduct.Id)
	helper.PanicIfError(err)
}

func (repository *OrderProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderProductId int) (domain.OrderProduct, error) {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderProductId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderProduct := domain.OrderProduct{}
	if rows.Next() {
		err := rows.Scan(&orderProduct.Id, &orderProduct.OrderId, &orderProduct.ProductId, &orderProduct.Qty, &orderProduct.Price, &orderProduct.Amount)
		helper.PanicIfError(err)
		return orderProduct, nil
	} else {
		return orderProduct, errors.New("order product is not found")
	}
}

func (repository *OrderProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderProducts []domain.OrderProduct
	for rows.Next() {
		orderProduct := domain.OrderProduct{}
		err := rows.Scan(&orderProduct.Id, &orderProduct.OrderId, &orderProduct.ProductId, &orderProduct.Qty, &orderProduct.Price, &orderProduct.Amount)
		helper.PanicIfError(err)
		orderProducts = append(orderProducts, orderProduct)
	}
	return orderProducts
}
