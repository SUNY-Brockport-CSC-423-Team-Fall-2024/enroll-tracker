package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/models"
	"enroll-tracker/internal/services"
	"errors"
	"net/http"
	"strconv"
)

func CreateCourseHandler(s *services.CourseService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Decode http request params
	var courseCreation models.CourseCreation
	if err := json.NewDecoder(r.Body).Decode(&courseCreation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := s.CreateCourse(courseCreation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func GetCoursesHandler(s *services.CourseService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Query params we want to take
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

	courses, err := s.GetCourses(queryParams)
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
func GetCourseHandler(s *services.CourseService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Get course id
	courseIDParam := r.PathValue("courseID")

	if courseIDParam == "" {
		http.Error(w, "Course ID not provided", http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		http.Error(w, "Course ID not in valid format", http.StatusBadRequest)
		return
	}

	course, err := s.GetCourse(courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(course); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func UpdateCourseHandler(s *services.CourseService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	var courseUpdate models.CourseUpdate
	courseIDParam := r.PathValue("courseID")

	if courseIDParam == "" {
		http.Error(w, "Course ID not provided", http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		http.Error(w, "Course ID not in valid format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&courseUpdate); err != nil {
		http.Error(w, "Error ocurred while updating course", http.StatusInternalServerError)
		return
	}

	success, err := s.UpdateCourse(courseID, courseUpdate)
	if err != nil {
		if errors.Is(err, models.NoAffectedRows) {
			http.Error(w, "No course was found to update", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !success {
		http.Error(w, "Error occured while updating course", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func DeleteCourseHandler(s *services.CourseService, w http.ResponseWriter, r *http.Request) {
	//Set CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//Get course id to delete
	courseIDParam := r.PathValue("courseID")

	if courseIDParam == "" {
		http.Error(w, "Course ID not provided", http.StatusBadRequest)
		return
	}

	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		http.Error(w, "Course ID not in valid format", http.StatusBadRequest)
		return
	}

	//Delete course
	success, err := s.DeleteCourse(courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !success {
		http.Error(w, "Error occured when deleting course", http.StatusBadRequest)
		return
	}
	//Write 204 back to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}
