package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"errors"
	"net/http"
	"strconv"
)

func CreateMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode http request params
		var majorCreation models.MajorCreation
		if err := json.NewDecoder(r.Body).Decode(&majorCreation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := s.CreateMajor(majorCreation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
func GetMajorsHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query params we want to take
		query := r.URL.Query()
		limitParam := query.Get("limit")
		pageParam := query.Get("page")
		nameParam := query.Get("name")
		descriptionParam := query.Get("description")
		statusParam := query.Get("status")

		var limit *int
		var offset *int
		var name *string
		var description *string
		var status *string

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
		if nameParam != "" {
			name = new(string)
			(*name) = nameParam
		}
		if descriptionParam != "" {
			description = new(string)
			(*description) = descriptionParam
		}
		if statusParam != "" {
			status = new(string)
			(*status) = statusParam
		}

		queryParams := models.MajorQueryParams{
			Limit:       limit,
			Offset:      offset,
			Name:        name,
			Description: description,
			Status:      status,
		}

		majors, err := s.GetMajors(queryParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(majors); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
func GetMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get major id
		majorIDParam := r.PathValue("majorID")

		if majorIDParam == "" {
			http.Error(w, "Major ID not provided", http.StatusBadRequest)
			return
		}

		majorID, err := strconv.Atoi(majorIDParam)
		if err != nil {
			http.Error(w, "Major ID not in valid format", http.StatusBadRequest)
			return
		}

		major, err := s.GetMajor(majorID)
		if err != nil {
			http.Error(w, "Unable to get major", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(major); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
func UpdateMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var majorUpdate models.MajorUpdate
		majorIDParam := r.PathValue("majorID")

		if majorIDParam == "" {
			http.Error(w, "Major ID not provided", http.StatusBadRequest)
			return
		}

		majorID, err := strconv.Atoi(majorIDParam)
		if err != nil {
			http.Error(w, "Major ID not in valid format", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&majorUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		success, err := s.UpdateMajor(majorID, majorUpdate)
		if err != nil {
			if errors.Is(err, models.NoAffectedRows) {
				http.Error(w, "No major was found to update", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !success {
			http.Error(w, "Error occured while updating major", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
func DeleteMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get major id to delete
		majorIDParam := r.PathValue("majorID")

		if majorIDParam == "" {
			http.Error(w, "Major ID not provided", http.StatusBadRequest)
			return
		}

		majorID, err := strconv.Atoi(majorIDParam)
		if err != nil {
			http.Error(w, "Major ID not in valid format", http.StatusBadRequest)
			return
		}

		//Delete major
		success, err := s.DeleteMajor(majorID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !success {
			http.Error(w, "Error occured when deleting major", http.StatusBadRequest)
			return
		}
		//Write 204 back to indicate successful deletion
		w.WriteHeader(http.StatusNoContent)
	}
}

func AddCourseToMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		majorIDParam := r.PathValue("majorID")

		if majorIDParam == "" {
			http.Error(w, "Major ID not provided", http.StatusBadRequest)
			return
		}
		majorID, err := strconv.Atoi(majorIDParam)
		if err != nil {
			http.Error(w, "Major ID not in valid format", http.StatusBadRequest)
			return
		}

		//Get course id
		var kv map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		courseIDParam, ok := kv["courseID"]
		if !ok {
			http.Error(w, "Course ID not provided in body", http.StatusBadRequest)
			return
		}
		courseID, ok := courseIDParam.(float64) // by default numbers are decoded as float64 when interface{} is the value type
		if !ok {
			http.Error(w, "Course ID not in valid format", http.StatusBadRequest)
			return
		}

		success, err := s.AddCourseToMajor(majorID, int(courseID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !success {
			http.Error(w, "Unable to add course to major", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
func GetCoursesAssoicatedWithMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get major id
		majorIDParam := r.PathValue("majorID")
		if majorIDParam == "" {
			http.Error(w, "Major ID not provided", http.StatusBadRequest)
			return
		}
		majorID, err := strconv.Atoi(majorIDParam)
		if err != nil {
			http.Error(w, "Major ID not in valid format", http.StatusBadRequest)
			return
		}

		//Get query params
		query := r.URL.Query()
		limitParam := query.Get("limit")
		pageParam := query.Get("page")
		nameParam := query.Get("name")
		descriptionParam := query.Get("description")
		teacherParam := query.Get("teacher_id")
		maxEnrollmentParam := query.Get("max_enrollment")
		minEnrollmentParam := query.Get("min_enrollment")
		maxNumCreditsParam := query.Get("max_num_credits")
		minNumCreditsParam := query.Get("min_num_credits")
		statusParam := query.Get("status")

		var limit *int
		var offset *int
		var name *string
		var description *string
		var teacherID *int
		var maxEnrollment *int
		var minEnrollment *int
		var maxNumCredits *int
		var minNumCredits *int
		var status *string

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
		if nameParam != "" {
			name = new(string)
			(*name) = nameParam
		}
		if descriptionParam != "" {
			description = new(string)
			(*description) = descriptionParam
		}
		if teacherParam != "" {
			i, err := strconv.Atoi(teacherParam)
			if err != nil {
				http.Error(w, "Invalid teacher id param", http.StatusBadRequest)
				return
			}
			teacherID = new(int)
			*teacherID = i
		}
		if maxEnrollmentParam != "" {
			i, err := strconv.Atoi(maxEnrollmentParam)
			if err != nil {
				http.Error(w, "Invalid max enrollment param", http.StatusBadRequest)
				return
			}
			maxEnrollment = new(int)
			*maxEnrollment = i
		}
		if minEnrollmentParam != "" {
			i, err := strconv.Atoi(minEnrollmentParam)
			if err != nil {
				http.Error(w, "Invalid min enrollment param", http.StatusBadRequest)
				return
			}
			minEnrollment = new(int)
			*minEnrollment = i
		}
		if maxNumCreditsParam != "" {
			i, err := strconv.Atoi(maxNumCreditsParam)
			if err != nil {
				http.Error(w, "Invalid max num credits param", http.StatusBadRequest)
				return
			}
			maxNumCredits = new(int)
			*maxNumCredits = i
		}
		if minNumCreditsParam != "" {
			i, err := strconv.Atoi(minNumCreditsParam)
			if err != nil {
				http.Error(w, "Invalid min num credits param", http.StatusBadRequest)
				return
			}
			minNumCredits = new(int)
			*minNumCredits = i
		}
		if statusParam != "" {
			status = new(string)
			(*status) = statusParam
		}

		queryParams := models.CourseQueryParams{
			Limit:         limit,
			Offset:        offset,
			Name:          name,
			Description:   description,
			TeacherID:     teacherID,
			MaxEnrollment: maxEnrollment,
			MinEnrollment: minEnrollment,
			MaxNumCredits: maxNumCredits,
			MinNumCredits: minNumCredits,
			Status:        status,
		}

		courses, err := s.GetCoursesAssociatedWithMajor(majorID, queryParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(courses); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
func DeleteCourseFromMajorHandler(s *services.MajorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		majorIDParam := r.PathValue("majorID")
		courseIDParam := r.PathValue("courseID")

		if majorIDParam == "" {
			http.Error(w, "Major ID not provided", http.StatusBadRequest)
			return
		}
		if courseIDParam == "" {
			http.Error(w, "Course ID not provided", http.StatusBadRequest)
			return
		}

		majorID, err := strconv.Atoi(majorIDParam)
		if err != nil {
			http.Error(w, "Major ID not in valid format", http.StatusBadRequest)
			return
		}
		courseID, err := strconv.Atoi(courseIDParam)
		if err != nil {
			http.Error(w, "Course ID not in valid format", http.StatusBadRequest)
			return
		}

		success, err := s.RemoveCourseFromMajor(majorID, courseID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !success {
			http.Error(w, "Unable to remove course from major", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
