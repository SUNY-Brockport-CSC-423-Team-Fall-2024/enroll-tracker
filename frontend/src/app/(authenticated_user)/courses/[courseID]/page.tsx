"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";
import styles from "../styles.module.css";
import { useAuthHeader } from "@/app/providers/auth-header-provider";
import { Roles, Major } from "@/app/lib/definitions";
import { getCourseMajors } from "@/app/lib/client/data";
import PencilIcon from "@/app/components/icons/pencil";
import CourseStudents from "@/app/components/courses/course-students";

interface Course {
  id: number;
  name: string;
  description: string;
  num_credits: number;
  current_enrollment: number;
  max_enrollment: number;
  teacher_id: number;
}

interface Student {
  student_id: number;
  name: string;
  email: string;
  is_enrolled: boolean;
}

export default function CourseDetail() {
  const [course, setCourse] = useState<Course | null>(null);
  const [majors, setMajors] = useState<Major[]>([]);
  const [isEnrolled, setIsEnrolled] = useState<boolean | "dropped" | null>(null);
  const [selectedButton, setSelectedButton] = useState<string>("General");

  const { setPageTitle } = useAuthHeader();
  const { userRole, userID } = useAuth();
  const router = useRouter();
  const params = useParams();
  const courseID = params.courseID;

  const handleButtonClick = (button: string) => {
    setSelectedButton(button);
  };

  const fetchCourse = async () => {
    try {
      const response = await fetch(`http://localhost:8002/api/courses/${courseID}`);
      if (!response.ok) throw new Error("Failed to fetch course details");

      const data: Course = await response.json();
      const majors: Major[] = await getCourseMajors(data.id);
      setCourse(data);
      setPageTitle(data.name);
      setMajors(majors);
    } catch (error) {
      console.error(error);
    }
  };

  const checkEnrollment = async () => {
    try {
      const response = await fetch(`http://localhost:8002/api/enrollments/${courseID}/students`);
      if (response.ok) {
        const students: Student[] = await response.json();
        const studentEnrollment = students.find((student) => student.student_id === userID);

        if (studentEnrollment) {
          setIsEnrolled(studentEnrollment.is_enrolled ? true : "dropped");
        } else {
          setIsEnrolled(false); // User is not enrolled in the course
        }
      }
    } catch (error) {
      console.error(error);
    }
  };

  const buttons = ["General", "Enrolled Students", "Unenrolled Students"];

  useEffect(() => {
    if (courseID && userID) {
      fetchCourse();
      if (userRole === Roles.STUDENT) {
        checkEnrollment();
      }
    }
  }, [courseID, userID]);

  const handleEnrollment = async () => {
    const url = `http://localhost:8002/api/enrollments/${courseID}/${userID}`;
    const method = isEnrolled ? "DELETE" : "POST";

    try {
      const response = await fetch(url, { method });
      if (!response.ok) throw new Error(`Failed to update enrollment STATUS: ${response.status}`);
      setIsEnrolled((prevState) => (prevState === true ? "dropped" : true)); // Toggle enrollment status based on previous state
    } catch (error) {
      console.error(error);
    }
  };

  if (!course) return <p>Loading...</p>;

  return (
    <div className={styles.courses_root}>
      {userRole === Roles.STUDENT && (
        <>
          <header className={styles.header}>
            <button onClick={() => router.push("/courses")} className={styles.left_button}>
              Back
            </button>
          </header>
          <div>
            <hr />
          </div>
        </>
      )}
      {userRole === Roles.TEACHER && (
        <div className={styles.nav_container}>
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
          <div
            className={styles.edit_course_button_container}
            onClick={() => router.push(`/courses/${course.id}/edit`)}
          >
            <div className={styles.edit_course_button}>
              <PencilIcon fill="#FFFFFF" stroke="#FFFFFF" />
            </div>
          </div>
        </div>
      )}
      {(userRole === Roles.STUDENT || selectedButton === "General") && (
        <div>
          <p>
            <b>Description:</b> {course.description}
          </p>
          <p>
            <b>Credits:</b> {course.num_credits}
          </p>
          <p>
            <b>Current Enrollment:</b> {course.current_enrollment}
          </p>
          <p>
            <b>Max Enrollment:</b> {course.max_enrollment}
          </p>
          <div className={styles.course_majors_list}>
            <p>
              <b>Majors: </b>&nbsp;
            </p>
            <div className={styles.majors_list}>
              {majors.length > 0 &&
                majors.map((major, i) => {
                  if (i === majors.length - 1) {
                    return <p key={i}>{major.name}</p>;
                  } else {
                    return <p key={i}>{major.name},&nbsp;</p>;
                  }
                })}
            </div>
          </div>
        </div>
      )}
      {userRole === Roles.TEACHER && selectedButton === "Enrolled Students" && (
        <div className={styles.courses_students_table}>
          <CourseStudents isEnrolled={true} courseID={course.id} />
        </div>
      )}
      {userRole === Roles.TEACHER && selectedButton === "Unenrolled Students" && (
        <div className={styles.courses_students_table}>
          <CourseStudents isEnrolled={false} courseID={course.id} />
        </div>
      )}
      {userRole === Roles.STUDENT && (
        <>
          {isEnrolled === null ? (
            <p>Loading...</p>
          ) : isEnrolled === "dropped" ? (
            <p className={styles.red_center_text}>Dropped</p>
          ) : (
            <button onClick={handleEnrollment} className={styles.centered_button}>
              {isEnrolled ? "Drop" : "Enroll"}
            </button>
          )}
        </>
      )}
    </div>
  );
}
