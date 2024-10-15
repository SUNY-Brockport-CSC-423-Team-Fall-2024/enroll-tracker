package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"net/http"
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

func GetStudentHandler(s *services.StudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var kv map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rawUsername, ok := kv["username"]
		if !ok {
			http.Error(w, "Username not supplied", http.StatusBadRequest)
			return
		}

		username, ok := rawUsername.(string)
		if !ok {
			http.Error(w, "Username invalid type", http.StatusBadRequest)
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

		if err := json.NewDecoder(r.Body).Decode(&studentUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		student, err := studentService.UpdateStudent(studentUpdate)
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
