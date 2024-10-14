package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	CurArgon2IdParams = Argon2IdParams{
		Version:    argon2.Version,
		Memory:     7168,
		Iterations: 5,
		Threads:    1,
		KeyLength:  128,
	} //Recommendations from OWASP
)

type Argon2IdParams struct {
	Version    uint8
	Memory     uint32
	Iterations uint32
	Threads    uint8 //Paralleism degree
	KeyLength  uint32
}

func generateSalt(len uint32) ([]byte, error) {
	salt := make([]byte, len)
	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string, params Argon2IdParams) (string, error) {
	salt, err := generateSalt(10)

	if err != nil {
		return "", errors.New(`Error generating salt`)
	}

	hash := argon2.IDKey([]byte(password),
		salt,
		CurArgon2IdParams.Iterations,
		CurArgon2IdParams.Memory,
		CurArgon2IdParams.Threads,
		CurArgon2IdParams.KeyLength,
	)

	encodedPasswordHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	encodedPassInfo := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", CurArgon2IdParams.Version, CurArgon2IdParams.Memory, CurArgon2IdParams.Iterations, CurArgon2IdParams.Threads, encodedSalt, encodedPasswordHash)

	return encodedPassInfo, nil
}

func VerifyPassword(password string, encodedPasswordHash string) (bool, error) {
	//decode argon2 params, salt, and digest before hashing password
	parts := strings.Split(encodedPasswordHash, "$")
	var argonParams Argon2IdParams

	if len(parts) != 6 {
		return false, errors.New("Invalid encoded password hash")
	}

	//read version
	_, err := fmt.Sscanf(parts[2], "v=%d", &argonParams.Version)
	if err != nil {
		return false, errors.New("Invalid version in encoded password hash")
	}

	//read memory, time, threads
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &argonParams.Memory, &argonParams.Iterations, &argonParams.Threads)
	if err != nil {
		return false, err
	}

	//read salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	//read password digest
	digest, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	//hash provided password with argon2id params used when password was set
	newHash := argon2.IDKey([]byte(password),
		salt,
		argonParams.Iterations,
		argonParams.Memory,
		argonParams.Threads,
		argonParams.KeyLength,
	)

	//Compare two digest
	return subtle.ConstantTimeCompare(newHash, digest) == 1, nil
}
