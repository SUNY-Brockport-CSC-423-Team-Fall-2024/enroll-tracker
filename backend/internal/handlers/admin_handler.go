package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"net/http"
    "strconv"
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

func GetAdministratorsHandler(s *services.AdministratorService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
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

		queryParams := models.AdministratorQueryParams{
			Limit:       limit,
			Offset:      offset,
			Username:    username,
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			PhoneNumber: phoneNumber,
			Office: office,
		}

		administrators, err := s.GetAdministrators(queryParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(administrators); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
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

func DeleteAdministratorHandler(adminService *services.AdministratorService, userSessionService *services.UserSessionService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		//Get username to delete
		username := r.PathValue("username")
		if username == "" {
			http.Error(w, "Administrator username not provided", http.StatusBadRequest)
			return
		}
		//Revoke any active user sessions associated with the user
		if _, err := userSessionService.RevokeUserSessionWithUsername(username); err != nil {
			http.Error(w, "Error occured when deleting administrator", http.StatusInternalServerError)
			return
		}
		//Delete administrator
		success, err := adminService.DeleteAdministrator(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !success {
			http.Error(w, "Error occured when deleting administrator", http.StatusInternalServerError)
			return
		}
		//Write 204 back to indicate successful deletion
		w.WriteHeader(http.StatusNoContent)
    }
}
