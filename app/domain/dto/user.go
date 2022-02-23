package dto

import "time"

type User struct {
	FirstName string    `json:"first_name"`
	Surname   string    `json:"surname"`
	DOB       time.Time `json:"dob"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	AreaCode  string    `json:"area_code"`
	Number    string    `json:"number"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

type UserRepository interface {
	SaveUser(user User) error
	FindAll() ([]*User, error)
}
