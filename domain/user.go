package domain

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Display_name string `json:"display_name"`
}

type UserRepository interface {
	Login(username *string, password *string) (*User, error)
	Logout() error
	CreateUser(username *string, password *string, displayName *string) (*User, string, error)
	ListUsers() ([]*User, error)
}

type UserUseCase interface {
	Login(username *string, password *string) (*User, error)
	Logout() error
	CreateUser(username *string, password *string, displayName *string) (*User, string, error)
	ListUsers() ([]*User, error)
}
