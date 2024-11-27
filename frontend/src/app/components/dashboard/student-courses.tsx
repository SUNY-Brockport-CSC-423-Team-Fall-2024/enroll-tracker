import Table from "@/app/components/table/table";
import { currentUser } from "@/app/lib/server/actions";
import { getStudent, getStudentCourses } from "@/app/lib/server/data";
import { ITableRow, Student } from "@/app/lib/definitions";

export async function StudentCoursesTable() {
    let student: Student;
    let courses: ITableRow[] = [];

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
    ]

    try {
        const userData = await currentUser();
        if(!userData) {
            throw new Error("Unable to get student username")
        }
        student = await getStudent(userData.username);

        const studentCourses = await getStudentCourses(student.id, true);

        studentCourses.map((course) => courses.push({ content: [course.course_name, course.num_credits, new Date(course.enrolled_date).toLocaleDateString('en-US')], clickable: true, href: `/courses/${course.course_id}`}));
        
        return (
            <>
                {courses.length > 0 && <Table headers={courseHeaders} rows={courses} />}
                {courses.length === 0 && <p>Not enrolled in any courses.</p>}
            </>
        )
    } catch(err) {
        console.error(err)
        return (<p>Unable to load student courses.</p>)
    }
}
