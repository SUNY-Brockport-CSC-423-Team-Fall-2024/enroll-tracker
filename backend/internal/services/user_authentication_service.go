package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"errors"
	"regexp"
)

type UserAuthenticationService struct {
	repository repositories.UserAuthenticationRepository
}

func NewUserAuthenticationService(repo repositories.UserAuthenticationRepository) *UserAuthenticationService {
	return &UserAuthenticationService{repository: repo}
}

func (s *UserAuthenticationService) CreateNewUserAuthentication(username string, password string) (models.UserAuthentication, error) {
	var userAuth models.UserAuthentication

	//Validate Username & Password
	usernameRe := regexp.MustCompile(`^[a-z,A-Z,0-9]{8,50}$`)
	passwordRe := regexp.MustCompile(`^[a-z,A-Z,0-9,!,@,*,$,#]$`)
	if !usernameRe.MatchString(username) || !passwordRe.MatchString(password) {
		return userAuth, errors.New(`Invalid username or password`)
	}

	//Check user auth doesn't already exist
	ua, err := s.repository.GetUserAuthentication(username)
	if err != nil {
		return userAuth, err
	}
	if !utils.IsZeroValueStruct(ua) {
		return userAuth, errors.New(`User already exists`)
	}

	//Hash password
	passwordHash, err := utils.HashPassword(password, utils.CurArgon2IdParams)
	if err != nil {
		return userAuth, errors.New(`Error generating password`)
	}

	//Create user auth
	userAuth, err = s.repository.CreateUserAuthentication(username, passwordHash)
	return userAuth, err
}
