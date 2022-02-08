package entities

import (
	"time"
)

type User struct {
	ID         string    `json:"id"`
	first_name string    `json:"first_name"`
	surname    string    `json:"surname"`
	dob        time.Time `json:"dob"`
	password   string    `json:password`
	email      string    `json:email`
	number     string    `json:number`
}
