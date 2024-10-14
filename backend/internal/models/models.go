package models

import "time"

type StudentCreation struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Phone_Number string `json:"phone_number"`
	Email        string `json:"email"`
}

type Student struct {
	ID           int       `json:"id"`
	First_Name   string    `json:"first_name"`
	Last_Name    string    `json:"last_name"`
	Auth_ID      int       `json:"auth_id"`
	Major_ID     int       `json:"major_id"`
	Phone_Number string    `json:"phone_number"`
	Email        string    `json:"email"`
	Last_Login   time.Time `json:"last_login"`
	Created_At   time.Time `json:"created_at"`
	Updated_At   time.Time `json:"updated_at"`
}

type UserAuthentication struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}
