package repositories

import "messageapp/models"

//  UserRepository
type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, user models.User) (*models.User, error)

	FindUserById(id string) (*models.User, error)
	FindUserByUserId(userId int64) (*models.User, error)
	FindUserByName(name string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}