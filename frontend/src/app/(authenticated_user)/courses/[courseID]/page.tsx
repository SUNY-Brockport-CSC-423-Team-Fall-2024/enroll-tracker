"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";
import styles from "../styles.module.css";

interface Course {
  id: number;
  name: string;
  description: string;
  num_credits: number;
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
  const [isEnrolled, setIsEnrolled] = useState<boolean | "dropped" | null>(null);
  const { userID } = useAuth();
  const router = useRouter();
  const params = useParams();
  const courseID = params.courseID;

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        const response = await fetch(`http://localhost:8002/api/courses/${courseID}`);
        if (!response.ok) throw new Error("Failed to fetch course details");

        const data = await response.json();
        setCourse(data);
      } catch (error) {
        console.error(error);
      }
    };

    const checkEnrollment = async () => {
      try {
        const response = await fetch(`http://localhost:8002/api/enrollments/${courseID}/students`);
        if (response.ok) {
          const students: Student[] = await response.json();
          const studentEnrollment = students.find(student => student.student_id === userID);
          
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

    if (courseID && userID) {
      fetchCourse();
      checkEnrollment();
    }
  }, [courseID, userID]);

  const handleEnrollment = async () => {
    const url = `http://localhost:8002/api/enrollments/${courseID}/${userID}`;
    const method = isEnrolled ? "DELETE" : "POST";

    try {
      const response = await fetch(url, { method });
      if (!response.ok) throw new Error(`Failed to update enrollment STATUS: ${response.status}`);
      setIsEnrolled(prevState => prevState === true ? "dropped" : true); // Toggle enrollment status based on previous state
    } catch (error) {
      console.error(error);
    }
  };

  if (!course) return <p>Loading...</p>;

  return (
    <div className={styles.courses_root}>
      <header className={styles.header}>
        <h1>{course.name}</h1>
        <button onClick={() => router.push("/courses")} className={styles.right_button}>
          Back
        </button>
      </header>
      <div>
        <hr />
      </div>
      <div>
        <p><b>Description:</b> {course.description}</p>
        <p><b>Credits:</b> {course.num_credits}</p>
        <p><b>Max Enrollment:</b> {course.max_enrollment}</p>
      </div>
      {isEnrolled === null ? (
        <p>Loading...</p>
      ) : isEnrolled === "dropped" ? (
        <p className={styles.red_center_text}>Dropped</p>
      ) : (
        <button onClick={handleEnrollment} className={styles.centered_button}>
          {isEnrolled ? "Drop" : "Enroll"}
        </button>
      )}
    </div>
  );
}

