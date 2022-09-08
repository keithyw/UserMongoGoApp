package services

import (
	"messageapp/models"
	"messageapp/repositories"
)

type UserServiceImpl struct {
	Ur repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (svc *UserServiceImpl) CreateUser(user models.User) (*models.User, error) {
	newUser, err := svc.Ur.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (svc *UserServiceImpl) DeleteUser(id string) error {
	err := svc.Ur.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *UserServiceImpl) UpdateUser(id string, user models.User) (*models.User, error) {
	updatedUser, err := svc.Ur.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (svc *UserServiceImpl) FindUserById(id string) (*models.User, error) {
	user, err := svc.Ur.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserServiceImpl) FindUserByName(name string) (*models.User, error) {
	user, err := svc.Ur.FindUserByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	users, err := svc.Ur.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}