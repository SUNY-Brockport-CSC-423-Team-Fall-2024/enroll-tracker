import { ITableHeader, ITableRow } from "@/app/lib/definitions";
import Table from "../table/table";
import { useEffect, useState } from "react";
import { getMajorsCourses } from "@/app/lib/client/data";

interface MajorCoursesProps {
  majorID: number | undefined;
}

export default function MajorCourses({ majorID }: MajorCoursesProps) {
  const [courses, setCourses] = useState<ITableRow[]>([]);

  const tableHeaders: ITableHeader[] = [
    {
      title: "Course Title",
    },
    {
      title: "Description",
    },
  ];

  const getCourses = async () => {
    try {
      if (majorID === undefined) {
        setCourses([]);
        return;
      }

      let majorCourses = await getMajorsCourses(majorID);

      let courseRows: ITableRow[] = [];

      majorCourses.map((course) =>
        courseRows.push({
          content: [course.name, course.description],
          clickable: true,
          href: `/courses/${course.id}`,
        }),
      );

      setCourses(courseRows);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    getCourses();
  }, [majorID]);

  return (
    <>
      {courses.length > 0 && <Table headers={tableHeaders} rows={courses} />}
      {courses.length === 0 && <p>No courses.</p>}
    </>
  );
}
