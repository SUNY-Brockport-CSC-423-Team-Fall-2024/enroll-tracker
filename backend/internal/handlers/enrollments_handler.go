package handlers

import (
	"encoding/json"
	"enroll-tracker/internal/services"
	"net/http"
	"strconv"
)

func EnrollStudentHandler(s *services.EnrollmentsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseIDParam := r.PathValue("courseID")
		studentIDParam := r.PathValue("studentID")

		if courseIDParam == "" {
			http.Error(w, "Missing course id", http.StatusBadRequest)
			return
		}
		if studentIDParam == "" {
			http.Error(w, "Missing student id", http.StatusBadRequest)
			return
		}

		courseID, err := strconv.Atoi(courseIDParam)
		if err != nil {
			http.Error(w, "Course id invalid format", http.StatusBadRequest)
			return
		}
		studentID, err := strconv.Atoi(studentIDParam)
		if err != nil {
			http.Error(w, "Course id invalid format", http.StatusBadRequest)
			return
		}

		success, err := s.EnrollStudent(courseID, studentID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !success {
			http.Error(w, "Unable to enroll student in course", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UnenrollStudentHandler(s *services.EnrollmentsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseIDParam := r.PathValue("courseID")
		studentIDParam := r.PathValue("studentID")

		if courseIDParam == "" {
			http.Error(w, "Missing course id", http.StatusBadRequest)
			return
		}
		if studentIDParam == "" {
			http.Error(w, "Missing student id", http.StatusBadRequest)
			return
		}

		courseID, err := strconv.Atoi(courseIDParam)
		if err != nil {
			http.Error(w, "Course id invalid format", http.StatusBadRequest)
			return
		}
		studentID, err := strconv.Atoi(studentIDParam)
		if err != nil {
			http.Error(w, "Course id invalid format", http.StatusBadRequest)
			return
		}

		success, err := s.UnenrollStudent(courseID, studentID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !success {
			http.Error(w, "Unable to unenroll student in course", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetCoursesStudentsHandler(s *services.EnrollmentsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courseIDParam := r.PathValue("courseID")
		isEnrolledParam := r.PathValue("isEnrolled")

		var isEnrolled *bool = nil

		if courseIDParam == "" {
			http.Error(w, "Missing course id", http.StatusBadRequest)
			return
		}

		courseID, err := strconv.Atoi(courseIDParam)
		if err != nil {
			http.Error(w, "Course id invalid format", http.StatusBadRequest)
			return
		}

		if isEnrolledParam != "" {
			isEnrolled = new(bool)
			(*isEnrolled), err = strconv.ParseBool(isEnrolledParam)
			if err != nil {
				http.Error(w, "Is enrolled invalid format", http.StatusBadRequest)
				return
			}
		}

		students, err := s.GetCoursesStudents(courseID, isEnrolled)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(students); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetStudentsCoursesHandler(s *services.EnrollmentsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentIDParam := r.PathValue("studentID")
		isEnrolledParam := r.PathValue("isEnrolled")

		var isEnrolled *bool = nil

		if studentIDParam == "" {
			http.Error(w, "Missing student id", http.StatusBadRequest)
			return
		}

		studentID, err := strconv.Atoi(studentIDParam)
		if err != nil {
			http.Error(w, "Student id invalid format", http.StatusBadRequest)
			return
		}

		if isEnrolledParam != "" {
			isEnrolled = new(bool)
			(*isEnrolled), err = strconv.ParseBool(isEnrolledParam)
			if err != nil {
				http.Error(w, "Is enrolled invalid format", http.StatusBadRequest)
				return
			}
		}

		courses, err := s.GetStudentsCourses(studentID, isEnrolled)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(courses); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
