package handler

import(
	"fmt"
	"klepon/config"
)

func AuthenticateUser(username, password string) (int, error) {
    // Buat koneksi ke database
    db, err := config.ConnDB()
    if err != nil {
        return 0, err // Mengembalikan kesalahan sebagai nilai balik
    }
    defer db.Close()

    // Query ke database untuk mencari pengguna berdasarkan username
    var storedPassword string
    var userID int
    err = db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&userID, &storedPassword)
    if err != nil {
        return 0, err // Username tidak ditemukan
    }

    // Bandingkan password yang dimasukkan dengan password yang ada di database
    if storedPassword != password {
        return 0, fmt.Errorf("Authentication failed")
    }

    return userID, nil
}