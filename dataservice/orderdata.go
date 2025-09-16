package dataservice

import (
	"database/sql"
	"errors"

	"github.com/varlaalekya/goproject/model"
)

// Create
func CreateOrder(db *sql.DB, order model.Order) error {
	const q = `
		INSERT INTO orders
			(id, customer_name, paymentmethod, placedondate, deliveredon, item, address, amount)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(q,
		order.Id,
		order.Customer_Name,
		order.PaymentMethod,
		order.PlacedOnDate,
		order.DeliveredOn,
		order.Item,
		order.Address,
		order.Amount,
	)
	return err
}

// Update: expects full order with Id set (simple & clear for your task)
func UpdateOrder(db *sql.DB, order model.Order) error {
	if order.Id == 0 {
		return errors.New("id is required")
	}
	const q = `
		UPDATE orders
		   SET customer_name = ?,
		       paymentmethod = ?,
		       placedondate  = ?,
		       deliveredon   = ?,
		       item          = ?,
		       address       = ?,
		       amount        = ?
		 WHERE id = ?`
	res, err := db.Exec(q,
		order.Customer_Name,
		order.PaymentMethod,
		order.PlacedOnDate,
		order.DeliveredOn,
		order.Item,
		order.Address,
		order.Amount,
		order.Id,
	)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err == nil && n == 0 {
		return errors.New("order not found")
	}
	return err
}

// Delete by id
func DeleteOrder(db *sql.DB, id int) error {
	const q = `DELETE FROM orders WHERE id = ?`
	res, err := db.Exec(q, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err == nil && n == 0 {
		return errors.New("order not found")
	}
	return err
}
