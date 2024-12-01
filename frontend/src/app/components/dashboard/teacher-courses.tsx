"use client";

import Table from "@/app/components/table/table";
import { getTeacher, getTeachersCourses } from "@/app/lib/client/data";
import { ITableRow } from "@/app/lib/definitions";
import { useAuth } from "@/app/providers/auth-provider";
import { useEffect, useState } from "react";

export default function TeacherCoursesTable({ isActive }: { isActive?: boolean }) {
  const { username } = useAuth();
  const [courses, setCourses] = useState<ITableRow[]>([]);

  const courseHeaders = [
    {
      title: "Course Title",
    },
    {
      title: "Enrolled/Max",
    },
    {
      title: "Last Updated",
    },
  ];

  const getCourses = async () => {
    try {
      if (username === undefined) return;
      const teacher = await getTeacher(username);
      let teachersCourses = await getTeachersCourses(teacher.id);

      let tableRows: ITableRow[] = [];

      //Filter courses if prop passed
      if (isActive !== undefined) {
        teachersCourses = teachersCourses.filter((course) =>
          isActive ? course.status === "active" : course.status === "inactive",
        );
      }

      //Create table rows
      teachersCourses.map((course) =>
        tableRows.push({
          content: [
            course.course_name,
            `${course.current_enrollment}/${course.max_enrollment}`,
            new Date(course.last_updated).toLocaleDateString("en-US"),
          ],
          clickable: true,
          href: `/courses/${course.course_id}`,
        }),
      );
      setCourses(tableRows);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    getCourses();
  }, [isActive]);

  return (
    <>
      {courses.length > 0 && <Table headers={courseHeaders} rows={courses} />}
      {courses.length === 0 && isActive === undefined && <p>Not teaching any courses.</p>}
      {courses.length === 0 && isActive === true && <p>No active courses.</p>}
      {courses.length === 0 && isActive === false && <p>No inactive courses.</p>}
    </>
  );
}
