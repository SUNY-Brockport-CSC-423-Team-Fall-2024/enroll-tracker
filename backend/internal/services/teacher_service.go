package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"errors"
)

type TeacherService struct {
	repository                repositories.TeacherRepository
	userAuthenticationService *UserAuthenticationService
}

func CreateTeacherService(repo repositories.TeacherRepository, uaService *UserAuthenticationService) *TeacherService {
	return &TeacherService{repository: repo, userAuthenticationService: uaService}
}

func (s *TeacherService) CreateTeacher(username string, password string, firstName string, lastName string, phoneNumber string, email string, office string) (models.Teacher, error) {
	//create teacher
	var teacher models.Teacher

	//Validate Username & Password
	validUsername := utils.ValidUsername(username)
	validPassword := utils.ValidPassword(password)
	if !validUsername || !validPassword {
		return teacher, errors.New(`Invalid username or password`)
	}

	//Create new user auth...this checks if user auth already exists.
	userAuth, err := s.userAuthenticationService.CreateNewUserAuthentication(username, password)
	if err != nil {
		return teacher, err
	}

	//Create teacher
	teacher, err = s.repository.CreateTeacher(firstName, lastName, userAuth.ID, phoneNumber, email, office)

	return teacher, err
}

func (s *TeacherService) GetTeachers(queryParams models.TeacherQueryParams) ([]models.Teacher, error) {
	//Validate query params
	if queryParams.Limit != nil && (*queryParams.Limit < 10 || *queryParams.Limit > 100) {
		return nil, errors.New("Limit is not within range of 10-100")
	}

	//Get teachers
	teachers, err := s.repository.GetTeachers(queryParams)
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (s *TeacherService) GetTeacher(username string) (models.Teacher, error) {
	var teacher models.Teacher

	//Validate username
	if !utils.ValidUsername(username) {
		return teacher, errors.New("Invalid username")
	}

	//Get teacher
	teacher, err := s.repository.GetTeacher(username)

	return teacher, err
}

func (s *TeacherService) UpdateTeacher(username string, teacherUpdates models.TeacherUpdate) (models.Teacher, error) {
	var teacher models.Teacher

	//Validate updates
	if !utils.ValidUsername(username) {
		return teacher, errors.New("Invalid username")
	}
	if teacherUpdates.FirstName != nil && !utils.ValidName(*teacherUpdates.FirstName) {
		return teacher, errors.New("Invalid first name")
	}
	if teacherUpdates.LastName != nil && !utils.ValidName(*teacherUpdates.LastName) {
		return teacher, errors.New("Invalid last name")
	}
	if teacherUpdates.PhoneNumber != nil && !utils.ValidPhoneNumber(*teacherUpdates.PhoneNumber) {
		return teacher, errors.New("Invalid phone number")
	}
	if teacherUpdates.Email != nil && !utils.ValidEmail(*teacherUpdates.Email) {
		return teacher, errors.New("Invalid email")
	}

	//Make sure there is at least one update. If not return
	teacher, err := s.repository.UpdateTeacher(username, teacherUpdates)

	return teacher, err
}

func (s *TeacherService) DeleteTeacher(username string) (bool, error) {
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
