package usecases

import ("log"
	"../domain/dto"
)

type UserInteractor struct {
	UserRepository dto.UserRepository
}

func NewUserInteractor (repository dto.UserRepository) UserInteractor {
	return UserInteractor{repository}
}

func (interactor *UserInteractor) CreateUser(user dto.User) error {
	err := interactor.UserRepository.SaveUser(user)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (interactor *UserInteractor) FindAll() ([]*dto.User, error) {
	results, err := interactor.UserRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return results, nil
}
