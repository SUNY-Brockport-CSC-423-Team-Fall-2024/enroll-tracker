package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"net/http"
	"strconv"
)

func CreateTeacherHandler(s *services.TeacherService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Decode http request params
	var teacherCreation models.TeacherCreation
	if err := json.NewDecoder(r.Body).Decode(&teacherCreation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := s.CreateTeacher(teacherCreation.Username, teacherCreation.Password, teacherCreation.FirstName, teacherCreation.LastName, teacherCreation.PhoneNumber, teacherCreation.Email, teacherCreation.Office)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetTeachersHandler(s *services.TeacherService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Query params we want to take
	query := r.URL.Query()
	limitParam := query.Get("limit")
	pageParam := query.Get("page")
	usernameParam := query.Get("username")
	firstNameParam := query.Get("first_name")
	lastNameParam := query.Get("last_name")
	phoneNumberParam := query.Get("phone_number")
	emailParam := query.Get("email")
	officeParam := query.Get("office")

	var limit *int
	var offset *int
	var username *string
	var firstName *string
	var lastName *string
	var phoneNumber *string
	var email *string
	var office *string

	//Get parameters
	if limitParam != "" {
		i, err := strconv.Atoi(limitParam)
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		limit = new(int)
		*limit = i
	}
	if pageParam != "" && limit != nil {
		i, err := strconv.Atoi(pageParam)
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		//Can't have page 0. Page num starts at 1
		if i <= 0 {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		} else if i == 1 {
			offset = new(int)
			*offset = 0
		} else {
			//If we want the 3rd page with 10 records a page. We want to offset the query by 20 so limit * page - 1
			offset = new(int)
			*offset = (*limit) * (i - 1)
		}
	}
	if firstNameParam != "" {
		firstName = new(string)
		(*firstName) = firstNameParam
	}
	if lastNameParam != "" {
		lastName = new(string)
		*lastName = lastNameParam
	}
	if usernameParam != "" {
		username = new(string)
		*username = usernameParam
	}
	if emailParam != "" {
		email = new(string)
		*email = emailParam
	}
	if phoneNumberParam != "" {
		phoneNumber = new(string)
		*phoneNumber = phoneNumberParam
	}
	if officeParam != "" {
		office = new(string)
		*office = officeParam
	}

	queryParams := models.TeacherQueryParams{
		Limit:       limit,
		Offset:      offset,
		Username:    username,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Office:      office,
	}

	teachers, err := s.GetTeachers(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(teachers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetTeacherHandler(s *services.TeacherService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Get username
	username := r.PathValue("username")

	if username == "" {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	teacher, err := s.GetTeacher(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateTeacherHandler(teacherService *services.TeacherService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	var teacherUpdate models.TeacherUpdate
	username := r.PathValue("username")

	if username == "" {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&teacherUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teacher, err := teacherService.UpdateTeacher(username, teacherUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTeacherHandler(teacherService *services.TeacherService, userSessionService *services.UserSessionService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Get username to delete
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Teacher username not provided", http.StatusBadRequest)
		return
	}
	//Revoke any active user sessions associated with the user
	if _, err := userSessionService.RevokeUserSessionWithUsername(username); err != nil {
		http.Error(w, "Error occured when deleting teacher", http.StatusInternalServerError)
		return
	}
	//Delete teacher
	success, err := teacherService.DeleteTeacher(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !success {
		http.Error(w, "Error occured when deleting teacher", http.StatusInternalServerError)
		return
	}
	//Write 204 back to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}
