package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/juragankoding/golang_graphql_training/domain"
)

var issuer = []byte("github/thealmarques")

func GenerateJWT(userID int64, expiredAt int64) string {
	claims := domain.ClaimsUser{
		UserID: int(userID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    string(issuer),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(issuer)

	if err != nil {
		fmt.Print(err.Error())
	}

	return signedToken
}
