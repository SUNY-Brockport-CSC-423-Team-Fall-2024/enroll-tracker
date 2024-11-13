"use client";

import { useState, useEffect } from "react";
import styles from "./styles.module.css";
import { useAuth } from "@/app/providers/auth-provider";
import { useRouter } from "next/navigation";

interface Course {
  course_id: number;
  course_name: string;
  course_description: string;
  teacher_id: number;
  max_enrollment: number;
  num_credits: number;
  status: string;
  last_updated: string;
  created_at: string;
  is_enrolled: boolean;
  enrolled_date: string;
  unenrolled_date: string | null;
}

interface Teacher {
  id: number;
  last_name: string;
}

export default function Courses() {
  const router = useRouter();
  const { userRole, userID, username } = useAuth();
  const [selectedButton, setSelectedButton] = useState<string>(
    userRole === "instructor" ? "Active" : "My Major",
  );
  const [enrolledCourses, setEnrolledCourses] = useState<Course[]>([]);
  const [majorCourses, setMajorCourses] = useState<Course[]>([]);
  const [droppedCourses, setDroppedCourses] = useState<Course[]>([]);
  const [teacherNames, setTeacherNames] = useState<{ [id: number]: string }>({});
  const [error, setError] = useState<string | null>(null);
  const [selectedCourseID, setSelectedCourseID] = useState<number | null>(null);

  const handleButtonClick = (button: string) => {
    setSelectedButton(button);
  };
  // Fetch enrolled courses when the "Enrolled" button is selected
  useEffect(() => {
    if (userID) {
      const fetchCourses = async () => {
        try {
          let url = "";
          let data: Course[] = [];

          if (selectedButton === "Enrolled" || selectedButton === "Dropped") {
            // Fetch enrolled or dropped courses
            url = `http://localhost:8002/api/students/${userID}/courses`;
            const response = await fetch(url);
            if (!response.ok) {
              throw new Error("Error fetching courses");
            }
            data = await response.json();

            // Filter courses based on `is_enrolled` status
            if (selectedButton === "Enrolled") {
              setEnrolledCourses(data.filter((course) => course.is_enrolled));
            } else if (selectedButton === "Dropped") {
              setDroppedCourses(data.filter((course) => !course.is_enrolled));
            }
          } else if (selectedButton === "My Major") {
            // Fetch courses for the major
            url = `http://localhost:8002/api/majors/${userID}/courses`;
            const response = await fetch(url);
            if (!response.ok) {
              throw new Error("Error fetching major courses");
            }
            data = await response.json();
            // Normalize major courses (same structure as enrolled/dropped courses)
            const normalizedCourses = data.map((course: any) => ({
              ...course,
              course_id: course.course_id || course.id, // Ensure uniform naming
              course_name: course.name || course.course_name, // Ensure uniform naming
            }));
            setMajorCourses(normalizedCourses); // Set major courses with normalization
          }
        } catch (err) {
          setError(err instanceof Error ? err.message : "An unknown error occurred");
        }
      };

      fetchCourses();
    }
  }, [selectedButton, userID]); // Trigger fetch when selectedButton or userID changes

  // Fetch all teachers and store in an object when the component mounts
  useEffect(() => {
    const fetchTeachers = async () => {
      try {
        const response = await fetch("http://localhost:8002/api/teachers");
        if (!response.ok) throw new Error("Error fetching teachers");

        const data: Teacher[] = await response.json();
        const teacherDict = data.reduce(
          (acc, teacher) => {
            acc[teacher.id] = teacher.last_name;
            return acc;
          },
          {} as { [id: number]: string },
        );
        console.log(teacherDict); // Check the teacherDict here
        setTeacherNames(teacherDict);
      } catch (err) {
        setError("Failed to fetch teacher last names");
      }
    };

    fetchTeachers();
  }, []);
  // Use useEffect for navigation
  useEffect(() => {
    if (selectedCourseID) {
      router.push(`/courses/${selectedCourseID}`);
    }
  }, [selectedCourseID, router]);

  // Handle course click by setting the selected course ID
  const handleCourseClick = (courseID: number) => {
    setSelectedCourseID(courseID);
  };

  const buttons =
    userRole === "instructor" ? ["Active", "Inactive"] : ["My Major", "Enrolled", "Dropped"];

  return (
    <div className={styles.courses_root}>
      <header className={styles.header}>
        <h1>Courses</h1>
        <button className={styles.right_button}>{username}</button>
      </header>

      <nav className={styles.nav_bar}>
        {buttons.map((button) => (
          <button
            key={button}
            onClick={() => handleButtonClick(button)}
            className={`${styles.nav_button} ${selectedButton === button ? styles.selected : ""}`}
          >
            {button}
          </button>
        ))}
      </nav>

      <div className={styles.scroll_list}>
        {selectedButton === "Active" && (
          <>
            <div className={styles.list_item}>Active Course 1 - Example Text</div>
            <div className={styles.list_item}>Active Course 2 - Example Text</div>
            <div className={styles.list_item}>Active Course 3 - Example Text</div>
            <div className={styles.list_item}>Active Course 4 - Example Text</div>
          </>
        )}
        {selectedButton === "Inactive" && (
          <>
            <div className={styles.list_item}>Inactive Course 1 - Example Text</div>
            <div className={styles.list_item}>Inactive Course 2 - Example Text</div>
            <div className={styles.list_item}>Inactive Course 3 - Example Text</div>
            <div className={styles.list_item}>Inactive Course 4 - Example Text</div>
          </>
        )}
        {selectedButton === "My Major" && (
          <>
            {error && <p className={styles.error}>Error: {error}</p>}

            <div className={styles.header_bar}>
              <span className={styles.column_header}>Course Name</span>
              <span className={styles.column_header}>Credits</span>
              <span className={styles.column_header}>Professor</span>
              <span className={styles.column_header}>Max Enrollment</span>
            </div>

            {majorCourses.length > 0 ? (
              majorCourses.map((course, index) => (
                <div
                  key={index}
                  className={styles.list_item}
                  onClick={() => handleCourseClick(course.course_id)}
                  style={{ cursor: "pointer" }}
                >
                  <span className={styles.course_name}>{course.course_name}</span>
                  <span className={styles.course_credits}>{course.num_credits}</span>
                  <span className={styles.teacher_name}>
                    {teacherNames[course.teacher_id] || course.teacher_id}
                  </span>
                  <span className={styles.max_enrollment}>{course.max_enrollment}</span>
                </div>
              ))
            ) : (
              <p>No courses for this major found.</p>
            )}
          </>
        )}
        {selectedButton === "Enrolled" && (
          <>
            {error && <p className={styles.error}>Error: {error}</p>}

            <div className={styles.header_bar}>
              <span className={styles.column_header}>Course Name</span>
              <span className={styles.column_header}>Credits</span>
              <span className={styles.column_header}>Professor</span>
              <span className={styles.column_header}>Max Enrollment</span>
            </div>

            {enrolledCourses.length > 0 ? (
              enrolledCourses.map((course, index) => (
                <div
                  key={index}
                  className={styles.list_item}
                  onClick={() => handleCourseClick(course.course_id)}
                  style={{ cursor: "pointer" }}
                >
                  <span className={styles.course_name}>{course.course_name}</span>
                  <span className={styles.course_credits}>{course.num_credits}</span>
                  <span className={styles.teacher_name}>
                    {teacherNames[course.teacher_id] || course.teacher_id}
                  </span>
                  <span className={styles.max_enrollment}>{course.max_enrollment}</span>
                </div>
              ))
            ) : (
              <p>No enrolled courses found.</p>
            )}
          </>
        )}
        {selectedButton === "Dropped" && (
          <>
            {error && <p className={styles.error}>Error: {error}</p>}

            <div className={styles.header_bar}>
              <span className={styles.column_header}>Course Name</span>
              <span className={styles.column_header}>Credits</span>
              <span className={styles.column_header}>Professor</span>
              <span className={styles.column_header}>Max Enrollment</span>
            </div>

            {droppedCourses.length > 0 ? (
              droppedCourses.map((course, index) => (
                <div
                  key={index}
                  className={styles.list_item}
                  onClick={() => handleCourseClick(course.course_id)}
                  style={{ cursor: "pointer" }}
                >
                  <span className={styles.course_name}>{course.course_name}</span>
                  <span className={styles.course_credits}>{course.num_credits}</span>
                  <span className={styles.teacher_name}>
                    {teacherNames[course.teacher_id] || course.teacher_id}
                  </span>
                  <span className={styles.max_enrollment}>{course.max_enrollment}</span>
                </div>
              ))
            ) : (
              <p>No dropped courses found.</p>
            )}
          </>
        )}
      </div>
    </div>
  );
}
