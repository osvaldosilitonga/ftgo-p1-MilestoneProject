package handler

import (
	"context"
	"klepon/config"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (bool, error) {
	// Connect to DB
	db, err := config.ConnDB()
	if err != nil {
		log.Fatal("Failed to connect DB", err)
	}
	defer db.Close()

	var hashedPassword string
	var isAdmin bool

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.QueryRowContext(ctx, "SELECT password, isAdmin FROM users WHERE username = ?", username).Scan(&hashedPassword, &isAdmin)
	if err != nil {
		return isAdmin, err
	}

	// Password compare
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return isAdmin, err // Password not match
	}

	return isAdmin, nil // Authentication Success

}
