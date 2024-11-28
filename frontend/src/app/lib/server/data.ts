import { currentUser } from "./actions";
import { Course, Student, Teacher, Major, StudentCourse, TeacherCourse } from "../definitions";

export async function getStudentCourses(studentID: number, isEnrolled?: boolean): Promise<StudentCourse[]> {
    try {
        const url = `http://api:443/api/students/${studentID}/courses${isEnrolled !== undefined ? `?isEnrolled=${isEnrolled}` : ``}`
        const resp = await fetch(url, {
            method: "GET"
        });
        const data: StudentCourse[] = await resp.json();
        if (resp.ok) {
            return data
        } else {
            return []
        }
    } catch (err) {
        console.error(err)
        return []
    }
}

export async function getTeachersCourses(studentID: number): Promise<TeacherCourse[]> {
    try {
        const url = `http://api:443/api/teachers/${studentID}/courses`
        const resp = await fetch(url, {
            method: "GET"
        });
        const data: TeacherCourse[] = await resp.json();
        if (resp.ok) {
            return data
        } else {
            return []
        }
    } catch (err) {
        console.error(err)
        return []
    }
}

export async function getCoursesStudents(courseID: number, isEnrolled: boolean | undefined): Promise<Course[]> {
    const url = `http://api:443/api/courses/${courseID}/students${isEnrolled !== undefined ? `?isEnrolled=${isEnrolled}` : ``}`
    
    try {
        const resp = await fetch(url, {
            method: "GET"
        });
        const data: Course[] = await resp.json();
        if (resp.ok) {
            return data
        } else {
            return []
        }
    } catch (err) {
        console.error(err)
        return []
    }
}

export async function getStudentMajor(): Promise<Major | null> {
    const user = await currentUser();
    if (!user) {
        return null
    }
    const { username } = user;
    try {
        const url = `http://api:443/api/students/${username}`
        const resp = await fetch(url, {
            method: "GET"
        });

        const student: Student = await resp.json();

        if (resp.ok) {
            if (student.major_id !== undefined) {
                return await getMajor(student.major_id)
            }
            throw new Error("Unable to get student major. They have not declared one.")
        } else {
            throw new Error("Error retrieving student information")
        }
    } catch (err) {
        console.error(err)
        return null
    }
}

export async function getMajor(majorID: number): Promise<Major> {
    try {
        const url = `http://api:443/api/majors/${majorID}`
        const resp = await fetch(url, {
            method: "GET"
        })

        if (resp.ok) {
            const major: Major = await resp.json();
            return major
        } else {
            throw new Error("Request to get major information was not successful")
        }
    } catch (err) {
       console.error("Error retrieving major information")
       throw new Error("Error retrieving major information")
    }
}

export async function getStudent(username: string): Promise<Student> {
    try {
        const url = `http://api:443/api/students/${username}`
        const resp = await fetch(url, {
            method: "GET"
        })

        if (resp.ok) {
            const student: Student = await resp.json();
            return student;
        } else {
            throw new Error("Request to get student information was not successful")
        }
    } catch(err) {
       console.error("Error retrieving student information")
       throw new Error("Error retrieving student information")
    }
}

export async function getTeacher(username: string): Promise<Teacher> {
    try {
        const url = `http://api:443/api/teachers/${username}`
        const resp = await fetch(url, {
            method: "GET"
        })

        if (resp.ok) {
            const teacher: Teacher = await resp.json();
            return teacher;
        } else {
            throw new Error("Request to get teacher information was not successful")
        }
    } catch(err) {
       console.error("Error retrieving teacher information")
       throw new Error("Error retrieving teacher information")
    }
}

