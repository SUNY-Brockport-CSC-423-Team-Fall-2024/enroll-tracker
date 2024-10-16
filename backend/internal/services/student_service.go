package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"errors"
)

type StudentService struct {
	repository                repositories.StudentRepository
	userAuthenticationService *UserAuthenticationService
}

func CreateStudentService(repo repositories.StudentRepository, uaService *UserAuthenticationService) *StudentService {
	return &StudentService{repository: repo, userAuthenticationService: uaService}
}

func (s *StudentService) CreateStudent(username string, password string, firstName string, lastName string, phoneNumber string, email string) (models.Student, error) {
	//create student
	var student models.Student

	//Validate Username & Password
	validUsername := utils.ValidUsername(username)
	validPassword := utils.ValidPassword(password)
	if !validUsername || !validPassword {
		return student, errors.New(`Invalid username or password`)
	}

	//Create new user auth...this checks if user auth already exists.
	userAuth, err := s.userAuthenticationService.CreateNewUserAuthentication(username, password)
	if err != nil {
		return student, err
	}

	//Create student
	student, err = s.repository.CreateStudent(firstName, lastName, userAuth.ID, phoneNumber, email)

	return student, err
}

func (s *StudentService) GetStudent(username string) (models.Student, error) {
	var student models.Student

	//Validate username
	if !utils.ValidUsername(username) {
		return student, errors.New("Invalid username")
	}

	//Get student
	student, err := s.repository.GetStudent(username)

	return student, err
}

func (s *StudentService) UpdateStudent(studentUpdates models.StudentUpdate) (models.Student, error) {
	var student models.Student

	//Validate updates
	if !utils.ValidUsername(studentUpdates.Username) {
		return student, errors.New("Invalid username")
	}
	if studentUpdates.FirstName != nil && !utils.ValidName(*studentUpdates.FirstName) {
		return student, errors.New("Invalid first name")
	}
	if studentUpdates.LastName != nil && !utils.ValidName(*studentUpdates.LastName) {
		return student, errors.New("Invalid last name")
	}
	if studentUpdates.PhoneNumber != nil && !utils.ValidPhoneNumber(*studentUpdates.PhoneNumber) {
		return student, errors.New("Invalid phone number")
	}
	if studentUpdates.Email != nil && !utils.ValidEmail(*studentUpdates.Email) {
		return student, errors.New("Invalid email")
	}

	//Make sure there is at least one update. If not return
	student, err := s.repository.UpdateStudent(studentUpdates)

	return student, err
}
