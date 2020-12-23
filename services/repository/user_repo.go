package repository

import (
	"database/sql"
	"log"

	"github.com/juragankoding/golang_graphql_training/domain"
)

type userRepo struct {
	Conn *sql.DB
}

func GenerateNewUserRepository(conn *sql.DB) domain.UserRepository {
	return &userRepo{
		Conn: conn,
	}
}

func (u *userRepo) Login(username *string, password *string) (*domain.User, error) {
	return nil, nil
}

func (u *userRepo) Logout() error {
	return nil
}

func (u *userRepo) CreateUser(username *string, password *string, displayName *string) (*domain.User, string, error) {

	statement, err := u.Conn.Prepare("INSERT INTO users (username, password, display_name) values (?,?,?)")

	if err != nil {
		return nil, "", err
	}

	if result, err := statement.Exec(username, password, displayName); err != nil {
		return nil, "", err
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return nil, "", err
		} else {
			return &domain.User{
				ID:           int(id),
				Username:     *username,
				Password:     *password,
				Display_name: *displayName,
			}, "", nil
		}
	}

}

func (u *userRepo) ListUsers() ([]*domain.User, error) {
	var users []*domain.User

	query, err := u.Conn.Query("SELECT * FROM users")

	if err != nil {
		log.Printf("cannot get user %s", err.Error())

		return users, err
	}

	for query.Next() {
		var user domain.User

		switch err := query.Scan(user.ID, user.Username, user.Password, user.Display_name); err {
		case sql.ErrNoRows:
			return users, sql.ErrNoRows
		case nil:
			users = append(users, &user)
		}
	}

	return users, nil
}
