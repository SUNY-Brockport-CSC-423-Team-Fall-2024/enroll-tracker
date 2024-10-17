package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/services"
	"enroll-tracker/pkg/utils"
	"net/http"
	"strings"
)

func LoginHandler(userSessionService *services.UserSessionService, userAuthService *services.UserAuthenticationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Login failed: Invalid user ID or password", http.StatusBadRequest)
			return
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "" || password == "" {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
				_, err2 := userSessionService.GetUserSession(token)
				if err2 == nil {
					w.WriteHeader(http.StatusOK)
					return
				}
				http.Error(w, err2.Error(), http.StatusBadRequest)
				return
			}

			http.Error(w, "Login failed: Invalid user ID or password", http.StatusBadRequest)
			return
		}

		//Get userauth
		userAuth, err := userAuthService.GetUserAuthentication(username)
		if err != nil {
			http.Error(w, "Login failed: Invalid user ID or password", http.StatusBadRequest)
			return
		}
		//Validate username & password
		valid, err := utils.VerifyPassword(password, userAuth.PasswordHash)
		if err != nil || !valid {
			http.Error(w, "Login failed: Invalid user ID or password", http.StatusBadRequest)
			return
		}

		//Create user sess
		_, token, err := userSessionService.CreateUserSession(userAuth.ID, userAuth.Username)
		if err != nil {
			http.Error(w, "Login failed: Invalid user ID or password", http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(&token); err != nil {
			http.Error(w, "Login failed: Invalid user ID or password", http.StatusBadRequest)
			return
		}
		w.WriteHeader(200)
	}
}

func RefreshTokenHandler(userSessionService *services.UserSessionService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
         
    }
}

func LogoutHandler(userSessionService *services.UserSessionService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
         
    }
}
