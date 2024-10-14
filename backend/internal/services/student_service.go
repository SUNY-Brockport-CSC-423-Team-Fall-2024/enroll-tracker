package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
)

type StudentService struct {
	repository                repositories.StudentRepository
	userAuthenticationService *UserAuthenticationService
}

func CreateNewStudentService(repo repositories.StudentRepository, uaService *UserAuthenticationService) *StudentService {
	return &StudentService{repository: repo, userAuthenticationService: uaService}
}

func (s *StudentService) CreateNewStudent(username string, password string, firstName string, lastName string, phoneNumber string, email string) (models.Student, error) {
	//create student
	var student models.Student

	//Create new user auth...this checks if user auth already exists.
	userAuth, err := s.userAuthenticationService.CreateNewUserAuthentication(username, password)
	if err != nil {
		return student, err
	}

	//Create student
	student, err = s.repository.CreateStudent(firstName, lastName, userAuth.ID, phoneNumber, email)

	return student, err
}
