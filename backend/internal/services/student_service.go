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

func (s *StudentService) GetStudents(queryParams models.StudentQueryParams) ([]models.Student, error) {
	//Validate query params
	if queryParams.Limit != nil && (*queryParams.Limit < 10 || *queryParams.Limit > 100) {
		return nil, errors.New("Limit is not within range of 10-100")
	}

	//Get students
	students, err := s.repository.GetStudents(queryParams)
	if err != nil {
		return nil, err
	}

	return students, nil
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

func (s *StudentService) UpdateStudent(username string, studentUpdates models.StudentUpdate) (models.Student, error) {
	var student models.Student

	//Validate updates
	if !utils.ValidUsername(username) {
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
	student, err := s.repository.UpdateStudent(username, studentUpdates)

	return student, err
}

func (s *StudentService) DeleteStudent(username string) (bool, error) {
	//Make sure username is valid
	if !utils.ValidUsername(username) {
		return false, errors.New("Username isn't valid")
	}

	success, err := s.userAuthenticationService.DeleteUserAuthentication(username)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *StudentService) AddStudentToMajor(majorID int, studentID int) (bool, error) {
	success, err := s.repository.AddStudentToMajor(majorID, studentID)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}
	return true, nil
}
