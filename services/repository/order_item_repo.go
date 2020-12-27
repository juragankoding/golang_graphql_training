package repository

import (
	"database/sql"
	"errors"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type orderItemRepository struct {
	Conn *sql.DB
}

func NewGenerateOrderItemRepository(Conn *sql.DB) domain.OrderItemRepository {
	return &orderItemRepository{
		Conn: Conn,
	}
}

func (o *orderItemRepository) All() ([]*domain.OrderItem, error) {
	var listOrderItem []*domain.OrderItem

	query, err := o.Conn.Query("SELECT * FROM order_items")

	if err != nil {
		return nil, err
	}

	for query.Next() {
		var orderItem *domain.OrderItem

		switch err := query.Scan(orderItem.ItemID, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.ListPrice, orderItem.Discount); err {
		case sql.ErrNoRows:
			return nil, sql.ErrNoRows
		case nil:
			listOrderItem = append(listOrderItem, orderItem)
		}
	}

	return listOrderItem, nil
}

func (o *orderItemRepository) Single(id int) (*domain.OrderItem, error) {
	queryRow := o.Conn.QueryRow("SELECT * FROM order_items WHERE item_id=?", id)

	var orderItem *domain.OrderItem

	switch err := queryRow.Scan(orderItem.ItemID, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.ListPrice, orderItem.Discount); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return orderItem, nil
	}

	return nil, errors.New("Nothing action in this functions")
}

func (o *orderItemRepository) Insert(orderItem domain.OrderItem) (int64, error) {
	statement, err := o.Conn.Prepare("INSERT INTO order_items (order_id, product_id, quantity, list_price, discount) VALUES (?,?,?,?,?)")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.ListPrice, orderItem.Discount); err != nil {
		return -1, err
	} else {
		return result.LastInsertId()
	}
}

func (o *orderItemRepository) Update(orderItem domain.OrderItem) (int64, error) {
	statement, err := o.Conn.Prepare("UPDATE order_items SET order_id=?, product_id=?, quantity=?, list_price=?, discount=? WHERE item_id=?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.ListPrice, orderItem.Discount, orderItem.ItemID); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}

func (o *orderItemRepository) Delete(id int) (int64, error) {
	statement, err := o.Conn.Prepare("DELETE FROM order_items WHERE item_id = ?")

	if err != nil {
		return -1, err
	}

	if result, err := statement.Exec(id); err != nil {
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
