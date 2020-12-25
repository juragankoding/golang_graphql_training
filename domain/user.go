package domain

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Display_name string `json:"display_name"`
}

type UserRepository interface {
	SingleUserFromID(id int64) (*User, error)
	SingleUser(username *string) (*User, error)
	CreateUser(username *string, password *string, displayName *string) (*User, string, error)
	ListUsers() ([]*User, error)
}

type UserUseCase interface {
	Login(username *string, password *string) (*User, string, error)
	Logout() error
	CreateUser(username *string, password *string, displayName *string) (*User, string, error)
	ListUsers() ([]*User, error)
	Users(token string) (*User, error)
	SingleUserFromID(id int64) (*User, error)
}

type ClaimsUser struct {
	jwt.StandardClaims
	UserID int `json:"Username"`
}
