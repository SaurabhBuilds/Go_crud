package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

var Users []User
