package services

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

func ConnectToDB() (*sql.DB, error) {
	dbURL, err := createDBURL("db")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getDBUsername() (string, error) {
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return "", errors.New("DB username not found")
	}
	return user, nil
}

func getDBPassword() (string, error) {
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return "", errors.New("DB password not found")
	}
	return password, nil
}

func getDBName() (string, error) {
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return "", errors.New("DB name not found")
	}
	return dbName, nil
}

func getDBPort() (string, error) {
	dbPort, ok := os.LookupEnv("DB_CONTAINER_PORT")
	if !ok {
		return "", errors.New("DB Container port not found")
	}
	return dbPort, nil
}

func createDBURL(dbHostName string) (string, error) {
	username, err := getDBUsername()
	if err != nil {
		return "", err
	}

	password, err := getDBPassword()
	if err != nil {
		return "", err
	}

	dbName, err := getDBName()
	if err != nil {
		return "", err
	}

	dbPort, err := getDBPort()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, dbHostName, dbPort, dbName)

	return url, nil
}
