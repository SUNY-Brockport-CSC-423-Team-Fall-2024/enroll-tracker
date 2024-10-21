package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"net/http"
)

func CreateTeacherHandler(s *services.TeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func GetTeacherHandler(s *services.TeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func UpdateTeacherHandler(teacherService *services.TeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
