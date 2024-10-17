package main

import (
	"enroll-tracker/internal/handlers"
	"enroll-tracker/internal/middleware"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/internal/services"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	db, err := services.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	uaRepo := repositories.CreatePostgresUserAuthenticationRepository(db)
	studentRepo := repositories.CreatePostgresStudentRepository(db)
	userSessionRepo := repositories.CreatePostgresUserSessionRepository(db)
	userAuthService := services.CreateUserAuthenticationService(uaRepo)
	studentService := services.CreateStudentService(studentRepo, userAuthService)
	userSessionService := services.CreateUserSessionService(userSessionRepo)

	stdMux := http.NewServeMux()
	stdMux.HandleFunc("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Healthy!")
		w.WriteHeader(200)
	})

	//Student routes
    stdMux.HandleFunc("POST /api/students", handlers.CreateStudentHandler(studentService))
	stdMux.HandleFunc("GET /api/students/{username}", handlers.GetStudentHandler(studentService))
	stdMux.HandleFunc("PUT /api/students/{username}", handlers.UpdateStudentHandler(studentService))

	//Auth routes
	stdMux.HandleFunc("POST /auth/login", handlers.LoginHandler(userSessionService, userAuthService))
    stdMux.HandleFunc("POST /auth/token-refresh", handlers.RefreshTokenHandler(userSessionService))
    stdMux.HandleFunc("POST /auth/logout", handlers.LogoutHandler(userSessionService))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	loggedRouter := loggingMiddleware(stdMux)

	l, ok := os.LookupEnv("API_CONTAINER_PORT")

	if !ok {
		fmt.Print("No specified port number")
		return
	}

	listeningPortNumber, err := strconv.Atoi(l)

	if err != nil {
		fmt.Print("Invalid port number")
		return
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", listeningPortNumber),
		Handler: loggedRouter,
	}

	fmt.Printf("Listening on %s...\n", server.Addr)
	_ = server.ListenAndServe()

	shutdown := make(chan os.Signal, 1)
	fmt.Println("Press Ctrl+C to stop.")
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	<-shutdown
}
