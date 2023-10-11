package handler

import (
	"context"
	"klepon/config"
	"klepon/entity"
	"log"
	"time"
)

func GetPaymentOrders() ([]entity.AdminOrderList, error) {
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
		SELECT orders.id, users.username, orders.order_date, menu.nama, menu.harga, order_details.qty, orders.discount FROM orders
		LEFT JOIN order_details ON order_details.order_id = orders.id
		JOIN users ON orders.user_id = users.id
		JOIN menu ON order_details.menu_id = menu.id
		WHERE orders.status = "Payment"
		ORDER BY orders.order_date DESC
	`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o entity.AdminOrderList

		err := rows.Scan(&o.ID, &o.Username, &o.OrderDate, &o.Menu, &o.Price, &o.Qty, &o.Discount)
		if err != nil {
			return nil, err
		}

		orderList = append(orderList, o)
	}

	return orderList, nil
}

func PaymentProcess(id, discount, amount int) string {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
				UPDATE orders
				SET status = "Success", discount = ?, amount = ?
				WHERE id = ?
			`
	_, err = db.ExecContext(ctx, query, discount, amount, id)
	if err != nil {
		return "Payment Failed."
	}

	return "Payment Success."
}
