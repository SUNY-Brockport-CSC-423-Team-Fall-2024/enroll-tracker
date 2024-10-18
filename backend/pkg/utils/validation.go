package utils

import "regexp"

func ValidUsername(username string) bool {
	usernameRe := regexp.MustCompile(`^[a-z,A-Z,0-9]{8,50}$`)
	if !usernameRe.MatchString(username) {
		return false
	}
	return true
}

func ValidPassword(password string) bool {
	passwordRe := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+=\[\]{};':"\\|,.<>\/?~-]{8,30}$`)
	if !passwordRe.MatchString(password) {
		return false
	}
	return true
}

func ValidEmail(email string) bool {
	emailRe := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRe.MatchString(email) {
		return false
	}
	return true
}

func ValidPhoneNumber(phoneNumber string) bool {
	phoneNumberRe := regexp.MustCompile(`^\+?(\d{1,3})?[-. (]*\d{3}[-. )]*\d{3}[-. ]*\d{4}$`)
	if !phoneNumberRe.MatchString(phoneNumber) {
		return false
	}
	return true
}

func ValidName(name string) bool {
	nameRe := regexp.MustCompile(`^[a-zA-Z]+([ '-][a-zA-Z]+)*$`)
	if !nameRe.MatchString(name) {
		return false
	}
	return true
}

func IsString(text interface{}) bool {
	if _, ok := text.(string); !ok {
		return false
	}
	return true
}
