package usecase

import (
	"github.com/juragankoding/golang_graphql_training/domain"
)

type userUsecase struct {
	UserRepo domain.UserRepository
}

func NewGenerateUserUseCase(UserRepo *domain.UserRepository) domain.UserUseCase {
	return &userUsecase{
		UserRepo: *UserRepo,
	}
}

func (u *userUsecase) Login(username *string, password *string) (*domain.User, error) {
	return u.UserRepo.Login(username, password)
}

func (u *userUsecase) Logout() error {
	return u.UserRepo.Logout()
}

func (u *userUsecase) CreateUser(username *string, password *string, diplayName *string) (*domain.User, string, error) {
	return u.UserRepo.CreateUser(username, password, diplayName)
}

func (u *userUsecase) ListUsers() ([]*domain.User, error) {
	return u.UserRepo.ListUsers()
}
