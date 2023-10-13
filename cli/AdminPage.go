package cli

import (
	"bufio"
	"fmt"
	"klepon/handler"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk sistem Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func AdminPage(username string) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println() // Separator
		fmt.Println("*****	Welcome To Admin Page	*****")
		fmt.Println("1. Order List")
		fmt.Println("2. Payment")
		// fmt.Println("3. Menu")
		fmt.Println("3. Report")
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
				continue
			}

			// Output
			// Retrive Orders Data
			var username, orderDate string
			counter := 1
			for _, o := range orders {
				date := strings.Split(o.OrderDate, " ")

				if username != o.Username || o.OrderDate != orderDate {
					counter = 1

					fmt.Println() // Separator
					fmt.Println("Order ID		: ", o.ID)
					fmt.Println("Username		: ", o.Username)
					fmt.Println("Order Date		: ", date[0])
					fmt.Println("Menu Items : ")
					fmt.Printf("  %v. %v --- x%v\n", counter, o.Menu, o.Qty)
				} else {
					counter += 1
					fmt.Printf("  %v. %v --- x%v\n", counter, o.Menu, o.Qty)
					continue
				}

				orderDate = o.OrderDate
				username = o.Username
			}

			fmt.Println()                                            // Separator
			fmt.Println("-----------------------------------------") // Separator

			OrderListAction()
			clearScreen()

		case "2":
			fmt.Printf("\n----------	Payment Menu - Order List	----------\n")

			orders, err := handler.GetPaymentOrders()
			if err != nil {
				fmt.Printf("\n*** Failed to get order list. Please try again! \n[ERR] ---> %v\n\n", err)
				continue
			}

			// Output
			// Retrive Orders Data
			var username, orderDate string
			counter := 1
			for _, o := range orders {
				if username != o.Username || o.OrderDate != orderDate {
					counter = 1
					fmt.Println() // Separator
					fmt.Println("Order ID		: ", o.ID)
					fmt.Println("Username		: ", o.Username)
					fmt.Println("Order Date		: ", o.OrderDate)
					fmt.Println("Menu Items : ")
					fmt.Printf("  %v. %v -- Rp.%v | x%v\n", counter, o.Menu, o.Price, o.Qty)

				} else {
					counter += 1
					fmt.Printf("  %v. %v -- Rp.%v | x%v\n", counter, o.Menu, o.Price, o.Qty)
					continue
				}

				orderDate = o.OrderDate
				username = o.Username
			}

			fmt.Println()                                            // Separator
			fmt.Println("-----------------------------------------") // Separator

			for {
				// Order Id input
				fmt.Print("Insert ID (type 0 to exit) : ")
				scanner.Scan()
				orderId, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
				if err != nil {
					fmt.Println("Invalid input")
					continue
				}

				if orderId == 0 {
					break
				}

				// Discount input
				fmt.Print("Insert Discount : ")
				scanner.Scan()
				discount, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
				if err != nil {
					fmt.Println("Invalid input")
					continue
				}
				if discount < 0 || discount > 100 {
					fmt.Println("Discount must be greater than 0 and less then 100")
					continue
				}

				type OrderDetail struct {
					Menu  string
					Price int
					Qty   int
					Total int
				}

				var orderDetails []OrderDetail

				// Retrive Order Detail
				var username, orderDate string
				var totalAmount int
				for _, o := range orders {
					var od OrderDetail

					if o.ID == orderId {
						username = o.Username
						orderDate = o.OrderDate

						od.Menu = o.Menu
						od.Price = o.Price
						od.Qty = o.Qty
						od.Total = o.Price * o.Qty

						totalAmount += od.Total

						orderDetails = append(orderDetails, od)
					}
				}

				totalAmount -= totalAmount * discount / 100

				// Output
				fmt.Println() // Separator
				fmt.Println("------	Payment Confirmation	------")
				fmt.Println("Order ID		: ", orderId)
				fmt.Println("Username		: ", username)
				fmt.Println("Order Date		: ", orderDate)
				fmt.Println("Menu Items : ")

				counter = 1
				for _, v := range orderDetails {
					fmt.Printf("  %v. %v -- Rp.%v | x%v ---> Rp.%v\n", counter, v.Menu, v.Price, v.Qty, v.Total)
					counter += 1
				}

				fmt.Printf("\n* Discount : %v%%\n", discount)
				fmt.Printf("* Total Amount : %v\n\n", totalAmount)

				// Confirmation
				fmt.Print("Confirm [Y/N] : ")
				scanner.Scan()
				confirm := strings.TrimSpace(strings.ToLower(scanner.Text()))

				if confirm == "y" {
					msg := handler.PaymentProcess(orderId, discount, totalAmount)
					fmt.Printf("\n*** %v ***\n", msg)
					fmt.Println()
					break
				} else if confirm == "n" {
					fmt.Println("*** Payment Cancel ***")
					break
				} else {
					fmt.Println("Invalid input")
				}
			}
			clearScreen()

		case "3":
			for {
				fmt.Printf("\n----------	Report Menu	----------\n")
				fmt.Println("1. Order Report")
				fmt.Println("2. Menu Report")
				fmt.Println("3. Customer Report")
				fmt.Println("0. Back")

				fmt.Println() // Separator

				fmt.Print("Choice : ")
				scanner.Scan()
				choice = strings.TrimSpace(scanner.Text())

				if choice == "0" {
					break
				}

				switch choice {
				case "1":
					report, status := handler.GeneralReport()
					if !status {
						fmt.Println("Failed to generate report")
						continue
					}

					// Date
					current := time.Now()
					y, m, _ := current.Date()

					// Output
					fmt.Printf("\n---------- General Report ----------\n")
					fmt.Printf("Date			: %v %v\n", m, y)
					fmt.Println("Order Success		:", report.Success)
					fmt.Println("Order Cancel		:", report.Cancel)
					fmt.Printf("Order Revenue		: Rp.%v\n", report.Revenue)
					fmt.Println("---------------------------------------")
					

				case "2":
					report, err := handler.MenuReport()
					if err != nil {
						fmt.Println("Failed to generate report")
						continue
					}

					// Date
					current := time.Now()
					y, m, _ := current.Date()

					// Output
					fmt.Printf("\n---------- Menu Report ----------\n\n")
					fmt.Printf("Date : %v %v\n\n", m, y)
					for _, v := range report {
						fmt.Printf("ID			: %v\n", v.ID)
						fmt.Printf("Menu			: %v\n", v.Menu)
						fmt.Printf("Category		: %v\n", v.Category)
						fmt.Printf("Success			: %vx\n\n", v.OrderSuccess)
					}

					fmt.Println("---------------------------------------")

				case "3":
					report, err := handler.CustomerReport()
					if err != nil {
						fmt.Println("Failed to generate report")
						continue
					}

					// Date
					current := time.Now()
					y, m, _ := current.Date()

					// Output
					fmt.Printf("\n---------- Customer Report ----------\n\n")
					fmt.Printf("Date : %v %v\n\n", m, y)
					for _, v := range report {
						fmt.Printf("Username		: %v\n", v.Username)
						fmt.Printf("Order Success		: %v\n", v.OrderSuccess)
						fmt.Printf("Order Cancel		: %v\n\n", v.OrderCancel)
					}

					fmt.Println("---------------------------------------")

				default:
					fmt.Println("Invalid Input")
				}
			}

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
