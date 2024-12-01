"use client";
import Button from "@/app/components/button";
import styles from "./styles.module.css";
import { MajorName } from "@/app/components/dashboard/major-name";
import StudentCoursesTable from "@/app/components/dashboard/student-courses";
import TeacherCoursesTable from "@/app/components/dashboard/teacher-courses";
import { ButtonType, Roles } from "@/app/lib/definitions";
import { useAuth } from "@/app/providers/auth-provider";

export default function Dashboard() {
  const { userRole } = useAuth();

  return (
    <div className={styles.dashboard_root}>
      {userRole === Roles.STUDENT && (
        <div className={styles.dashboard_student_major}>
          <h2>Your Major</h2>
          <MajorName />
        </div>
      )}
      {(userRole === Roles.STUDENT || userRole === Roles.TEACHER) && (
        <div className={styles.dashboard_courses_section}>
          <h2>Your Courses</h2>
          {userRole === Roles.STUDENT && (
            <div className={styles.dashboard_student_courses_table}>
              <StudentCoursesTable isEnrolled={true} />
            </div>
          )}
          {userRole === Roles.TEACHER && (
            <>
              <div className={styles.dashboard_teacher_courses_table}>
                <TeacherCoursesTable isActive={true} />
              </div>
              <div className={styles.dashboard_action_button_container}>
                <div className={styles.dashboard_action_button}>
                  <Button
                    btnTitle="Manage Courses"
                    href="/courses"
                    btnType={ButtonType.SECONDARY}
                  />
                </div>
              </div>
            </>
          )}
        </div>
      )}
      {userRole === Roles.ADMIN && (
        <>
          <h2>Users</h2>
          <div className={styles.dashboard_action_button_container}>
            <div className={styles.dashboard_action_button}>
              <Button btnTitle="Manage Users" href="/users" btnType={ButtonType.SECONDARY} />
            </div>
          </div>
          <h2>Courses</h2>
          <div className={styles.dashboard_action_button_container}>
            <div className={styles.dashboard_action_button}>
              <Button btnTitle="Manage Majors" href="/majors" btnType={ButtonType.SECONDARY} />
            </div>
          </div>
        </>
      )}
    </div>
  );
}
