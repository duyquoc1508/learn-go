package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	config "go-demo/configs"
	handler "go-demo/handlers"
	util "go-demo/utils"
)

var jwtKey []byte = []byte(config.JWT_SECRET_KEY)

func Auth(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			util.ResponseErr(w, http.StatusBadRequest, "Access token is required")
			return
		}
		tokenSplit := strings.Split(accessToken, " ")
		if len(tokenSplit) != 2 {
			util.ResponseErr(w, http.StatusBadRequest, "Invalid access token")
			return
		}
		tokenWithoutBearer := tokenSplit[1]
		claims := &handler.Claims{}

		token, err := jwt.ParseWithClaims(tokenWithoutBearer, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			util.ResponseErr(w, http.StatusBadRequest, err.Error())
			return
		}
		if !token.Valid {
			util.ResponseErr(w, http.StatusUnauthorized, "Unauthorized")
		}
		ctx := context.WithValue(r.Context(), handler.Claims{}, token.Claims) // attach data to request
		originalHandler.ServeHTTP(w, r.WithContext(ctx))
	})
}
