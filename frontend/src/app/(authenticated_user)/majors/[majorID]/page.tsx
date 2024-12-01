"use client";

import { useEffect, useState } from "react";
import { Major, Roles } from "@/app/lib/definitions";
import { useAuthHeader } from "@/app/providers/auth-header-provider";
import { declareMajor, getMajor, getStudent } from "@/app/lib/client/data";
import { useParams, useRouter } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";
import styles from "../styles.module.css";
import MajorCourses from "@/app/components/majors/major-courses";
import PencilIcon from "@/app/components/icons/pencil";

export default function MajorPage() {
  const router = useRouter();
  const { username, userRole, userID } = useAuth();
  const { setPageTitle } = useAuthHeader();
  const [major, setMajor] = useState<Major | undefined>();
  const [resultSuccess, setResultSuccess] = useState<boolean | undefined>(undefined);
  const [resultMessage, setResultMessage] = useState<string | undefined>(undefined);
  const [studentsMajorID, setStudentsMajorID] = useState<number | undefined>(undefined);
  const params = useParams();
  const majorID = params.majorID;

  const getStudentsMajorID = async () => {
    try {
      if (username !== undefined) {
        const student = await getStudent(username);
        setStudentsMajorID(student.major_id !== null ? student.major_id : undefined);
      } else {
        setStudentsMajorID(undefined);
      }
    } catch (err) {
      console.error(err);
      setStudentsMajorID(undefined);
    }
  };

  const declareStudentMajor = async () => {
    try {
      if (major?.id === undefined || userID === undefined) {
        setResultMessage("Error occured while declaring for major.");
        setResultSuccess(false);
        return;
      }
      const resp = await declareMajor(userID, major.id);

      if (resp.success) {
        setStudentsMajorID(major.id);
      } else {
        setResultMessage("Error occured while declaring for major.");
        setResultSuccess(false);
      }
    } catch (err) {
      setResultMessage("Error occured while declaring for major.");
      setResultSuccess(false);
    }
  };

  const fetchMajor = async () => {
    try {
      if (majorID instanceof Array) {
        setMajor(undefined);
        return;
      }
      const m = await getMajor(parseInt(majorID));
      setMajor(m);
      setPageTitle(m.name);
    } catch (err) {
      setMajor(undefined);
    }
  };

  useEffect(() => {
    fetchMajor();
    if (userRole === Roles.STUDENT) {
      getStudentsMajorID();
    }
  }, []);

  return (
    <div className={styles.majors_root}>
      {userRole === Roles.ADMIN && (
        <div className={styles.nav_container}>
          <div
            className={styles.edit_major_button_container}
            onClick={() => {
              if (major) {
                router.push(`/majors/${major.id}/edit`);
              }
            }}
          >
            <div className={styles.edit_major_button}>
              <PencilIcon />
            </div>
          </div>
        </div>
      )}
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
      <p>
        <strong>Description:</strong> {major?.description}
      </p>
      {userRole === Roles.STUDENT && (
        <>
          {studentsMajorID === undefined && (
            <button className={styles.centered_button} onClick={() => declareStudentMajor()}>
              Declare Major
            </button>
          )}
          {studentsMajorID !== undefined && studentsMajorID === major?.id && (
            <p>You are in this major.</p>
          )}
        </>
      )}
      <div className={styles.majors_courses_container}>
        <h2>Courses</h2>
        <div className={styles.majors_courses_table}>
          <MajorCourses majorID={major?.id} />
        </div>
      </div>
    </div>
  );
}
