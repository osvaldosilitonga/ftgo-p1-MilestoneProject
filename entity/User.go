package entity


type MenuItem struct {
	ID    int
	Name  string
	Price float64
}

type Order struct {
	ID       int
	Customer string
	Items    []MenuItem
	Total    float64
	Status   string // Pending, Processed, Completed, Canceled, etc.
}

type User struct {
	Name     string
	Address  string
	Email    string
	Username string
	Password string
	IsAdmin  bool
}
