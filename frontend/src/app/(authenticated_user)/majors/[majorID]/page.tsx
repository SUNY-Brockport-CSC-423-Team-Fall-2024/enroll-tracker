'use client'

import { useEffect, useState } from "react";
import { Major, Roles } from "@/app/lib/definitions";
import { useAuthHeader } from "@/app/providers/auth-header-provider"
import { getMajor, getStudent } from "@/app/lib/client/data";
import { useParams, useRouter } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";
import styles from "../styles.module.css";
import MajorCourses from "@/app/components/majors/major-courses";

export default function MajorPage() {
  const router = useRouter();
  const { username, userRole } = useAuth();
  const { setPageTitle } = useAuthHeader();
  const [major, setMajor] = useState<Major | undefined>();
  const [studentsMajorID, setStudentsMajorID] = useState<number | undefined>(undefined);
  const params = useParams();
  const majorID = params.majorID;

  const getStudentsMajorID = async () => {
    try {
      if(username !== undefined) {
        const student = await getStudent(username);
        setStudentsMajorID(student.major_id);
      } else {
        setStudentsMajorID(undefined);
      }
    } catch (err) {
      console.error(err)
      setStudentsMajorID(undefined);
    }
  }

  const fetchMajor = async () => {
    try {
      if(majorID instanceof Array) {
        setMajor(undefined);
        return
      }
      const m = await getMajor(parseInt(majorID));
      setMajor(m);
      setPageTitle(m.name);
    } catch(err) {
      setMajor(undefined);
    } 
  }

  useEffect(() => {
    fetchMajor()
    getStudentsMajorID()
  }, [])

  return (
    <div className={styles.majors_root}>
      {userRole === Roles.STUDENT && (
        <>
          <header className={styles.header}>
            <button onClick={() => router.push("/majors")} className={styles.left_button}>
              Back
            </button>
          </header>
          <div>
            <hr />
          </div>
        </>
      )}
      <p><strong>Description:</strong> {major?.description}</p>
      {userRole === Roles.STUDENT && (
        <>
          {studentsMajorID === undefined && (
            <button className={styles.centered_button}>
              Declare Major
            </button>
          )}
        </>
      )}
      <p>Courses</p>
      <div className={styles.majors_courses_table }>
        <MajorCourses majorID={major?.id}/>
      </div>
    </div>
  )
}
