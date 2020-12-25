package repository

import (
	"database/sql"
	"errors"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type productsRepository struct {
	Conn *sql.DB
}

func NewGenerateProductsRepository(Conn *sql.DB) domain.ProductsRepository {
	return &productsRepository{
		Conn: Conn,
	}
}

func (a *productsRepository) Single(id int) (*domain.Products, error) {
	queryRow := a.Conn.QueryRow("SELECT * FROM orders WHERE product_id = ?")

	var product *domain.Products

	switch err := queryRow.Scan(product.ProductID, product.ProductName, product.BrandID, product.CategoryID, product.ModelYear, product.ListPrice); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return product, nil
	}

	return nil, errors.New("Nothing action in this function")
}

func (a *productsRepository) All() ([]*domain.Products, error) {
	var listProducts []*domain.Products

	query, err := a.Conn.Query("SELECT * FROM orders")

	if err != nil {
		return listProducts, err
	}

	for query.Next() {
		var product *domain.Products

		switch err := query.Scan(product.ProductID, product.ProductName, product.BrandID, product.CategoryID, product.ModelYear, product.ListPrice); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			listProducts = append(listProducts, product)
		}
	}

	return listProducts, nil
}

func (a *productsRepository) Insert(products domain.Products) (int64, error) {
	statement, err := a.Conn.Prepare("INSERT INTO products (product_name, brands_id, category_id, model_year, list_price) values (?,?,?,?,?)")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(products.ProductName, products.BrandID, products.CategoryID, products.ModelYear, products.ListPrice); err != nil {
		return -1, err
	} else {
		return result.LastInsertId()
	}

}

func (a *productsRepository) Update(products domain.Products) (int64, error) {
	statement, err := a.Conn.Prepare("UPDATE products SET product_name=?, brand_id=?, category_id=?, model_year=?, list_price=? WHERE product_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(products.ProductName, products.BrandID, products.CategoryID, products.ModelYear, products.ListPrice, products.ProductID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}

func (a *productsRepository) Delete(id int) (int64, error) {
	statement, err := a.Conn.Prepare("DELETE FROM products WHERE product_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(id); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
