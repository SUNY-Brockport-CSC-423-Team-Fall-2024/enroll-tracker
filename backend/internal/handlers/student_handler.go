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

		_, err := s.CreateNewStudent(studentCreation.Username, studentCreation.Password, studentCreation.First_Name, studentCreation.Last_Name, studentCreation.Phone_Number, studentCreation.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(201)
	}
}
