// data/data.go
package data

import (
    "database/sql"
    "klepon/entity"
)

var db *sql.DB

func SetDB(database *sql.DB) {
    db = database
}

// GetMenuList mengambil data menu dari database
func GetMenuList() ([]entity.MenuItem, error) {
    // Query untuk mengambil data menu dari database
    query := "SELECT id, name, price FROM menu_items"
    
    // Eksekusi query
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var menuList []entity.MenuItem

    for rows.Next() {
        var item entity.MenuItem
        if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
            return nil, err
        }
        menuList = append(menuList, item)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return menuList, nil
}

// GetOrderHistory mengambil riwayat pesanan dari database berdasarkan username
func GetOrderHistory(username string) ([]entity.Order, error) {
	// Query untuk mengambil data riwayat pesanan dari database
	query := "SELECT id, menu_id, quantity, order_time FROM orders WHERE username = ?"
	
	// Eksekusi query
	rows, err := db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderHistory []entity.Order

	for rows.Next() {
		var order entity.Order
		if err := rows.Scan(&order.ID, &order.MenuID, &order.Quantity, &order.OrderTime); err != nil {
			return nil, err
		}
		orderHistory = append(orderHistory, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orderHistory, nil
}
// Fungsi-fungsi lainnya untuk operasi database
