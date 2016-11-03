package models

// Name of user
type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
}

// User model
type User struct {
	Name      Name   `json:"name"`
	Age       int    `json:"age"`
	Job       string `json:"job"`
	Seniority string `json:"seniority"`
}
