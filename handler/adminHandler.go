package handler

import (
	"context"
	"database/sql"
	"klepon/config"
	"klepon/entity"
	"log"
	"time"
)

func GetOrders() ([]entity.AdminOrderList, error) {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	// Orders Object
	var orderList []entity.AdminOrderList

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT orders.id, users.username, orders.order_date, menu.nama, order_details.qty 
		FROM orders
		LEFT JOIN order_details ON order_details.order_id = orders.id
		JOIN users ON orders.user_id = users.id
		JOIN menu ON order_details.menu_id = menu.id
		WHERE orders.status = "waiting"
		ORDER BY orders.order_date DESC;
	`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o entity.AdminOrderList

		err := rows.Scan(&o.ID, &o.Username, &o.OrderDate, &o.Menu, &o.Qty)
		if err != nil {
			return nil, err
		}

		orderList = append(orderList, o)
	}

	return orderList, nil
}

func ProcessOrder(id int) string {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var orderId int

	// Check if order id exist
	query := `
		SELECT id FROM orders
		WHERE id = ?
	`
	row := db.QueryRowContext(ctx, query, id)

	switch err := row.Scan(&orderId); err {
	// order id not exist
	case sql.ErrNoRows:
		return "Orders Id not found."

	// order id exist
	case nil:
		// Updating order status to 'Payment'
		query := `
				UPDATE orders
				SET status = "Payment"
				WHERE id = ?
			`
		_, err := db.ExecContext(ctx, query, orderId)
		if err != nil {
			return "Update order status failed."
		}

		return "Order status has been updated."

	default:
		panic(err)
	}
}
