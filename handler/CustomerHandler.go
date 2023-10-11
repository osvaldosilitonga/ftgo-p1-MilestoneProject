package handler


import (
    "database/sql"
    "klepon/entity"
	"klepon/data"
)

func DisplayMenuList(db *sql.DB) ([]entity.MenuItem, error) {
    // Query untuk mengambil data menu dari database
    rows, err := db.Query("SELECT id, name, price FROM menu_items")
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

func MakeOrder(username string, cart *[]entity.MenuItem) error {
    // Implementasi untuk membuat pesanan dan menambahkannya ke keranjang (cart)
    // Anda dapat mengambil item-menu yang dipilih oleh pengguna dan menambahkannya ke cart
    // Selanjutnya, Anda dapat menyimpan pesanan ini ke database
    return nil
}

func ManageCart(cart *[]entity.MenuItem) {
    // Implementasi untuk melihat, mengedit, dan menghapus item dalam keranjang
}

func DisplayOrderHistory(username string) ([]entity.Order, error) {
    // Implementasi untuk menampilkan riwayat pesanan oleh pengguna
    orderHistory, err := data.GetOrderHistory(username)
    if err != nil {
        return nil, err
    }
    return orderHistory, nil
}
