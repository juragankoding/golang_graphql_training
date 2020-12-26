package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type brandsRepository struct {
	Conn *sql.DB
}

func NewGenerateBrandsRepository(Conn *sql.DB) domain.BrandsRepository {
	return &brandsRepository{
		Conn: Conn,
	}
}

func (b *brandsRepository) Single(id int) (*domain.Brands, error) {
	queryRows := b.Conn.QueryRow("SELECT * FROM brands WHERE brand_id=?", id)

	var brands domain.Brands

	switch err := queryRows.Scan(&brands.ID, &brands.Name); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return &brands, nil
	}

	return nil, errors.New("no have action on this function")
}

func (b *brandsRepository) All() ([]*domain.Brands, error) {
	var listBrands []*domain.Brands

	query, err := b.Conn.Query("SELECT * FROM brands")

	if err != nil {
		return nil, err
	}

	for query.Next() {
		var brands domain.Brands

		if err := query.Scan(&brands.ID, &brands.Name); err != nil {
			fmt.Print("Failed scan 1 field")
		} else {
			listBrands = append(listBrands, &brands)
		}

	}

	return listBrands, nil
}

func (b *brandsRepository) Insert(brands domain.Brands) (int64, error) {
	statement, err := b.Conn.Prepare("INSERT INTO brands  (brand_name) values (?)")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(brands.Name)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (b *brandsRepository) Update(brands domain.Brands) (int64, error) {

	statement, err := b.Conn.Prepare("UPDATE brands SET brand_name=? where brand_id=?")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(brands.Name, brands.ID)

	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func (b *brandsRepository) Delete(id int) (int64, error) {
	statement, err := b.Conn.Prepare("DELETE FROM brands WHERE brand_id=?")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(id)

	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
