package cli

import (
	"bufio"
	"fmt"
	"klepon/handler"
	"os"
	"strconv"
)

func AdminPage(username string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println() // Separator
		fmt.Println("*****	Welcome To Admin Page	*****")
		fmt.Println("1. Order List")
		fmt.Println("2. Payment")
		fmt.Println("3. Menu")
		fmt.Println("4. User")
		fmt.Println("5. Report")
		fmt.Println("0. Logout")

		// User Input
		fmt.Println() // Separator
		fmt.Print("Choice : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Printf("\n----------	Order List	----------\n")

			orders, err := handler.GetOrders()
			if err != nil {
				fmt.Printf("\n*** Failed to get order list. Please try again! \n[ERR] ---> %v\n\n", err)
			}

			// Output
			// Retrive Orders Data
			var username string
			counter := 1
			for _, o := range orders {
				if username != o.Username {
					fmt.Println() // Separator
					fmt.Println("Order ID		: ", o.ID)
					fmt.Println("Username		: ", o.Username)
					fmt.Println("Order Date		: ", o.OrderDate)
					fmt.Println("Menu Items : ")
					fmt.Printf("  %v. %v --- x%v\n", counter, o.Menu, o.Qty)

					counter = 1
				} else {
					counter += 1
					fmt.Printf("  %v. %v --- x%v\n", counter, o.Menu, o.Qty)
					continue
				}

				username = o.Username
			}

			fmt.Println()                                            // Separator
			fmt.Println("-----------------------------------------") // Separator

			OrderListAction()

		case "0":
			fmt.Printf("\n*****	You have successfully logged out!	*****\n\n")
			return
		default:
			fmt.Println("Invalid input. Please try again!")
		}
	}
}

func OrderListAction() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("\n****\t Actions\t ****\n")
		fmt.Println("1. Process Order")
		fmt.Println("2. Cancel Order")
		fmt.Println("0. Back")
		fmt.Println() // Separator

		// User Input
		fmt.Print("Action : ")
		scanner.Scan()
		action := scanner.Text()

		switch action {
		case "1":
			fmt.Println()
			fmt.Println("Process Order")
			fmt.Print("Insert Order ID : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			msg := handler.ProcessOrder(id)
			fmt.Println(msg)

		case "2":
			fmt.Println()
			fmt.Println("Cancel Order")
			fmt.Print("Insert Order ID : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			msg := handler.CancelOrder(id)
			fmt.Println(msg)

		case "0":
			return

		default:
			fmt.Println("*Invalid input!")
		}
	}
}
