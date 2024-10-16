package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"errors"
	"time"
)

type UserSessionService struct {
	repository                repositories.UserSessionRepository
	userAuthenticationService *UserAuthenticationService
}

func CreateUserSessionService(repo repositories.UserSessionRepository) *UserSessionService {
	return &UserSessionService{repository: repo}
}

func (s *UserSessionService) CreateUserSession(userID int, username string) (*models.UserSession, string, error) {
	sessID := utils.CreateSessionID()
	expiresAt := time.Now().Add(time.Second * 15)
	issuedAt := time.Now()
	notBefore := time.Now()
	token, err := utils.CreateJWT(sessID, username, expiresAt, issuedAt, notBefore)
	if err != nil {
		return nil, "", err
	}

	userSession, err := s.repository.CreateUserSession(userID, sessID, issuedAt, expiresAt, notBefore)
	if err != nil {
		return nil, "", err
	}

	return &userSession, token, nil
}

func (s *UserSessionService) GetUserSession(tokenString string) (models.UserSession, error) {
	//verify token
	var userSession models.UserSession
	verifiedClaims, err := utils.VerifyJWT(tokenString)
	if err != nil {
		return userSession, err
	}
	if verifiedClaims == nil {
		return userSession, errors.New("No verified JWT")
	}

	userSession, err = s.repository.GetUserSession(verifiedClaims.SessID)
	if err != nil {
		return userSession, err
	}

	//Check if current time is after usersession exp time
	if time.Now().Compare(verifiedClaims.ExpiresAt.Time) == 1 {
		return userSession, errors.New("Session expired")
	}

	return userSession, nil
}
