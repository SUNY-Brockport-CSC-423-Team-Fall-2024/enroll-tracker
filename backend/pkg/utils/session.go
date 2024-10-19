package utils

import (
	"crypto/rand"
	"encoding/base64"
	"enroll-tracker/internal/models"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateUUID() string {
	return uuid.NewString()
}

func CreateJWT(username string, role string, expiresAt time.Time, issuedAt time.Time, notBefore time.Time) (string, error) {
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
		Role: role,
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

func CreateRefreshToken(length int) (string, error) {
	bytes := make([]byte, length)

	_, err := rand.Reader.Read(bytes)
	if err != nil {
		return "", errors.New("Unable to create refresh token")
	}

	refreshToken := base64.RawStdEncoding.EncodeToString(bytes)

	return refreshToken, nil
}

func ExtractAccessTokenFromAuthHeader(bearerToken string) string {
	return strings.TrimSpace(strings.TrimPrefix(bearerToken, "Bearer"))
}
