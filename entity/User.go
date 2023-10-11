package entity


type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
    ID        int
    Customer  string
    MenuID    int // Kolom menu_id dalam database
    Quantity  int // Kolom quantity dalam database
    OrderTime string // Kolom order_time dalam database
    Total     float64
    Status    string
    // Field lain yang sesuai dengan struktur database Anda
}

type User struct {
	Name     string
	Address  string
	Email    string
	Username string
	Password string
	IsAdmin  bool
}



