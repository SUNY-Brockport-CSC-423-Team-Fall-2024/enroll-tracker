package utils

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"enroll-tracker/internal/models"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateUUID() string {
	return uuid.NewString()
}

func CreateJWT(username string, userID int, role string, expiresAt time.Time, issuedAt time.Time, notBefore time.Time) (string, error) {
	signingKeyEncoded, ok := os.LookupEnv("ENROLL_TRACKER_RSA_PRIVATE_KEY")
	if !ok {
		return "", errors.New("Can't sign JWT")
	}
	signingKeyDecoded, err := base64.StdEncoding.DecodeString(signingKeyEncoded)
	signingKey := formatKeyToPEM(string(signingKeyDecoded), "PRIVATE")
	block, _ := pem.Decode([]byte(signingKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("Failed to parse RSA private key PEM")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", errors.New("Failed to parse RSA private key")
	}

	claims := models.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			NotBefore: jwt.NewNumericDate(notBefore),
			Issuer:    "enroll-tracker",
			Subject:   username,
			Audience:  jwt.ClaimStrings{"enroll-tracker-client"},
		},
		Role:   role,
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(tokenString string) (*models.CustomClaims, error) {
	var claims models.CustomClaims
	//Parse token
	signingKeyEncoded, ok := os.LookupEnv("ENROLL_TRACKER_RSA_PUBLIC_KEY")
	if !ok {
		return nil, errors.New("JWT public key not found")
	}
	verifyingKeyDecoded, err := base64.StdEncoding.DecodeString(signingKeyEncoded)
	verifyingKey := formatKeyToPEM(string(verifyingKeyDecoded), "PUBLIC")
	block, _ := pem.Decode([]byte(verifyingKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("Failed to parse RSA public key PEM")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Failed to parse RSA public key")
	}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}, jwt.WithValidMethods([]string{"RS256"}))
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

// Converts base64 encoded string to utf-16 in PEM format
func formatKeyToPEM(base64key string, keyType string) string {
	var formattedKey strings.Builder

	formattedKey.WriteString(fmt.Sprintf("-----BEGIN %s KEY-----\n", keyType))

	for i := 0; i < len(base64key); i += 64 {
		end := i + 64
		if end > len(base64key) {
			end = len(base64key)
		}
		formattedKey.WriteString(base64key[i:end] + "\n")
	}

	formattedKey.WriteString(fmt.Sprintf("-----END %s KEY-----\n", keyType))

	return formattedKey.String()
}
