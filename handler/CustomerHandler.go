package handler

import (
	"bufio"
	"fmt"
	"klepon/config"
	"strconv"
	"strings"

	"klepon/entity"
	"log"
	"os"
)

func DisplayMenuList() ([]entity.MenuItem, error) {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	// Query untuk mengambil data menu dari database
	rows, err := db.Query("SELECT id, nama, harga FROM menu")
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

func MakeOrder(username string, cart *[]entity.MenuItem, status string) error {
	// Koneksi ke DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Gagal terhubung ke DB", err)
		return err
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// var menuID, quantity int

		fmt.Print("Masukkan ID menu yang ingin dipesan (0 untuk selesai): ")
		scanner.Scan()
		menuID, err := strconv.Atoi(scanner.Text())
		// fmt.Scan(&menuID)

		// Keluar dari loop jika pengguna memasukkan 0
		if menuID == 0 {
			return nil
		}
		if menuID == 0 && len(*cart) > 0 {
			break
		}

		fmt.Print("Masukkan jumlah: ")
		scanner.Scan()
		quantity, err := strconv.Atoi(scanner.Text())
		// fmt.Scan(&quantity)

		// Query item menu berdasarkan ID yang dimasukkan
		var menuItem entity.MenuItem
		menuItem.Qty = quantity
		err = db.QueryRow("SELECT id, nama, harga FROM menu WHERE id = ?", menuID).Scan(&menuItem.ID, &menuItem.Name, &menuItem.Price)
		if err != nil {
			log.Printf("Gagal mengambil item menu: %v\n", err)
			return err
		}

		// Tambahkan item menu yang dipilih ke dalam keranjang berdasarkan jumlah yang dimasukkan
		// for i := 0; i < quantity; i++ {
		// 	*cart = append(*cart, menuItem)
		// }
		*cart = append(*cart, menuItem)

		fmt.Printf("%s ditambahkan ke dalam keranjang Anda.\n", menuItem.Name)
	}

	// Set status pesanan secara default

	// Masukkan item ke dalam tabel 'cart'
	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		log.Printf("Gagal mendapatkan ID pengguna: %v\n", err)
		return err
	}

	for _, item := range *cart {
		_, err := db.Exec("INSERT INTO cart (user_id, menu_id, qty) VALUES (?, ?, 1)", userID, item.ID)
		if err != nil {
			log.Printf("Gagal memasukkan item ke dalam keranjang: %v\n", err)
			return err
		}
	}

	fmt.Println("Order successfully created and added to your cart.")

	return nil
}

