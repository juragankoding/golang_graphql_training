package repository

import (
	"database/sql"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type customerRepository struct {
	Conn *sql.DB
}

func NewGenerateCustomersRepository(Conn *sql.DB) domain.CustomersRepository {
	return &customerRepository{
		Conn: Conn,
	}
}

func (a *customerRepository) All() ([]*domain.Customers, error) {
	var customers []*domain.Customers

	query, err := a.Conn.Query("SELECT * FROM customers")

	if err != nil {
		return nil, err
	}

	for query.Next() {
		var customer domain.Customers

		switch err := query.Scan(&customer.CustomerID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Phone,
			&customer.Email,
			&customer.Street,
			&customer.City,
			&customer.State,
			&customer.ZipCode); err {
		case sql.ErrNoRows:
			return nil, err

		case nil:
			customers = append(customers, &customer)
		}
	}

	return customers, nil
}

func (a *customerRepository) Get(id int) (*domain.Customers, error) {

	queryRow := a.Conn.QueryRow("SELECT * FROM customers WHERE customer_id = ?", id)

	var customer domain.Customers

	switch err := queryRow.Scan(&customer.CustomerID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&customer.Email,
		&customer.Street,
		&customer.City,
		&customer.State,
		&customer.ZipCode); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return &customer, nil
	default:
		return nil, err
	}
}

func (a *customerRepository) Insert(customer domain.Customers) (int64, error) {
	statement, err := a.Conn.Prepare("INSERT INTO customers (first_name, last_name, phone, email, street, city, zip_code) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(customer.FirstName, customer.LastName, customer.Phone, customer.Email, customer.Street, customer.City, customer.ZipCode); err != nil {
		return -1, err
	} else {
		return result.LastInsertId()
	}

}

func (a *customerRepository) Update(customers domain.Customers) (int64, error) {
	statement, err := a.Conn.Prepare("UPDATE customers SET first_name=?,last_name=?,phone=?,email=?,street=?,city=?,state=?,zip_code=? where customer_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(customers.FirstName, customers.LastName, customers.Phone, customers.Email, customers.Street, customers.City, customers.State, customers.ZipCode, customers.CustomerID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}

}

func (a *customerRepository) Delete(id int) (int64, error) {
	statement, err := a.Conn.Prepare("DELETE FROM customers WHERE customer_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(id); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
