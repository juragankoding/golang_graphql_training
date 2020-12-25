package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type categoriesRepository struct {
	Conn *sql.DB
}

func NewGenerateCategoriesRepository(Conn *sql.DB) domain.CategoriesRepository {
	return &categoriesRepository{
		Conn: Conn,
	}
}

func (r *categoriesRepository) Get(id int) (res *domain.Categories, err error) {
	row := r.Conn.QueryRow("SELECT * FROM categories WHERE category_id = ?", id)

	var categories domain.Categories

	switch err := row.Scan(&categories.ID, &categories.Name); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return &categories, nil
	}

	return nil, errors.New("no have data")
}

func (r *categoriesRepository) Fetch() (res []*domain.Categories, err error) {

	var listData []*domain.Categories

	rows, err := r.Conn.Query("SELECT * FROM categories")

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	for rows.Next() {
		var categories domain.Categories

		rows.Scan(&categories.ID, &categories.Name)

		listData = append(listData, &categories)
	}

	return listData, nil
}

func (r *categoriesRepository) Insert(categories *domain.Categories) (int64, error) {
	statement, err := r.Conn.Prepare("INSERT INTO categories (category_name) VALUES (?)")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(categories.Name)

	if err != nil {
		return -1, err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return lastID, nil

}
func (r *categoriesRepository) Update(categories *domain.Categories) (int64, error) {
	statement, err := r.Conn.Prepare("UPDATE categories set category_name=? where category_id=?")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(categories.Name, categories.ID)

	if err != nil {
		return -1, err
	}

	lastID, err := result.RowsAffected()

	if err != nil {
		return -1, err
	}

	return lastID, nil
}
func (r *categoriesRepository) Delete(id int) (int64, error) {
	statement, err := r.Conn.Prepare("DELETE FROM categories WHERE category_id=?")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(id)

	if err != nil {
		return -1, err
	}

	lastID, err := result.RowsAffected()

	if err != nil {
		return -1, err
	}

	return lastID, nil
}
