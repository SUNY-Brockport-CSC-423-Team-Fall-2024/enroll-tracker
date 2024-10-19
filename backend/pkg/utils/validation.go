package utils

import "regexp"

var RegexRules = struct {
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	Name        string
}{
	Username:    `^[a-z,A-Z,0-9]{8,50}$`,
	Password:    `^[a-zA-Z0-9!@#$%^&*()_+=\[\]{};':"\\|,.<>\/?~-]{8,30}$`,
	Email:       `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
	PhoneNumber: `^\+?(\d{1,3})?[-. (]*\d{3}[-. )]*\d{3}[-. ]*\d{4}$`,
	Name:        `^[a-zA-Z]+([ '-][a-zA-Z]+)*$`,
}

func ValidUsername(username string) bool {
	usernameRe := regexp.MustCompile(RegexRules.Username)
	if !usernameRe.MatchString(username) {
		return false
	}
	return true
}

func ValidPassword(password string) bool {
	passwordRe := regexp.MustCompile(RegexRules.Password)
	if !passwordRe.MatchString(password) {
		return false
	}
	return true
}

func ValidEmail(email string) bool {
	emailRe := regexp.MustCompile(RegexRules.Email)
	if !emailRe.MatchString(email) {
		return false
	}
	return true
}

func ValidPhoneNumber(phoneNumber string) bool {
	phoneNumberRe := regexp.MustCompile(RegexRules.PhoneNumber)
	if !phoneNumberRe.MatchString(phoneNumber) {
		return false
	}
	return true
}

func ValidName(name string) bool {
	nameRe := regexp.MustCompile(RegexRules.Name)
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
