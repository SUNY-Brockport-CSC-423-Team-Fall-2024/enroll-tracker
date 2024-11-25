package main

import (
	"enroll-tracker/internal/handlers"
	"enroll-tracker/internal/middleware"
	"enroll-tracker/internal/redis"
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

	redis, err := redis.CreateNewRedisClient()
	if err != nil {
		panic(err.Error())
	}
	defer redis.Close()

	//Create repos
	uaRepo := repositories.CreatePostgresUserAuthenticationRepository(db)
	studentRepo := repositories.CreatePostgresStudentRepository(db)
	teacherRepo := repositories.CreatePostgresTeacherRepository(db)
	administratorRepo := repositories.CreatePostgresAdministratorRepository(db)
	userSessionRepo := repositories.CreatePostgresUserSessionRepository(db)
	redisRepo := repositories.CreateRedisRepository(redis)
	courseRepo := repositories.CreateCourseRepository(db)
	majorRepo := repositories.CreateMajorRepository(db)
	courseMajorRepo := repositories.CreateCourseMajorRepository(db)
	enrollmentsRepo := repositories.CreateEnrollmentsRepository(db)

	//Create services
	userAuthService := services.CreateUserAuthenticationService(uaRepo)
	studentService := services.CreateStudentService(studentRepo, userAuthService)
	teacherService := services.CreateTeacherService(teacherRepo, userAuthService)
	administratorService := services.CreateAdministratorService(administratorRepo, userAuthService)
	userSessionService := services.CreateUserSessionService(userSessionRepo)
	redisSession := services.CreateRedisService(redisRepo)
	courseService := services.CreateCourseService(courseRepo)
	majorService := services.CreateMajorService(majorRepo, courseMajorRepo)
	enrollmentsService := services.CreateEnrollmentsService(enrollmentsRepo)

	//Create http multiplexer
	stdMux := http.NewServeMux()
	stdMux.HandleFunc("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Healthy!")
		w.WriteHeader(200)
	})

	//Student routes
	stdMux.HandleFunc("/api/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateStudentHandler(studentService, w, r)
		case http.MethodGet:
			handlers.GetStudentsHandler(studentService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/students/{username}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetStudentHandler(studentService, w, r)
		case http.MethodPut:
			handlers.UpdateStudentHandler(studentService, w, r)
		case http.MethodDelete:
			handlers.DeleteStudentHandler(studentService, userSessionService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/students/{studentID}/courses", handlers.GetStudentsCoursesHandler(enrollmentsService))
	stdMux.HandleFunc("/api/students/{studentID}/majors", handlers.AddStudentToMajorHandler(studentService))

	//Teacher routes
	stdMux.HandleFunc("/api/teachers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateTeacherHandler(teacherService, w, r)
		case http.MethodGet:
			handlers.GetTeachersHandler(teacherService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/teachers/{username}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTeacherHandler(teacherService, w, r)
		case http.MethodPut:
			handlers.UpdateTeacherHandler(teacherService, w, r)
		case http.MethodDelete:
			handlers.DeleteTeacherHandler(teacherService, userSessionService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})

	//Administrator routes
	stdMux.HandleFunc("/api/administrators", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateAdministratorHandler(administratorService, w, r)
		case http.MethodGet:
			handlers.GetAdministratorsHandler(administratorService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/administrators/{username}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAdministratorHandler(administratorService, w, r)
		case http.MethodPut:
			handlers.UpdateAdministratorHandler(administratorService, w, r)
		case http.MethodDelete:
			handlers.DeleteAdministratorHandler(administratorService, userSessionService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})

	//Auth routes
	stdMux.HandleFunc("/auth/login", handlers.LoginHandler(userSessionService, userAuthService))
	stdMux.HandleFunc("POST /auth/token-refresh", handlers.RefreshTokenHandler(userSessionService, redisSession))
	stdMux.HandleFunc("POST /auth/logout", handlers.LogoutHandler(userSessionService, redisSession))
	stdMux.HandleFunc("POST /auth/change-password", handlers.ChangePasswordHandler(userAuthService))

	//Courses routes
	stdMux.HandleFunc("/api/courses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateCourseHandler(courseService, w, r)
		case http.MethodGet:
			handlers.GetCoursesHandler(courseService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/courses/{courseID}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetCourseHandler(courseService, w, r)
		case http.MethodPut:
			handlers.UpdateCourseHandler(courseService, w, r)
		case http.MethodDelete:
			handlers.DeleteCourseHandler(courseService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})

	//Majors routes
	stdMux.HandleFunc("/api/majors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateMajorHandler(majorService, w, r)
		case http.MethodGet:
			handlers.GetMajorsHandler(majorService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/majors/{majorID}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetMajorHandler(majorService, w, r)
		case http.MethodPut:
			handlers.UpdateMajorHandler(majorService, w, r)
		case http.MethodDelete:
			handlers.DeleteMajorHandler(majorService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/majors/{majorID}/courses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.AddCourseToMajorHandler(majorService, w, r)
		case http.MethodGet:
			handlers.GetCoursesAssoicatedWithMajorHandler(majorService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})
	stdMux.HandleFunc("/api/majors/{majorID}/courses/{courseID}", handlers.DeleteCourseFromMajorHandler(majorService))

	//Enrollment routes
	stdMux.HandleFunc("/api/enrollments/{courseID}/students", handlers.GetCoursesStudentsHandler(enrollmentsService))
	stdMux.HandleFunc("/api/enrollments/{courseID}/{studentID}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.EnrollStudentHandler(enrollmentsService, w, r)
		case http.MethodDelete:
			handlers.UnenrollStudentHandler(enrollmentsService, w, r)
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(200)
		default:
			http.Error(w, "Method not allowed", http.StatusBadRequest)
		}
	})

	//Crete auth middleware
	authMiddleware := middleware.AuthMiddleware(redisSession)

	//Create logging middleware
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	loggingMiddleware := middleware.LoggingMiddleware(logger)

	//Create router with logging and auth middlewares
	loggedRouter := loggingMiddleware(authMiddleware(stdMux))

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
