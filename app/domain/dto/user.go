package dto

import "time"

type User struct {
	ID         primitive.ObjectID `json:"_id"`
	first_name string             `json:"first_name"`
	surname    string             `json:"surname"`
	dob        time.Time          `json:"dob"`
	password   string             `json:"password"`
	email      string             `json:"email"`
	area_code  string             `json:"area_code"`
	number     string             `json:"number"`
	created    time.Date          `json:created`
	updated    time.Date          `json:updated`
}

type UserRepository interface {
	SaveUser(user User) error
	FindAll() ([]*User, error)
}
