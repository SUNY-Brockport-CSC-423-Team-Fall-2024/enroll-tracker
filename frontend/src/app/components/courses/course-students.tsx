"use client";

import { ITableHeader, ITableRow } from "@/app/lib/definitions";
import Table from "../table/table";
import { useEffect, useState } from "react";
import { getCoursesStudents } from "@/app/lib/client/data";

interface CourseStudentsProps {
  isEnrolled: boolean;
  courseID: number;
}

export default function CourseStudents({ isEnrolled, courseID }: CourseStudentsProps) {
  const [students, setStudents] = useState<ITableRow[]>([]);

  const tableHeaders: ITableHeader[] = [
    {
      title: "First Name",
    },
    {
      title: "Last Name",
    },
    {
      title: "Email",
    },
    {
      title: "Date Enrolled",
    },
  ];

  if (isEnrolled === false) {
    tableHeaders.push({ title: "Date Unenrolled" });
  }

  const getStudents = async () => {
    try {
      const students = await getCoursesStudents(courseID, isEnrolled);
      let tableRows: ITableRow[] = [];

      students.map((student) => {
        let tableRowContent = [
          student.first_name,
          student.last_name,
          student.email,
          new Date(student.enrolled_date).toLocaleDateString("en-US"),
        ];

        if (!isEnrolled) {
          tableRowContent.push(
            student.unenrolled_date !== null
              ? new Date(student.unenrolled_date).toLocaleDateString("en-US")
              : "n/a",
          );
        }

        tableRows.push({
          content: tableRowContent,
          clickable: false,
        });
      });

      setStudents(tableRows);
    } catch (err) {
      console.error(err);
      setStudents([]);
    }
  };

  useEffect(() => {
    getStudents();
  }, []);

  return (
    <>
      {students.length > 0 && <Table headers={tableHeaders} rows={students} />}
      {students.length === 0 && <p>No students.</p>}
    </>
  );
}
