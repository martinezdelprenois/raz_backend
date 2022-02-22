package repository

import "raz.zaantu.com/m/v0/domain/dto"

type UserRepository struct {
	handler DatabaseHandler
}

func NewUserRepository(handler DatabaseHandler) UserRepository {
	return UserRepository{handler}
}

func (repository UserRepository) SaveUser(user dto.User) error {
	err := repository.handler.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (repository UserRepository) FindAll() ([]*dto.User,error) {
	result, err := repository.handler.FindAllUsers()
	if err != nil {
		return results, err
	}
	return results, nil
}
