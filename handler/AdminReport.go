package handler

import (
	// "context"
	"context"
	"database/sql"
	"fmt"
	"klepon/config"
	"klepon/entity"
	"log"
	"time"
)

func GeneralReport() (entity.GeneralReport, bool) {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT 
	COUNT(CASE orders.status WHEN "Success" THEN 1 ELSE NULL END) AS "Success", 
	COUNT(CASE orders.status WHEN "Cancel" THEN 1 ELSE NULL END) AS "Cancel", 
	(
		SELECT SUM(amount) FROM orders 
		WHERE status = "Success"
		AND order_date >= DATE_FORMAT(NOW(), '%Y-%m-01') AND order_date <= NOW()
	) AS "Revenue"
	FROM orders
	WHERE order_date >= DATE_FORMAT(NOW(), '%Y-%m-01')
	AND order_date <= NOW();
	`

	var gr entity.GeneralReport

	row := db.QueryRowContext(ctx, query)
	err = row.Scan(&gr.Success, &gr.Cancel, &gr.Revenue)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return gr, false
	case nil:
		return gr, true
	default:
		panic(err)
	}
}

func MenuReport() ([]entity.MenuReport, error) {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT menu_id AS "ID", menu.nama AS "Menu", menu.category AS "Category", COUNT(*) AS Success
	FROM order_details
	JOIN menu ON order_details.menu_id = menu.id
	JOIN orders ON order_details.order_id = orders.id
	WHERE order_date >= DATE_FORMAT(NOW(), '%Y-%m-01')
	GROUP BY menu_id
	ORDER BY Success DESC;
	`

	var menuReport []entity.MenuReport

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mr entity.MenuReport

		err = rows.Scan(&mr.ID, &mr.Menu, &mr.Category, &mr.OrderSuccess)
		if err != nil {
			return nil, err
		}

		menuReport = append(menuReport, mr)
	}

	return menuReport, nil
}

func CustomerReport() ([]entity.CustomerReport, error) {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT 
		users.username AS "Username", 
		COUNT(CASE orders.status WHEN "Success" THEN 1 ELSE NULL END) AS "Success",
		COUNT(CASE orders.status WHEN "Cancel" THEN 1 ELSE NULL END) AS "Cancel"
	FROM orders
	JOIN users ON orders.user_id = users.id
	WHERE order_date >= DATE_FORMAT(NOW(), '%Y-%m-01')
	GROUP BY users.username
	ORDER BY users.username;
	`

	var customerReport []entity.CustomerReport

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cr entity.CustomerReport

		err = rows.Scan(&cr.Username, &cr.OrderSuccess, &cr.OrderCancel)
		if err != nil {
			return nil, err
		}

		customerReport = append(customerReport, cr)
	}

	return customerReport, nil
}
