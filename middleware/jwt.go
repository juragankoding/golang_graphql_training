package middleware

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/juragankoding/golang_graphql_training/domain"
	"github.com/juragankoding/golang_graphql_training/utils"
)

type UserAUTH struct {
	UserID    int64
	Roles     []string
	IpAddress string
	Token     string
}

var userCtxKey = &contextKey{name: "user"}

type contextKey struct {
	name string
}

func JwtMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			token := TokenFromHTTPRequestgo(r)
			user := UserIDFromHTTPRequest(token)
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)

			if user >= 0 {
				userAuth := UserAUTH{
					UserID:    user,
					IpAddress: ip,
				}

				ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(rw, r)
		})
	}
}

func TokenFromHTTPRequestgo(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")

	var tokenString string

	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}

	return tokenString
}

func UserIDFromHTTPRequest(tokenString string) int64 {
	token, err := utils.DecodeJwt(tokenString)

	if err != nil {
		return -1
	}

	if claims, ok := token.Claims.(*domain.ClaimsUser); ok && token.Valid {
		if claims == nil {
			return -1
		}

		return int64(claims.UserID)
	}

	return -1
}

func GetUserFromContext(ctx context.Context) *UserAUTH {
	raw := ctx.Value(userCtxKey)

	if raw == nil {
		return nil
	}

	return raw.(*UserAUTH)
}
