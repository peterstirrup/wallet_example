package entities

type Order struct {
	OrderID string
	UserID  string
	Amount  float64
	Success bool
}
