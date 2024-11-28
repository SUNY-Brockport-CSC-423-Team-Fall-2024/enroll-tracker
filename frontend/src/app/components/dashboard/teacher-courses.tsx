"use client";

import Table from "@/app/components/table/table";
import { getTeacher, getTeachersCourses } from "@/app/lib/client/data";
import { ITableRow } from "@/app/lib/definitions";
import { useAuth } from "@/app/providers/auth-provider";
import { useEffect, useState } from "react";

export default function TeacherCoursesTable() {
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

      const teachersCourses = await getTeachersCourses(teacher.id);

      teachersCourses.map((course) =>
        setCourses([
          ...courses,
          {
            content: [
              course.course_name,
              `${course.current_enrollment}/${course.max_enrollment}`,
              new Date(course.last_updated).toLocaleDateString("en-US"),
            ],
            clickable: true,
            href: `/courses/${course.course_id}`,
          },
        ]),
      );
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
      {courses.length === 0 && <p>Not teaching any courses.</p>}
    </>
  );
}
