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
	var listStaffs []*domain.Staffs

	query, err := s.Conn.Query("SELECT * FROM staffs")

	if err != nil {
		return nil, err
	}

	for query.Next() {
		var staff domain.Staffs

		switch err := query.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName, &staff.Email, &staff.Phone, &staff.Active, &staff.StoreID, &staff.ManagerID); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			listStaffs = append(listStaffs, &staff)
		default:
			return nil, err
		}
	}

	return listStaffs, nil
}

func (s *staffsRepository) Single(id int) (*domain.Staffs, error) {
	query := s.Conn.QueryRow("SELECT * FROM staffs where staff_id=?", id)

	var staff domain.Staffs

	switch err := query.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName, &staff.Email, &staff.Phone, &staff.Active, &staff.StoreID, &staff.ManagerID); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return &staff, nil
	default:
		return nil, err
	}
}

func (s *staffsRepository) Insert(staffs domain.Staffs) (int64, error) {
	statement, err := s.Conn.Prepare("INSERT INTO staffs (first_name, last_name, email, phone, active, store_id, manager_id) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(staffs.FirstName, staffs.LastName, staffs.Email, staffs.Phone, staffs.Active, staffs.StoreID, staffs.ManagerID); err != nil {
		return -1, err
	} else {
		return result.LastInsertId()
	}
}

func (s *staffsRepository) Update(staffs domain.Staffs) (int64, error) {
	statement, err := s.Conn.Prepare("UPDATE staffs SET first_name=?, last_name=?, email=?, phone=?, active=?, store_id=?, manager_id=? WHERE staff_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(staffs.FirstName, staffs.LastName, staffs.Email, staffs.Phone, staffs.Active, staffs.StoreID, staffs.ManagerID, staffs.StaffID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}

func (s *staffsRepository) Delete(id int) (int64, error) {
	statement, err := s.Conn.Prepare("DELETE FROM staffs WHERE staff_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(id); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
