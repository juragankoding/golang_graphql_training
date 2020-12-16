package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type staffsRepository struct {
	Conn *sql.DB
}

func NewGenerateStaffsRepository(conn *sql.DB) domain.StaffsRepository {
	return &staffsRepository{
		Conn: conn,
	}
}

func (s *staffsRepository) All() ([]*domain.Staffs, error) {
	return nil, nil
}

func (s *staffsRepository) Single(id int) (*domain.Staffs, error) {
	return nil, nil
}

func (s *staffsRepository) Insert(staffs domain.Staffs) (int64, error) {
	return -1, nil
}

func (s *staffsRepository) Update(staffs domain.Staffs) (int64, error) {
	return -1, nil
}

func (s *staffsRepository) Delete(id int) (int64, error) {
	return -1, nil
}
