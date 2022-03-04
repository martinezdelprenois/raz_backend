package usecases

import (
	"log"
	"time"

	"raz.zaantu.com/m/v0/domain/dto"
	"raz.zaantu.com/m/v0/utils"
)

type UserInteractor struct {
	UserRepository dto.UserRepository
}

func NewUserInteractor(repository dto.UserRepository) UserInteractor {
	return UserInteractor{repository}
}

func hashPassword(user *dto.User) {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
}

func updateTime(user *dto.User) {
	user.Created = time.Now()
	user.Updated = time.Now()
}

func editUpdatedTime(user *dto.User) {
	user.Updated = time.Now()
}

func (interactor *UserInteractor) CreateUser(user dto.User) (*dto.User, error) {
	hashPassword(&user)
	updateTime(&user)

	result, err := interactor.UserRepository.SaveUser(user)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (interactor *UserInteractor) FindAll() ([]*dto.User, error) {
	results, err := interactor.UserRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return results, nil
}
