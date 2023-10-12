package entity


type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
    ID         int
    UserID     int
    Customer   string
    MenuID     int
	Menu string
    Quantity   int
    OrderTime  string
    TotalAmount float64
    Total      float64
    Status     string
    Price      float64 // Ubah tipe data Price ke float64 jika itu tipe data yang benar di tabel Anda
    // Field lain yang sesuai dengan struktur database Anda
}



type User struct {
	ID	int
	Name     string
	Address  string
	Email    string
	Username string
	Password string
	IsAdmin  bool
}



