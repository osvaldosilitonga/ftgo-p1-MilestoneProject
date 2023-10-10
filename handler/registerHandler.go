package handler

import (
	"klepon/config"
	"klepon/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Register(user *entity.User) error {
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password, name, address, email) VALUES (?,?,?,?,?)", user.Username, hashedPassword, user.Name, user.Address, user.Email)
	return err
}
