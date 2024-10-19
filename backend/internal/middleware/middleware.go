package middleware

import (
	"context"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"enroll-tracker/pkg/utils"
	"log/slog"
	"net/http"
	"regexp"
	"runtime/debug"
	"time"
)

var (
	protectedRoutes = map[*regexp.Regexp]models.RouteAuth{
		//regexp.MustCompile(`^/api/students$`): {
		//	MethodRoles: []models.MethodRoles{
		//		{
		//			Roles:       []string{utils.Roles.ADMIN},
		//			HTTPMethods: []string{http.MethodPost},
		//		},
		//	},
		//},
		regexp.MustCompile(`^/api/students/[a-zA-Z0-9]{8,50}(\?.*)?$`): { //matches /api/students/{username}?queryparams...
			MethodRoles: []models.MethodRoles{
				{
					Roles:       []string{utils.Roles.ADMIN, utils.Roles.TEACHER},
					HTTPMethods: []string{http.MethodPut},
				},
				{
					Roles:       []string{utils.Roles.STUDENT},
					HTTPMethods: []string{http.MethodGet},
				},
			},
		},
		regexp.MustCompile(`^/auth/token-refresh(\?.*)?$`): {
			MethodRoles: []models.MethodRoles{
				{
					Roles:       []string{utils.Roles.ADMIN, utils.Roles.TEACHER, utils.Roles.STUDENT},
					HTTPMethods: []string{http.MethodPost},
				},
			},
		},
		regexp.MustCompile(`^/auth/logout(\?.*)?$`): {
			MethodRoles: []models.MethodRoles{
				{
					Roles:       []string{utils.Roles.ADMIN, utils.Roles.TEACHER, utils.Roles.STUDENT},
					HTTPMethods: []string{http.MethodPost},
				},
			},
		},
	}
)

// LoggingMiddleware logs the incoming HTTP request, it's duration, and errors if there is a panic()
func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		//Middleware logic here
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error("error occured",
						"err", err,
						"trace", string(debug.Stack()),
					)
				}
			}()
			//Start timer to time it takes to complete next.ServeHTTP
			start := time.Now()
			//Wrapped response writer to capture response status
			wrapped := models.WrapResponseWriter(w)
			//Call the next handler
			next.ServeHTTP(wrapped, r)
			//Log after request has been served
			logger.Info("request received",
				"status", wrapped.Status(),
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
				"src_ip", r.RemoteAddr,
			)
		}

		return http.HandlerFunc(fn)
	}
}

func AuthMiddleware(redisService *services.RedisService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			//1. Identify path, method. And check if it's a protected route
			path := r.URL.EscapedPath()
			method := r.Method
			ok := false
			var routeAuth models.RouteAuth
			for regex, ra := range protectedRoutes {
				if regex.MatchString(path) {
					ok = true
					routeAuth = ra
					break
				}
			}
			if !ok {
				//Not a protected route
				next.ServeHTTP(w, r)
				return
			}

			//2. It's a protected route. Get bearer token
			bearerToken := r.Header.Get("Authorization")
			if bearerToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			accessToken := utils.ExtractAccessTokenFromAuthHeader(bearerToken)

			//3. Verify JWT access token (signature and expires at time)
			claims, err := utils.VerifyJWT(accessToken)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			//If the expires at time is equal to now() or is in the past. Request is unauthorized
			if time.Now().Compare(claims.ExpiresAt.Time) >= 0 {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			hasRole := false
			hasAppropriateHTTPMethod := false

			//4. Check roles and corresponding methods for route
			for _, methodRole := range routeAuth.MethodRoles {
				//Make sure the route with this methodrule allows the role
				if utils.Contains(claims.Role, &methodRole.Roles) {
					hasRole = true
				}

				if utils.Contains(method, &methodRole.HTTPMethods) {
					hasAppropriateHTTPMethod = true
				}

				//Break out of loop if we got sufficient roles
				if hasRole && hasAppropriateHTTPMethod {
					break
				}
			}

			if !(hasRole && hasAppropriateHTTPMethod) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
			//4. Check blacklist for access token
			ttl, err := redisService.Get(context.Background(), accessToken)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusInternalServerError)
				return
			}

			switch v := ttl.(type) {
			case string:
				//1 represents true. Redis doesn't have a boolean value
				if v == "1" {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			case nil:
				//Executes when the access token isn't on the blacklist
				break
			}

			//5. If auth, call next handler. If not return err
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
