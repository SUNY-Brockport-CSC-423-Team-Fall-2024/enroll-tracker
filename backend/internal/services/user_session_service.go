package services

import (
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/repositories"
	"enroll-tracker/pkg/utils"
	"time"
)

type UserSessionService struct {
	repository                repositories.UserSessionRepository
	userAuthenticationService *UserAuthenticationService
}

func CreateUserSessionService(repo repositories.UserSessionRepository) *UserSessionService {
	return &UserSessionService{repository: repo}
}

// Returns a UserSession object, access token, and an error
// The UserSession object contains the plaintext version of the refresh token
// Needs UserAuth table ID and the appropriate username
func (s *UserSessionService) CreateUserSession(userID int, username string) (*models.UserSession, string, error) {
	//Create refresh token
	refreshToken, err := utils.CreateRefreshToken(utils.RefreshTokenLength)
	if err != nil {
		return nil, "", err
	}
	//Hash refresh token
	hashedRefreshToken, err := utils.HashText(refreshToken, utils.CurArgon2IdParams)
	if err != nil {
		return nil, "", err
	}

	//Create claims for token
	role, err := s.repository.GetUserRole(username)
	if err != nil {
		return nil, "", err
	}

	accessTokenExpiresAt := time.Now().Add(time.Minute * 15)
	issuedAt := time.Now()
	notBefore := time.Now()
	accessToken, err := utils.CreateJWT(username, role, accessTokenExpiresAt, issuedAt, notBefore)
	if err != nil {
		return nil, "", err
	}

	//Create refresh token ID
	refreshTokenID := utils.CreateUUID()

	//Create refresh token exp date
	refreshTokenExpiresIn := time.Now().Add((time.Hour * 24) * 7) //refresh tokens expire in 7 days

	//Create user session
	userSession, err := s.repository.CreateUserSession(userID, username, hashedRefreshToken, refreshTokenID, issuedAt, refreshTokenExpiresIn)
	if err != nil {
		return nil, "", err
	}

	//Set refresh token on user session object to unhashed version so we can send back to user
	userSession.RefreshToken = refreshToken

	return &userSession, accessToken, nil
}

func (s *UserSessionService) GetUserSession(refreshTokenID string) (models.UserSession, error) {
	var userSession models.UserSession

	userSession, err := s.repository.GetUserSession(refreshTokenID)
	if err != nil {
		return userSession, err
	}

	return userSession, nil
}

func (s *UserSessionService) IsUserSessionValid(userSession models.UserSession) bool {
	//check if session is revoked
	if userSession.Revoked {
		return false
	}

	//Refresh token is past and not revoked. Revoke and return false
	if time.Now().Compare(userSession.ExpiresAt) == 1 {
		_ = s.repository.RevokeUserSession(userSession.RefreshTokenID)
		return false
	}

	return true
}

// Attempts to revoke user session in repo layer
//
// # Returns false, error if there is an error when revoking the user session
//
// Return true, nil if the user session was successfully revoked
func (s *UserSessionService) RevokeUserSession(refreshTokenID string) (bool, error) {
	if err := s.repository.RevokeUserSession(refreshTokenID); err != nil {
		return false, err
	}
	return true, nil
}
