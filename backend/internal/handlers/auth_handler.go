package handlers

import (
	"context"
	"encoding/json"
	"enroll-tracker/internal/middleware"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"enroll-tracker/pkg/utils"
	"net/http"
	"time"
)

func LoginHandler(userSessionService *services.UserSessionService, userAuthService *services.UserAuthenticationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		//Get request body
		var kv map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
			http.Error(w, "Login failed: Invalid username or password", http.StatusBadRequest)
			return
		}

		username, usernameOk := kv["username"]
		password, passwordOk := kv["password"]

		if !usernameOk || !passwordOk {
			http.Error(w, "Login failed: Invalid username or password", http.StatusBadRequest)
			return
		}
		u, uok := username.(string)
		p, pok := password.(string)
		if !uok || !pok {
			http.Error(w, "Login failed: Invalid username or password", http.StatusBadRequest)
			return
		}

		//Get userauth
		userAuth, err := userAuthService.GetUserAuthentication(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Validate username & password
		valid, err := utils.VerifyHashedText(p, userAuth.PasswordHash)
		if err != nil || !valid {
			http.Error(w, "Login failed: Invalid username and password", http.StatusBadRequest)
			return
		}

		//Revoke user user sessions
		_, err = userSessionService.RevokeUserSessionWithUsername(userAuth.Username)
		if err != nil {
			http.Error(w, "Login failed: Invalid username and password", http.StatusUnauthorized)
			return
		}

		//Create token resposne object
		var tokenResponse models.TokenResponse

		//Create user sess
		userSession, accessToken, err := userSessionService.CreateUserSession(userAuth.ID, userAuth.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Update last login
		success, err := userAuthService.UpdateLastLogin(userAuth.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !success {
			http.Error(w, "Error occured while logging user in", http.StatusInternalServerError)
			return
		}

		//Assign values to token response before serialization
		tokenResponse.AccessToken = accessToken
		tokenResponse.RefreshTokenID = userSession.RefreshTokenID
		tokenResponse.RefreshToken = userSession.RefreshToken
		tokenResponse.ExpiresIn = userSession.ExpiresAt

		if err := json.NewEncoder(w).Encode(&tokenResponse); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(200)
	}
}

func RefreshTokenHandler(userSessionService *services.UserSessionService, redisService *services.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		var kv map[string]interface{}

		//Get request body
		if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}

		//Validate refresh token and refresh token id
		rt, rtOk := kv["refresh_token"]
		rtid, rtidOk := kv["refresh_token_id"]
		if !rtOk || !rtidOk {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}
		refreshToken, refreshTokenOk := rt.(string)
		refreshTokenID, refreshTokenIDOk := rtid.(string)
		if !refreshTokenOk || !refreshTokenIDOk {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}

		//Get user session
		userSession, err := userSessionService.GetUserSession(refreshTokenID)
		if err != nil {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}

		//Verify refresh token
		ok, err := utils.VerifyHashedText(refreshToken, userSession.RefreshToken)
		if err != nil || !ok {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}

		//Validate user session
		ok = userSessionService.IsUserSessionValid(userSession)
		if !ok {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}

		//Refresh token is valid..invalidate old refresh token, create new access token, refresh token, refresh token id
		ok, err = userSessionService.RevokeUserSessionWithID(refreshTokenID)
		if err != nil || !ok {
			http.Error(w, "Unable to refresh token", http.StatusBadRequest)
			return
		}

		//Create token resposne object
		var tokenResponse models.TokenResponse

		//Create user sess
		newUserSession, newAccessToken, err := userSessionService.CreateUserSession(userSession.UserID, userSession.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Assign values to token response before serialization
		tokenResponse.AccessToken = newAccessToken
		tokenResponse.RefreshTokenID = newUserSession.RefreshTokenID
		tokenResponse.RefreshToken = newUserSession.RefreshToken
		tokenResponse.ExpiresIn = newUserSession.ExpiresAt

		if err := json.NewEncoder(w).Encode(&tokenResponse); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(200)
	}
}

func LogoutHandler(userSessionService *services.UserSessionService, redisService *services.RedisService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		var kv map[string]interface{}
		ctx := context.Background()
		bearerToken := r.Header.Get("Authorization")

		//Make sure authorization header was included
		if bearerToken == "" {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}

		//Get request body
		if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}

		//Validate refresh token and refresh token id
		rt, rtOk := kv["refresh_token"]
		rtid, rtidOk := kv["refresh_token_id"]
		if !rtOk || !rtidOk {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}
		refreshToken, refreshTokenOk := rt.(string)
		refreshTokenID, refreshTokenIDOk := rtid.(string)
		if !refreshTokenOk || !refreshTokenIDOk {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}

		//Get context info
		accessToken, accessTokenOK := (r.Context().Value(middleware.AccessTokenKey)).(string)
		claims, claimsOK := (r.Context().Value(middleware.ClaimsKey)).(*models.CustomClaims)
		if !accessTokenOK || !claimsOK || claims == nil {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}

		//Get user session
		userSession, err := userSessionService.GetUserSession(refreshTokenID)
		if err != nil {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}

		//Verify refresh token
		ok, err := utils.VerifyHashedText(refreshToken, userSession.RefreshToken)
		if err != nil || !ok {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}

		//Blacklist old access token in redis cache until it expires
		ttl := claims.ExpiresAt.Time.Sub(time.Now().UTC()).Seconds()
		if ttl > 0 {
			err := redisService.Set(ctx, accessToken, 1, time.Duration(ttl)*time.Second)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		//Revoke user session
		ok, err = userSessionService.RevokeUserSessionWithID(refreshTokenID)
		if err != nil || !ok {
			http.Error(w, "Unable to successfully logout", http.StatusBadRequest)
			return
		}
		w.WriteHeader(200)
	}
}

func ChangePasswordHandler(userAuthService *services.UserAuthenticationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		//Get request body
		var kv map[string]interface{}

		//Get request body
		if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
			http.Error(w, "Error ocurred while changing users password", http.StatusBadRequest)
			return
		}

		currentPassword, currentpasswordOk := kv["current_password"]
		newPassword, newPasswordOk := kv["new_password"]
		if !newPasswordOk || !currentpasswordOk {
			http.Error(w, "Current or new password not supplied", http.StatusBadRequest)
			return
		}
		cP, cpOk := currentPassword.(string)
		nP, npOk := newPassword.(string)
		if !npOk || !cpOk {
			http.Error(w, "Current or new password not valid format", http.StatusBadRequest)
			return
		}

		//Get username from context
		claims, claimsOK := (r.Context().Value(middleware.ClaimsKey)).(*models.CustomClaims)
		if !claimsOK || claims == nil {
			http.Error(w, "Error ocurred while changing users password", http.StatusBadRequest)
			return
		}
		username := claims.Subject

		//Change users password
		userAuth, err := userAuthService.ChangePassword(username, nP, cP)
		if err != nil || userAuth == nil {
			http.Error(w, "Error ocurred while changing users password", http.StatusBadRequest)
			return
		}
		w.WriteHeader(200)
	}
}
