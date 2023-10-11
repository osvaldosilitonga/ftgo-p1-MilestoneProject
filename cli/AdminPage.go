package cli

import (
	"fmt"
	"klepon/data"
	"klepon/entity"
	"klepon/handler"
	"strings"
	"bufio"
	"os"
)

func AdminPage(username string) {
	fmt.Printf("Welcome, Admin %s!\n", username)
	loggedInUser := data.GetUserByUsername(username)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nAdmin Menu:")
		fmt.Println("1. Add Menu Item")
		fmt.Println("2. Delete Menu Item")
		fmt.Println("3. Update Menu Item")
		fmt.Println("4. Process Order")
		fmt.Println("5. Cancel Order")
		fmt.Println("6. General Report")
		fmt.Println("7. Get Order Success by Date")
		fmt.Println("8. Get Cancel Order by Date")
		fmt.Println("9. Logout")
		fmt.Print("Select Menu: ")

		scanner.Scan()
		choice := scanner.Text()
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			AddMenuItem()
		case "2":
			DeleteMenuItem()
		case "3":
			UpdateMenuItem()
		case "4":
			ProcessOrder()
		case "5":
			CancelOrder()
		case "6":
			GenerateGeneralReport()
		case "7":
			GetOrderSuccessByDate()
		case "8":
			GetCancelOrderByDate()
		case "9":
			fmt.Println("Logging out.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// func AddMenuItem() {
// 	// Implement the logic to add a menu item here
// }

// func DeleteMenuItem() {
// 	// Implement the logic to delete a menu item here
// }

// func UpdateMenuItem() {
// 	// Implement the logic to update a menu item here
// }

// func ProcessOrder() {
// 	// Implement the logic to process an order here
// }

// func CancelOrder() {
// 	// Implement the logic to cancel an order here
// }

// func GenerateGeneralReport() {
// 	// Implement the logic to generate a general report here
// }

// func GetOrderSuccessByDate() {
// 	// Implement the logic to get order success by date here
// }

// func GetCancelOrderByDate() {
// 	// Implement the logic to get cancel order by date here
// }
