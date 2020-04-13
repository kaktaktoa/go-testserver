package app

import (
	"../models"
	util "../utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func responseWithMessage(response map[string] interface{}, w http.ResponseWriter, m string) {
	response = util.Message(false, m)
	w.WriteHeader(http.StatusForbidden)
	w.Header().Add("Content-Type", "application/json")
	util.Respond(w, response)
	return
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			notAuth := []string{"/api/user/new","/api/user/login"}
			requestPath := r.URL.Path
			for _, value := range notAuth {
				if value == requestPath {
					next.ServeHTTP(w, r)
					return
				}
			}
			response := make(map[string]interface{})
			tokenHeader := r.Header.Get("Authorization")
			splitter := strings.Split(tokenHeader, " ")
			if tokenHeader == "" || len(splitter) != 2{
				responseWithMessage(response, w, "Missing auth token")
				return
			}
			tk := &models.Token{}
			token, err := jwt.ParseWithClaims(splitter[1], tk,func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("token_password")), nil
			})
			if !token.Valid && err == nil {
				responseWithMessage(response, w, "Token is not valid")
				return
			}
			ctx := context.WithValue(r.Context(), "user", tk.UserId)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
	})
}