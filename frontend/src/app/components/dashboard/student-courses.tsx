"use client";

import Table from "@/app/components/table/table";
import { getStudent, getStudentCourses } from "@/app/lib/client/data";
import { ITableRow } from "@/app/lib/definitions";
import { useState, useEffect } from "react";
import { useAuth } from "@/app/providers/auth-provider";

export default function StudentCoursesTable({isEnrolled}:{isEnrolled?: boolean}) {
  const [courses, setCourses] = useState<ITableRow[]>([]);
  const { username } = useAuth();

  const courseHeaders = [
    {
      title: "Name",
    },
    {
      title: "# of Credits",
    },
    {
      title: "Date Enrolled",
    },
  ];

  const getCourses = async () => {
    try {
      if (username === undefined) return;

      const student = await getStudent(username);

      let studentCourses = await getStudentCourses(student.id);
    
      let courseRows: ITableRow[] = [];

      if(isEnrolled !== undefined) {
        studentCourses = studentCourses.filter(course => isEnrolled ? course.is_enrolled : !course.is_enrolled)
      }

      studentCourses.map((course) =>
          courseRows.push({
            content: [
              course.course_name,
              course.num_credits,
              new Date(course.enrolled_date).toLocaleDateString("en-US"),
            ],
            clickable: true,
            href: `/courses/${course.course_id}`,
          })
      );

      setCourses(courseRows);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    getCourses();
  }, []);

  return (
    <>
      {courses.length > 0 && <Table headers={courseHeaders} rows={courses} />}
      {courses.length === 0 && <p>Not enrolled in any courses.</p>}
    </>
  );
}
