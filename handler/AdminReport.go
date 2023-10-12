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
func OrderReport() {}
func MenuReport()  {}