func ManageCart(username string, cart *[]entity.MenuItem, defaultStatus string) {

	for {
		// Tampilkan isi keranjang
		fmt.Println("Your Cart:")
		for i, item := range *cart {
			fmt.Printf("%d. ID: %d, Name: %s, Price: %.2f, Qty: %v\n", i+1, item.ID, item.Name, item.Price, item.Qty)
		}

		// Berikan opsi kepada pengguna
		fmt.Println("Cart Menu:")
		fmt.Println("1. Edit Cart")
		fmt.Println("2. Delete Item")
		fmt.Println("3. Submit Order")
		fmt.Println("4. Back")

		// Membaca pilihan pengguna
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Select Menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			// Panggil fungsi untuk mengedit cart
			EditCart(cart)
		case "2":
			// Panggil fungsi untuk menghapus item dari cart
			DeleteItemFromCart(username, cart)
		case "3":
			// Panggil fungsi untuk menyubmit pesanan
			orderID, err := SubmitOrder(username, cart, defaultStatus)
			if err != nil {
				fmt.Printf("Failed to submit order: %v\n", err)
			} else {
				fmt.Printf("Order successfully submitted. Order ID: %d\n", orderID)
			}
			// Kembalikan pengguna ke menu utama
			return

		case "4":
			// Kembali ke menu utama
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func GetOrderID(username string) (int, error) {
	// Connect to the database
	db, err := config.ConnDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Query the database to get the order ID based on the username
	var orderID int
	err = db.QueryRow("SELECT id FROM orders WHERE user_id = ?", username).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func EditCart(cart *[]entity.MenuItem) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Println("Edit Cart")
	fmt.Println()

	// Tampilkan isi keranjang saat ini
	fmt.Println("Your Cart:")
	for i, item := range *cart {
		fmt.Printf("%d. ID: %d, Name: %s, Price: %.2f, Qty: %v\n", i+1, item.ID, item.Name, item.Price, item.Qty)
	}

	// Mintalah pengguna memasukkan nomor item yang ingin diedit
	fmt.Print("Enter the Item ID you want to edit (0 to finish): ")
	scanner.Scan()
	itemNumber, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	// var itemNumber int
	// fmt.Scan(&itemNumber)

	if itemNumber == 0 {
		return // Keluar dari Edit Cart jika pengguna memilih 0
	}

	if itemNumber < 1 || itemNumber > len(*cart) {
		fmt.Println("Invalid item number. Please try again.")
		return
	}

	// Mintalah pengguna memasukkan informasi item yang baru
	// var newQuantity int
	fmt.Print("Enter the new quantity: ")
	scanner.Scan()
	newQuantity, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	// fmt.Scan(&newQuantity)

	// Update jumlah item dalam keranjang
	for i, v := range *cart {
		if v.ID == itemNumber {
			(*cart)[i].Qty = newQuantity
		}
	}

	//(*cart)[itemNumber-1].Quantity = newQuantity

	fmt.Println("Item updated successfully.")
	fmt.Println()
}

func DeleteItemFromCart(username string, cart *[]entity.MenuItem) {
	fmt.Println("Delete Item from Cart")

	// Tampilkan isi keranjang saat ini
	fmt.Println("Your Cart:")
	for i, item := range *cart {
		fmt.Printf("%d. ID: %d, Name: %s, Price: %.2f\n", i+1, item.ID, item.Name, item.Price)
	}

	// Mintalah pengguna memasukkan nomor item yang ingin dihapus
	fmt.Print("Enter the item number you want to delete (0 to finish): ")
	var itemNumber int
	fmt.Scan(&itemNumber)

	if itemNumber == 0 {
		return // Keluar dari Delete Item jika pengguna memilih 0
	}

	if itemNumber < 1 || itemNumber > len(*cart) {
		fmt.Println("Invalid item number. Please try again.")
		return
	}

	// Hapus item dari keranjang
	*cart = append((*cart)[:itemNumber-1], (*cart)[itemNumber:]...)

	fmt.Println("Item deleted successfully.")
}

func SubmitOrder(username string, cart *[]entity.MenuItem, status string) (int, error) {
	// Connect to the database
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
		return 0, err
	}
	defer db.Close()

	// Get the user_id from the username
	var user_id int
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&user_id)
	if err != nil {
		log.Printf("Failed to get user_id: %v\n", err)
		return 0, err
	}

	// Calculate the total amount from the items in the cart
	// totalAmount := 0.0
	// for _, item := range *cart {
	// 	totalAmount += item.Price
	// }

	// Begin a database transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin transaction", err)
		return 0, err
	}

	// Insert the order into the 'orders' table, including the total amount
	result, err := tx.Exec("INSERT INTO orders (user_id, order_date, status) VALUES (?, NOW(), ?)", user_id, status)
	if err != nil {
		tx.Rollback() // Rollback the transaction on error
		log.Printf("Failed to insert order into orders table: %v\n", err)
		return 0, err
	}

	// Get the ID of the inserted order
	orderID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Printf("Failed to get order ID: %v\n", err)
		return 0, err
	}

	// Insert items into the 'order_details' table
	for _, item := range *cart {
		_, err := tx.Exec("INSERT INTO order_details (order_id, menu_id, qty) VALUES (?, ?, ?)", orderID, item.ID, item.Qty)
		if err != nil {
			tx.Rollback() // Rollback the transaction on error
			log.Printf("Failed to insert order details: %v\n", err)
			return 0, err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Failed to commit transaction", err)
		return 0, err
	}

	// Setelah transaksi COMMIT berhasil, hapus data dari tabel cart
	_, err = db.Exec("DELETE FROM cart WHERE user_id = ?", user_id)
	if err != nil {
		log.Printf("Failed to delete cart items: %v\n", err)
		return 0, err
	}

	return int(orderID), nil
}

func GetLastOrderID() (int, error) {
	// Connect to the database
	db, err := config.ConnDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Query the database to get the last inserted order ID
	var orderID int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&orderID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func DisplayOrderHistory(username string) ([]entity.Order, error) {
	fmt.Println("============== Order Customer History ==============")
	fmt.Println()
	// Connect to the database
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	// Query to retrieve order history for the specified username
	query := `
    SELECT
        o.id AS order_id,
        u.username AS customer,
        m.nama AS menu_name,
        m.harga AS total,
        o.status AS status
    FROM
        orders o
        JOIN users u ON o.user_id = u.id
        JOIN order_details od ON o.id = od.order_id
        JOIN menu m ON od.menu_id = m.id
    WHERE
        u.username = ?;
    `

	// Execute the query and scan the results into a slice of Order objects
	rows, err := db.Query(query, username)
	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var orderHistory []entity.Order

	fmt.Printf("%-10s | %-10s | %-13s | %-10s\n", "Order ID", "Customer", "Total", "Status")
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.Customer, &order.Menu, &order.TotalAmount, &order.Status)
		if err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			return nil, err
		}

		orderHistory = append(orderHistory, order)

		fmt.Printf("%-10d | %-10s | Rp.%-10.2f | %-10s\n", order.ID, order.Customer, order.TotalAmount, order.Status)
	}

	return orderHistory, nil
}
