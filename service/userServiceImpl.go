package service

import (
	"sadiq/Go_Rest_API/models"
	"sadiq/Go_Rest_API/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(ID string) error
	GetUserByID(ID string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(ID string) error {
	return s.userRepo.DeleteUser(ID)
}

func (s *userService) GetUserByID(ID string) (*models.User, error) {
	return s.userRepo.GetUserByID(ID)
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.userRepo.GetAllUsers()
}
