package repository

import "raz.zaantu.com/m/v0/domain/dto"

type DatabaseHandler interface {
	SaveUser(user dto.User) (*dto.User, error)
	FindAllUsers() ([]*dto.User, error)
}
