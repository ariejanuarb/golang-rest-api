package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	logrus.Info("product Create repository_impl start")
	SQL := "insert into product(name, price, category_id) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	logrus.Info("product Create repository_impl end")
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	logrus.Info("product Update repository_impl start")
	SQL := "update product set name = ?, price = ?, category_id = ? where id = ?"                   // created_at = current_timestamp, updated_by = current_timestamp
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId, product.Id) // &product.CreatedAt, &product.UpdatedAt
	helper.PanicIfError(err)

	logrus.Info("product Update repository_impl start")
	return product

}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	logrus.Info("product Delete repository_impl start")
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
	logrus.Info("product Delete repository_impl end")
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (web.ProductResponse, error) {
	logrus.Info("product FindById repository_impl start")
	SQL := "SELECT product.id, product.name, product.price, product.category_id, category.name as category_name FROM product INNER JOIN category on product.category_id = category.id where product.id = ?"
	//SQL := "select id, name, price, category_id from product where id = ?" // , created_at, updated_at
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := web.ProductResponse{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName) // &product.CategoryName, &product.CreatedAt, &product.UpdatedAt
		helper.PanicIfError(err)
		logrus.Info("product FindById repository_impl end")
		return product, nil
	} else {
		logrus.Info("product FindById repository_impl end")
		return product, errors.New("product is not found")
	}
}

//func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
//
//	//SQL := "SELECT product.id, product.name, product.price, product.category_id, category.name as category_name FROM product INNER JOIN category on product.category_id = category.id where product.id = ?"
//	SQL := "select id, name, price, category_id from product where id = ?
//	rows, err := tx.QueryContext(ctx, SQL, productId)
//	helper.PanicIfError(err)
//	defer rows.Close()
//
//	product := domain.Product{}
//	if rows.Next() {
//		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Id) // &product.CategoryName, &product.CreatedAt, &product.UpdatedAt
//		helper.PanicIfError(err)
//		logrus.Info("product repository end")
//		return product, nil
//	} else {
//		logrus.Info("product repository end")
//		return product, errors.New("product is not found")
//	}
//}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse {
	//select id, name, price, category_id from product // FindAll SQL syntax
	//SQL := "SELECT product.* category.name as category_name FROM product INNER JOIN category on product.category_id = category.id"
	//SQL := "SELECT p.id, p.name, p.price, p.category_id c.name AS category_name FROM product p INNER JOIN category c ON p.category_id = c.id"
	logrus.Info("product FindAll repository_impl start")
	SQL := "SELECT product.id, product.name, product.price, product.category_id, category.name as category_name FROM product INNER JOIN category on product.category_id = category.id" // , created_at, updated_at
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []web.ProductResponse
	for rows.Next() {
		product := web.ProductResponse{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName) // , &product.CreatedAt, &product.UpdatedAt
		helper.PanicIfError(err)
		products = append(products, product)
	}
	logrus.Info("product FindAll repository_impl start")
	return products
}
