package models

type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
