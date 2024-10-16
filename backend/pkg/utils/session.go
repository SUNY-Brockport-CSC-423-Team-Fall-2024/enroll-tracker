package utils

import (
	"enroll-tracker/internal/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateSessionID() string {
	return uuid.NewString()
}

func CreateJWT(sessId string, username string, expiresAt time.Time, issuedAt time.Time, notBefore time.Time) (string, error) {
	signingKey, ok := os.LookupEnv("ENROLL_TRACKER_RSA_PRIVATE_KEY")
	if !ok {
		return "", errors.New("Can't sign JWT")
	}

	claims := models.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt), // Expire in 15 minutes
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			NotBefore: jwt.NewNumericDate(notBefore),
			Issuer:    "enroll-tracker",
			Subject:   username,
			Audience:  jwt.ClaimStrings{"enroll-tracker-client"},
		},
		SessID: sessId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)

	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(tokenString string) (*models.CustomClaims, error) {
	var claims models.CustomClaims
	//Parse token
	signingKey, ok := os.LookupEnv("ENROLL_TRACKER_RSA_PRIVATE_KEY")
	if !ok {
		return nil, errors.New("Can't verify JWT")
	}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return &claims, nil
}
