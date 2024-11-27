import styles from "./styles.module.css";
import { MajorName } from "@/app/components/dashboard/major-name";
import { StudentCoursesTable } from "@/app/components/dashboard/student-courses";
import { Suspense } from "react";

export default function Dashboard() {
  return (
    <div className={styles.dashboard_root}>
      <div className={styles.dashboard_student_major}>
        <h2>Your Major</h2>
        <Suspense fallback={<p>Loading major...</p>}>
          <MajorName />
        </Suspense>
      </div>
      <div className={styles.dashboard_student_courses_section}>
        <h2>Your Courses</h2>
        <Suspense fallback={<p>Loading courses...</p>}>
          <StudentCoursesTable /> 
        </Suspense>
      </div>
    </div>
  );
}
