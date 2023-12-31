package cli

import (
	"bufio"
	"fmt"
	"os"

	//"klepon/entity"
	"klepon/entity"
	"klepon/handler"
)

func UserPage(username string) {
	var cart []entity.MenuItem
	defaultStatus := "Waiting"

	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println()
		fmt.Printf("Welcome, %s!\n", username)

		fmt.Println("User Menu:")
		fmt.Println("1. Menu List")
		fmt.Println("2. Make Order")
		fmt.Println("3. Cart")
		fmt.Println("4. Order History")
		fmt.Println("5. Logout")

		fmt.Print("Select Menu: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println()
			menuList, err := handler.DisplayMenuList()
			if err != nil {
				fmt.Printf("Failed to display menu list: %v\n", err)
			} else {
				// Tampilkan daftar menu ke pengguna
				fmt.Println("Menu List:")
				for _, item := range menuList {
					fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
				}
			}
		case "2":
			err := handler.MakeOrder(username, &cart, defaultStatus)
			if err != nil {
				fmt.Printf("Failed to make an order: %v\n", err)
				return
			}
			// else {
			// 	fmt.Println("Order successfully created and added to your cart.")

			// }

		case "3":
			fmt.Println()
			handler.ManageCart(username, &cart, defaultStatus)
		case "4":
			orderHistory, err := handler.DisplayOrderHistory(username)
			if err != nil {
				fmt.Printf("Failed to display order history: %v\n", err)
			} else {
				// Tampilkan riwayat pesanan ke pengguna
				fmt.Println("Order History:")
				for _, order := range orderHistory {
					fmt.Printf("Order ID: %d, Customer: %s, Total: %.2f, Status: %s\n", order.ID, order.Customer, order.Total, order.Status)
				}
			}
		case "5":
			fmt.Println("Logging out.")
			return
		default:
			if choice != "0" {
				break
			}
		}
	}
}
