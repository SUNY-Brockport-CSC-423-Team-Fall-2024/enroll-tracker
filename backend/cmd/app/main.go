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
	userAuthService := services.CreateUserAuthenticationService(uaRepo)
	studentService := services.CreateNewStudentService(studentRepo, userAuthService)

	stdMux := http.NewServeMux()
	stdMux.HandleFunc("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Healthy!")
		w.WriteHeader(200)
	})
	stdMux.HandleFunc("/api/students/create-student", handlers.CreateStudentHandler(studentService))
	stdMux.HandleFunc("/api/students/get-student", handlers.GetStudentHandler(studentService))
	stdMux.HandleFunc("/api/students/update-student", handlers.UpdateStudentHandler(studentService))

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
