package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"net/http"
	"strconv"
)

func CreateStudentHandler(s *services.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode http request params
		var studentCreation models.StudentCreation
		if err := json.NewDecoder(r.Body).Decode(&studentCreation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := s.CreateStudent(studentCreation.Username, studentCreation.Password, studentCreation.FirstName, studentCreation.LastName, studentCreation.PhoneNumber, studentCreation.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetStudentsHandler(s *services.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query params we want to take
		query := r.URL.Query()
		limitParam := query.Get("limit")
		pageParam := query.Get("page")
		usernameParam := query.Get("username")
		firstNameParam := query.Get("first_name")
		lastNameParam := query.Get("last_name")
		phoneNumberParam := query.Get("phone_number")
		majorParam := query["major"]
		emailParam := query.Get("email")

		var limit *int
		var offset *int
		var username *string
		var firstName *string
		var lastName *string
		var phoneNumber *string
		var email *string
		var majors *[]string

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
		if len(majorParam) > 0 {
			majors = new([]string)
			*majors = majorParam
		}

		queryParams := models.StudentQueryParams{
			Limit:       limit,
			Offset:      offset,
			Username:    username,
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			PhoneNumber: phoneNumber,
			Majors:      majors,
		}

		students, err := s.GetStudents(queryParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(students); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetStudentHandler(s *services.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get username
		username := r.PathValue("username")

		if username == "" {
			http.Error(w, "Username not provided", http.StatusBadRequest)
			return
		}

		student, err := s.GetStudent(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(student); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateStudentHandler(studentService *services.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studentUpdate models.StudentUpdate
		username := r.PathValue("username")

		if username == "" {
			http.Error(w, "Username not provided", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&studentUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		student, err := studentService.UpdateStudent(username, studentUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&student); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
