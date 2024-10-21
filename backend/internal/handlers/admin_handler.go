package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"net/http"
)

func CreateAdministratorHandler(s *services.AdministratorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode http request params
		var administratorCreation models.AdministratorCreation
		if err := json.NewDecoder(r.Body).Decode(&administratorCreation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := s.CreateAdministrator(administratorCreation.Username, administratorCreation.Password, administratorCreation.FirstName, administratorCreation.LastName, administratorCreation.PhoneNumber, administratorCreation.Email, administratorCreation.Office)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAdministratorHandler(s *services.AdministratorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get username
		username := r.PathValue("username")

		if username == "" {
			http.Error(w, "Username not provided", http.StatusBadRequest)
			return
		}

		administrator, err := s.GetAdministrator(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(administrator); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func UpdateAdministratorHandler(administratorService *services.AdministratorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var administratorUpdate models.AdministratorUpdate
		username := r.PathValue("username")

		if username == "" {
			http.Error(w, "Username not provided", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&administratorUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		administrator, err := administratorService.UpdateAdministrator(username, administratorUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(&administrator); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
