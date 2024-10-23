package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"errors"
)

type UserAuthenticationService struct {
	repository repositories.UserAuthenticationRepository
}

func CreateUserAuthenticationService(repo repositories.UserAuthenticationRepository) *UserAuthenticationService {
	return &UserAuthenticationService{repository: repo}
}

func (s *UserAuthenticationService) CreateNewUserAuthentication(username string, password string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	//Check user auth doesn't already exist
	ua, err := s.repository.GetUserAuthentication(username)
	if err != nil {
		return userAuth, err
	}
	if !utils.IsZeroValueStruct(ua) {
		return userAuth, errors.New(`User already exists`)
	}

	//Hash password
	passwordHash, err := utils.HashText(password, utils.CurArgon2IdParams)
	if err != nil {
		return userAuth, errors.New(`Error generating password`)
	}

	//Create user auth
	userAuth, err = s.repository.CreateUserAuthentication(username, passwordHash)
	return userAuth, err
}

func (s *UserAuthenticationService) GetUserAuthentication(username string) (*models.UserAuthentication, error) {
	userAuth, err := s.repository.GetUserAuthentication(username)
	if err != nil {
		return userAuth, err
	}
	return userAuth, nil
}

func (s *UserAuthenticationService) ChangePassword(username string, newPassword string, oldPassword string) (*models.UserAuthentication, error) {
	//TODO: Reduce number of calls to db by not calling two separate postgres functions
	ua, err := s.GetUserAuthentication(username)
	if err != nil {
		return nil, err
	}

	//Check if we actually got a user
	if ua == nil {
		return nil, errors.New("No user found to change password for.")
	}

	//Verify old password
	same, err := utils.VerifyHashedText(oldPassword, ua.PasswordHash)
	if err != nil {
		return nil, err
	}

	//Check if the passwords weren't the same
	if !same {
		return nil, errors.New("Old password is not correct.")
	}

	//Verify new password meets requirements
	if !utils.ValidPassword(newPassword) {
		return nil, errors.New("New password doesn't meet password requirements.")
	}

	//Hash new password
	hashedNewPassword, err := utils.HashText(newPassword, utils.CurArgon2IdParams)
	if err != nil {
		return nil, errors.New("Error when changing password.")
	}

	//Set new password
	newUa, err := s.repository.ChangePassword(ua.ID, hashedNewPassword)
	if err != nil {
		return nil, errors.New("Error when changing password.")
	}

	return newUa, nil
}
