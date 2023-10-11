package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"klepon/cli"
	"klepon/config"
	"klepon/entity"
	"klepon/handler"
)

func main() {
	db, err := config.ConnDB()
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	// // Main For
	for {
		fmt.Println()
		fmt.Println("*****	Welcome to Klepon Coffe Shop	*****")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")

		fmt.Println() // Separator
		fmt.Print("Select Menu : ")
		scanner.Scan()
		choice := scanner.Text()
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			Login()
		case "2":
			Register()
		case "0":
			fmt.Printf("\n***	Thank you and Goodbye	***\n\n")
			return
		}

	}
}

func Login() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n*****	Login Page	*****\n")

	fmt.Print("Username : ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Password : ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	isAdmin, err := handler.Login(username, password)
	if err != nil {
		fmt.Printf("\n*** Failed to login. Please try again. \n[Error] : %v\n", err)
		return
	}

	// Check if user is admin or not
	if isAdmin {
		cli.AdminPage(username)
	} else {
		cli.UserPage(username)
	}
}

func Register() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n*****	Register Page	*****\n")

	fmt.Print("Name : ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	fmt.Print("Address : ")
	scanner.Scan()
	address := strings.TrimSpace(scanner.Text())

	fmt.Print("Email : ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	fmt.Print("Username : ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Password : ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	// Object User
	user := entity.User{
		Name:     name,
		Address:  address,
		Email:    email,
		Username: username,
		Password: password,
	}

	err := handler.Register(&user)
	if err != nil {
		fmt.Printf("Failed register new user. \n[ERROR] : %v", err)
	} else {
		fmt.Printf("\n*** Register Success ***\n\n")
	}

}
