package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"errors"
)

type AdministratorService struct {
	repository                repositories.AdministratorRepository
	userAuthenticationService *UserAuthenticationService
}

func CreateAdministratorService(repo repositories.AdministratorRepository, uaService *UserAuthenticationService) *AdministratorService {
	return &AdministratorService{repository: repo, userAuthenticationService: uaService}
}

func (s *AdministratorService) CreateAdministrator(username string, password string, firstName string, lastName string, phoneNumber string, email string, office string) (models.Administrator, error) {
	//create administrator
	var administrator models.Administrator

	//Validate Username & Password
	validUsername := utils.ValidUsername(username)
	validPassword := utils.ValidPassword(password)
	if !validUsername || !validPassword {
		return administrator, errors.New(`Invalid username or password`)
	}

	//Create new user auth...this checks if user auth already exists.
	userAuth, err := s.userAuthenticationService.CreateNewUserAuthentication(username, password)
	if err != nil {
		return administrator, err
	}

	//Create administrator
	administrator, err = s.repository.CreateAdministrator(firstName, lastName, userAuth.ID, phoneNumber, email, office)

	return administrator, err
}

func (s *AdministratorService) GetAdministrators(queryParams models.AdministratorQueryParams) ([]models.Administrator, error) {
	//Validate query params
	if queryParams.Limit != nil && (*queryParams.Limit < 10 || *queryParams.Limit > 100) {
		return nil, errors.New("Limit is not within range of 10-100")
	}

	//Get administrators
	administrators, err := s.repository.GetAdministrators(queryParams)
	if err != nil {
		return nil, err
	}

	return administrators, nil
}

func (s *AdministratorService) GetAdministrator(username string) (models.Administrator, error) {
	var administrator models.Administrator

	//Validate username
	if !utils.ValidUsername(username) {
		return administrator, errors.New("Invalid username")
	}

	//Get administrator
	administrator, err := s.repository.GetAdministrator(username)

	return administrator, err
}

func (s *AdministratorService) UpdateAdministrator(username string, administratorUpdates models.AdministratorUpdate) (models.Administrator, error) {
	var administrator models.Administrator

	//Validate updates
	if !utils.ValidUsername(username) {
		return administrator, errors.New("Invalid username")
	}
	if administratorUpdates.FirstName != nil && !utils.ValidName(*administratorUpdates.FirstName) {
		return administrator, errors.New("Invalid first name")
	}
	if administratorUpdates.LastName != nil && !utils.ValidName(*administratorUpdates.LastName) {
		return administrator, errors.New("Invalid last name")
	}
	if administratorUpdates.PhoneNumber != nil && !utils.ValidPhoneNumber(*administratorUpdates.PhoneNumber) {
		return administrator, errors.New("Invalid phone number")
	}
	if administratorUpdates.Email != nil && !utils.ValidEmail(*administratorUpdates.Email) {
		return administrator, errors.New("Invalid email")
	}

	//Make sure there is at least one update. If not return
	administrator, err := s.repository.UpdateAdministrator(username, administratorUpdates)

	return administrator, err
}

func (s *AdministratorService) DeleteAdministrator(username string) (bool, error) {
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
