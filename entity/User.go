package entity

type User struct {
	ID                   int
	Name, Address, Email string
	Username, Password   string
	IsAdmin              bool
}
