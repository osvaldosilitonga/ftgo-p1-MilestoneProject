package entity

type GeneralReport struct {
	Success, Cancel, Revenue int
}

type MenuReport struct {
	ID           int
	Menu         string
	Category     string
	OrderSuccess int
}

type CustomerReport struct {
	Username     string
	OrderSuccess int
	OrderCancel  int
}
