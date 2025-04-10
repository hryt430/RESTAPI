package userService

import (
	"github.com/hryt430/RESTAPI/api/internal/domain/entity"
)

type UserDomainService struct {
	UserServiceRepository UserServiceRepository
}

func NewUserDomainService(UserServiceRepository UserServiceRepository) *UserDomainService {
	return &UserDomainService{UserServiceRepository: UserServiceRepository}
}

func (uds *UserDomainService) FindUser() ([]*entity.User, error) {
	users, err := uds.UserServiceRepository.FindUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uds *UserDomainService) FindUserById(id int) (*entity.User, error) {
	user, err := uds.UserServiceRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uds *UserDomainService) CreateUser(params *entity.User) (*entity.User, error) {
	id, err := uds.UserServiceRepository.Save(params)
	if err != nil {
		return nil, err
	}
	user, err := uds.UserServiceRepository.FindUserById(id)
	return user, err
}

func (uds *UserDomainService) EditUser(id int, param *entity.User) (*entity.User, error) {
	id, err := uds.UserServiceRepository.Edit(id, param)
	if err != nil {
		return nil, err
	}
	user, err := uds.UserServiceRepository.FindUserById(id)
	return user, err
}

func (uds *UserDomainService) DeleteUser(id int) error {
	err := uds.UserServiceRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
