package repository

import "../../domain/dto"

type DatabaseHandler interface {
	SaveUser(user dto.User) error
	FindAllUsers() ([]*dto.User, error)
}