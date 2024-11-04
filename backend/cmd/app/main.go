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
	stdMux.HandleFunc("POST /api/students", handlers.CreateStudentHandler(studentService))
	stdMux.HandleFunc("GET /api/students", handlers.GetStudentsHandler(studentService))
	stdMux.HandleFunc("GET /api/students/{username}", handlers.GetStudentHandler(studentService))
	stdMux.HandleFunc("PUT /api/students/{username}", handlers.UpdateStudentHandler(studentService))
	stdMux.HandleFunc("DELETE /api/students/{username}", handlers.DeleteStudentHandler(studentService, userSessionService))
	stdMux.HandleFunc("GET /api/students/{studentID}/courses", handlers.GetStudentsCoursesHandler(enrollmentsService))
	stdMux.HandleFunc("POST /api/students/{studentID}/majors", handlers.AddStudentToMajorHandler(studentService))

	//Teacher routes
	stdMux.HandleFunc("POST /api/teachers", handlers.CreateTeacherHandler(teacherService))
	stdMux.HandleFunc("GET /api/teachers", handlers.GetTeachersHandler(teacherService))
	stdMux.HandleFunc("GET /api/teachers/{username}", handlers.GetTeacherHandler(teacherService))
	stdMux.HandleFunc("PUT /api/teachers/{username}", handlers.UpdateTeacherHandler(teacherService))
	stdMux.HandleFunc("DELETE /api/teachers/{username}", handlers.DeleteTeacherHandler(teacherService, userSessionService))

	//Administrator routes
	stdMux.HandleFunc("POST /api/administrators", handlers.CreateAdministratorHandler(administratorService))
	stdMux.HandleFunc("GET /api/administrators", handlers.GetAdministratorsHandler(administratorService))
	stdMux.HandleFunc("GET /api/administrators/{username}", handlers.GetAdministratorHandler(administratorService))
	stdMux.HandleFunc("PUT /api/administrators/{username}", handlers.UpdateAdministratorHandler(administratorService))
	stdMux.HandleFunc("DELETE /api/administrators/{username}", handlers.DeleteAdministratorHandler(administratorService, userSessionService))

	//Auth routes
	stdMux.HandleFunc("/auth/login", handlers.LoginHandler(userSessionService, userAuthService))
	stdMux.HandleFunc("POST /auth/token-refresh", handlers.RefreshTokenHandler(userSessionService, redisSession))
	stdMux.HandleFunc("POST /auth/logout", handlers.LogoutHandler(userSessionService, redisSession))
	stdMux.HandleFunc("POST /auth/change-password", handlers.ChangePasswordHandler(userAuthService))

	//Courses routes
	stdMux.HandleFunc("POST /api/courses", handlers.CreateCourseHandler(courseService))
	stdMux.HandleFunc("GET /api/courses", handlers.GetCoursesHandler(courseService))
	stdMux.HandleFunc("GET /api/courses/{courseID}", handlers.GetCourseHandler(courseService))
	stdMux.HandleFunc("PUT /api/courses/{courseID}", handlers.UpdateCourseHandler(courseService))
	stdMux.HandleFunc("DELETE /api/courses/{courseID}", handlers.DeleteCourseHandler(courseService))

	//Majors routes
	stdMux.HandleFunc("POST /api/majors", handlers.CreateMajorHandler(majorService))
	stdMux.HandleFunc("GET /api/majors", handlers.GetMajorsHandler(majorService))
	stdMux.HandleFunc("GET /api/majors/{majorID}", handlers.GetMajorHandler(majorService))
	stdMux.HandleFunc("PUT /api/majors/{majorID}", handlers.UpdateMajorHandler(majorService))
	stdMux.HandleFunc("DELETE /api/majors/{majorID}", handlers.DeleteMajorHandler(majorService))
	stdMux.HandleFunc("POST /api/majors/{majorID}/courses", handlers.AddCourseToMajorHandler(majorService))
	stdMux.HandleFunc("GET /api/majors/{majorID}/courses", handlers.GetCoursesAssoicatedWithMajorHandler(majorService))
	stdMux.HandleFunc("DELETE /api/majors/{majorID}/courses/{courseID}", handlers.DeleteCourseFromMajorHandler(majorService))

	//Enrollment routes
	stdMux.HandleFunc("GET /api/enrollments/{courseID}/students", handlers.GetCoursesStudentsHandler(enrollmentsService))
	stdMux.HandleFunc("POST /api/enrollments/{courseID}/{studentID}", handlers.EnrollStudentHandler(enrollmentsService))
	stdMux.HandleFunc("DELETE /api/enrollments/{courseID}/{studentID}", handlers.UnenrollStudentHandler(enrollmentsService))

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
