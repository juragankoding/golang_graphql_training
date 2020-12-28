package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/utils"
)

type userUsecase struct {
	UserRepo domain.UserRepository
}

func NewGenerateUserUseCase(UserRepo *domain.UserRepository) domain.UserUseCase {
	return &userUsecase{
		UserRepo: *UserRepo,
	}
}

func (u *userUsecase) Login(username *string, password *string) (*domain.User, string, error) {
	user, err := u.UserRepo.SingleUser(username)

	if err != nil {
		return nil, "", err
	}

	if !utils.ComparePassword(*password, user.Password) {
		return nil, "", errors.New("password tidak sama")
	}

	expiredAt := int(time.Now().Add(time.Hour * 1).Unix())

	token := utils.GenerateJWT(int64(user.ID), int64(expiredAt))

	fmt.Println(token)

	return user, token, nil
}

func (u *userUsecase) Logout() error {
	return nil
}

func (u *userUsecase) CreateUser(username *string, password *string, diplayName *string) (*domain.User, string, error) {
	return u.UserRepo.Insert(&domain.User{
		Username:     *username,
		Password:     *password,
		Display_name: *diplayName,
	})
}

func (u *userUsecase) ListUsers() ([]*domain.User, error) {
	return u.UserRepo.ListUsers()
}

func (u *userUsecase) Users(token string) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) SingleUserFromID(id int64) (*domain.User, error) {
	return u.UserRepo.SingleUserFromID(id)
}
