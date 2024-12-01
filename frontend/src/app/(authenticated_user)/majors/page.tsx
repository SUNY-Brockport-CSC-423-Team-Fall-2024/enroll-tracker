"use client";

import { useEffect, useState } from "react";
import styles from "./styles.module.css";
import { Major, Roles } from "@/app/lib/definitions";
import { addCourseToMajors, getMajors, getStudent } from "@/app/lib/client/data";
import { useAuth } from "@/app/providers/auth-provider";
import MajorListItem from "@/app/components/majors/major-list-item";

export default function Majors() {
  const [majors, setMajors] = useState<Major[]>([]);
  const [studentsMajorID, setStudentsMajorID] = useState<number | undefined>(undefined);
  const { username, userRole } = useAuth();

  const fetchMajors = async () => {
    try {
      const retrievedMajors: Major[] = await getMajors();
      setMajors(retrievedMajors);
    } catch (err) {
      setMajors([]);
    }
  };

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

  useEffect(() => {
    getStudentsMajorID();
    fetchMajors();
  }, [username]);

  return (
    <div className={styles.majors_root}>
      {userRole === Roles.STUDENT && (
        <div className={styles.majors_list}>
          {majors.map(major => (
            <MajorListItem major={major} studentsMajorID={studentsMajorID}/>
          ))}
        </div>
      )}
    </div>
  );
}
