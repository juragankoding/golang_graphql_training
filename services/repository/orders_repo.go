package repository

import (
	"database/sql"
	"errors"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type ordersRepository struct {
	Conn *sql.DB
}

func NewGenerateOrdersRepository(Conn *sql.DB) domain.OrdersRepository {
	return &ordersRepository{
		Conn: Conn,
	}
}

func (o *ordersRepository) All() ([]*domain.Orders, error) {
	var listListOrders []*domain.Orders

	query, err := o.Conn.Query("SELECT * FROM orders ")

	if err != nil {
		return nil, err
	}

	for query.Next() {
		var order *domain.Orders

		switch err := query.Scan(order.OrderID, order.CustomerID, order.OrderStatus, order.OrderDate, order.RequiredDate, order.ShippedDate, order.StoreID, order.StaffID); err {
		case sql.ErrNoRows:
			return nil, sql.ErrNoRows
		case nil:
			listListOrders = append(listListOrders, order)
		}
	}

	return listListOrders, nil
}

func (o *ordersRepository) Single(id int) (*domain.Orders, error) {
	queryRow := o.Conn.QueryRow("SELECT * FROM orders WHERE order_id = ?", id)

	var order *domain.Orders

	switch err := queryRow.Scan(order.OrderID, order.CustomerID, order.OrderStatus, order.OrderDate, order.RequiredDate, order.ShippedDate, order.StoreID, order.StaffID); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return order, nil
	}

	return nil, errors.New("Nothing action in this function")
}

func (o *ordersRepository) Insert(orders domain.Orders) (int64, error) {
	statement, err := o.Conn.Prepare("INSERT INTO orders (costumer_id, order_status, order_date, required_date, shipped_date, store_id, staff_id) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(orders.CustomerID, orders.OrderStatus, orders.OrderDate, orders.RequiredDate, orders.ShippedDate, orders.StoreID, orders.StaffID); err != nil {
		return -1, err
	} else {
		return result.LastInsertId()
	}
}

func (o *ordersRepository) Update(orders domain.Orders) (int64, error) {
	statement, err := o.Conn.Prepare("UPDATE orders SET costumer_id=?, order_status=?, order_date=?, required_date=?, shipped_date=?, store_id=?, staff_id=? WHERE order_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(orders.CustomerID, orders.OrderStatus, orders.OrderDate, orders.RequiredDate, orders.ShippedDate, orders.StoreID, orders.StaffID, orders.OrderID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
func (o *ordersRepository) Delete(id int) (int64, error) {
	statement, err := o.Conn.Prepare("DELETE FROM orders WHERE order_id = ?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(id); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}

}
