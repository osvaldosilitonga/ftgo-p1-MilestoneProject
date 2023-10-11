package cli

import (
	"bufio"
	"fmt"
	"klepon/handler"
	"os"
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

		fmt.Println() // Separator
		fmt.Print("Choice : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("You select order list")
			handler.GetOrders()
		case "0":
			fmt.Printf("\n*****	You have successfully logged out!	*****\n\n")
			return
		default:
			fmt.Println("Invalid input. Please try again!")
		}
	}
}
