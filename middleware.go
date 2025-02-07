package djan_go

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qor/roles"
	"log"
	"net/http"
	"strings"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var hmacSampleSecret = []byte("my_secret_key")

func (s *HttpDataModel[T]) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(tokenHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			if _, ok := claims["role"]; !ok {
				http.Error(w, "role not found", http.StatusUnauthorized)
				return
			}
			if RoleChecker(s.Permissions, claims["role"].(string), r) {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "role not authorized", http.StatusUnauthorized)
				return
			}
			log.Println("Authenticated user claims:", claims)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Token Claims", http.StatusUnauthorized)
		}
	})
}

func RoleChecker(perm *roles.Permission, role string, r *http.Request) bool {
	if r.Method == http.MethodPost {
		if perm.HasPermission(roles.Create, role) {
			return true
		}
	}

	if r.Method == http.MethodGet {
		if perm.HasPermission(roles.Read, role) {
			return true
		}
	}

	if r.Method == http.MethodPut {
		if perm.HasPermission(roles.Update, role) {
			return true
		}
	}

	if r.Method == http.MethodDelete {
		if perm.HasPermission(roles.Delete, role) {
			return true
		}
	}
	return false
}
