package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
	errorJurganKoding "github.com/juragankoding/golang_graphql_training/error"
)

type storeRepo struct {
	Conn *sql.DB
}

func NewGenereateStoreRepository(conn *sql.DB) domain.StoresRepository {
	store := storeRepo{
		Conn: conn,
	}

	return &store
}

func (s *storeRepo) Single(id int) (*domain.Stores, error) {
	query := s.Conn.QueryRow("SELECT * FROM stores WHERE store_id = ?", id)
	var store *domain.Stores

	switch err := query.Scan(store.StoreID, store.StoreName, store.Phone, store.Email, store.Street, store.City, store.State, store.ZipCode); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return store, nil
	}

	return nil, errorJurganKoding.NotActionOnThisFunction
}

func (s *storeRepo) All() ([]*domain.Stores, error) {
	var stores []*domain.Stores

	query, err := s.Conn.Query("SELECT * FROM stores")

	if err != nil {
		return nil, err
	}

	for query.Next() {
		var store *domain.Stores

		switch err := query.Scan(store.StoreID, store.StoreName, store.Phone, store.Email, store.Street, store.City, store.State, store.ZipCode); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			stores = append(stores, store)
		}
	}

	return stores, nil
}
func (s *storeRepo) Insert(stores domain.Stores) (int64, error) {
	statement, err := s.Conn.Prepare("INSERT INTO stores (store_name, phone, email, street, city, state, zip_code) values (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(stores.StoreName, stores.Phone, stores.Email, stores.Street, stores.City, stores.State, stores.ZipCode)

	if err != nil {
		return -1, err
	}

	if lastID, err := result.LastInsertId(); err != nil {
		return -1, err
	} else {
		return lastID, nil
	}
}
func (s *storeRepo) Update(stores domain.Stores) (int64, error) {
	statement, err := s.Conn.Prepare("UPDATE stores SET store_name=?, phone=?, email=?, street=?, city=?, state=?, zip_code=? WHERE store_id = ?")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(stores.StoreName, stores.Phone, stores.Email, stores.Street, stores.City, stores.State, stores.ZipCode, stores.StoreID)

	if err != nil {
		return -1, err
	}

	if rowEffect, err := result.RowsAffected(); err != nil {
		return -1, err
	} else {
		return rowEffect, nil
	}
}
func (s *storeRepo) Delete(id int) (int64, error) {
	statement, err := s.Conn.Prepare("DELETE FROM stores  WHERE store_id = ?")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(id)

	if err != nil {
		return -1, err
	}

	if rowEffect, err := result.RowsAffected(); err != nil {
		return -1, err
	} else {
		return rowEffect, nil
	}
}
