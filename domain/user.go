type User struct{
	username string `json:"username"`
	password string `json:"password"`
	display_name string `json:"display_name"`
}

type UserRepository interface{
Login(username: *string, password: *string)(User, error)
Logout()(error)
CreateUser(username:*string, password:*string) (User, token, error)
}

type UserUseCase interface{
	Login(username: *string, password: *string)(User, error)
	Logout()(error)
	CreateUser(username:*string, password:*string) (User, token, error)
}