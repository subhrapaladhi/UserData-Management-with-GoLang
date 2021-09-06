package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func Validate(next http.Handler, allowedRoles []string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(viper.GetString("JWTSECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				role, found := claims["role"], false
				for _, v := range allowedRoles {
					if v == role {
						found = true
					}
				}
				if found {
					ctx := context.WithValue(r.Context(), "jwtData", claims)
					next.ServeHTTP(rw, r.WithContext(ctx))
				} else {
					rw.WriteHeader(http.StatusUnauthorized)
					rw.Write([]byte("Unauthorized"))
				}
			} else {
				fmt.Println(err)
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte("Unauthorized"))
			}
		}
	})
}
