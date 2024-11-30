import { TeacherCourse, Teacher, Student, User, Roles, Major, StudentCourse } from "../definitions";

export const passwordRegex = /^[a-zA-Z0-9!@#$%^&*()_+=\[\]{};':"\\|,.<>\/?~-]{8,30}$/;

export async function getUser(username: string, userType: string): Promise<User> {
  try {
    const userPath = resolveRouteFromRole(userType);
    const url = `http://localhost:8002/api/${userPath}/${username}`;
    const resp = await fetch(url, {
      method: "GET",
    });

    if (resp.ok) {
      const user: User = await resp.json();
      return user;
    } else {
      throw new Error("Request to get user information was not successful");
    }
  } catch (err) {
    console.error("Error retrieving user information");
    throw new Error("Error retrieving user information");
  }
}

function resolveRouteFromRole(role: string) {
  switch (role) {
    case Roles.STUDENT:
      return "students";
    case Roles.TEACHER:
      return "teachers";
    case Roles.ADMIN:
      return "administrators";
    default:
      throw new Error("Unable to resolve route from role");
  }
}

export async function getMajors(): Promise<Major[]> {
  try {
    const url = `http://localhost:8002/api/majors?limit=50`;
    const resp = await fetch(url, {
      method: "GET",
    });

    if (resp.ok) {
      const majors: Major[] = await resp.json();
      return majors;
    } else {
      throw new Error("Request to get major information was not successful");
    }
  } catch (err) {
    console.error("Error retrieving major information");
    throw new Error("Error retrieving major information");
  }
}

export async function getTeachers(query: string): Promise<Teacher[]> {
  try {
    const url = `http://localhost:8002/api/teachers?first_name=${query}`;
    const resp = await fetch(url, {
      method: "GET",
    });

    if (resp.ok) {
      const teacher: Teacher[] = await resp.json();
      return teacher;
    } else {
      throw new Error("Request to get teacher information was not successful");
    }
  } catch (err) {
    console.error("Error retrieving teacher information");
    throw new Error("Error retrieving teacher information");
  }
}

export async function getTeacher(username: string): Promise<Teacher> {
  try {
    const url = `http://localhost:8002/api/teachers/${username}`;
    const resp = await fetch(url, {
      method: "GET",
    });

    if (resp.ok) {
      const teacher: Teacher = await resp.json();
      return teacher;
    } else {
      throw new Error("Request to get teacher information was not successful");
    }
  } catch (err) {
    console.error("Error retrieving teacher information");
    throw new Error("Error retrieving teacher information");
  }
}

export async function getTeachersCourses(studentID: number): Promise<TeacherCourse[]> {
  try {
    const url = `http://localhost:8002/api/teachers/${studentID}/courses`;
    const resp = await fetch(url, {
      method: "GET",
    });
    const data: TeacherCourse[] = await resp.json();
    if (resp.ok) {
      return data;
    } else {
      return [];
    }
  } catch (err) {
    console.error(err);
    return [];
  }
}

export async function getStudentMajor(username: string): Promise<Major | null> {
  try {
    const url = `http://localhost:8002/api/students/${username}`;
    const resp = await fetch(url, {
      method: "GET",
    });

    const student: Student = await resp.json();
    console.log(student);

    if (resp.ok) {
      if (student.major_id !== undefined) {
        return await getMajor(student.major_id);
      }
      throw new Error("Unable to get student major. They have not declared one.");
    } else {
      throw new Error("Error retrieving student information");
    }
  } catch (err) {
    console.error(err);
    return null;
  }
}

export async function getMajor(majorID: number): Promise<Major> {
  try {
    const url = `http://localhost:8002/api/majors/${majorID}`;
    const resp = await fetch(url, {
      method: "GET",
    });

    if (resp.ok) {
      const major: Major = await resp.json();
      return major;
    } else {
      throw new Error("Request to get major information was not successful");
    }
  } catch (err) {
    console.error("Error retrieving major information");
    throw new Error("Error retrieving major information");
  }
}

export async function getStudentCourses(
  studentID: number,
  isEnrolled?: boolean,
): Promise<StudentCourse[]> {
  try {
    const url = `http://localhost:8002/api/students/${studentID}/courses${isEnrolled !== undefined ? `?isEnrolled=${isEnrolled}` : ``}`;
    const resp = await fetch(url, {
      method: "GET",
    });
    const data: StudentCourse[] = await resp.json();
    if (resp.ok) {
      return data;
    } else {
      return [];
    }
  } catch (err) {
    console.error(err);
    return [];
  }
}

export async function getStudent(username: string): Promise<Student> {
  try {
    const url = `http://localhost:8002/api/students/${username}`;
    const resp = await fetch(url, {
      method: "GET",
    });

    if (resp.ok) {
      const student: Student = await resp.json();
      return student;
    } else {
      throw new Error("Request to get student information was not successful");
    }
  } catch (err) {
    console.error("Error retrieving student information");
    throw new Error("Error retrieving student information");
  }
}

export interface AsyncResponse {
  success: boolean,
  errMessage?: string
}

export async function addCourseToMajors(courseID: number, majorIDs: number[]): Promise<AsyncResponse> {
  try {
      const url = `http://localhost:8002/api/courses/${courseID}/majors`
      const resp = await fetch(url, {
        method: "POST",
        body: JSON.stringify({
          majorIDs: majorIDs,
        })
      });

    if (resp.status === 201) {
        return { success: true }
    } else {
        return { success: false, errMessage: (await resp.text())}
    }
  } catch(err) {
    console.error(err)
    return { success: false, errMessage: err instanceof Error ? err.message : "Unknown error occurred" }
  }
}

export async function getCourseMajors(courseID: number)

export async function createCourse(courseName: string, courseDescription: string, courseTeacherID: number, maxEnrollment: number, numCredits: number, majorIDs: number[]): Promise<AsyncResponse> {
  try {
    const url = `http://localhost:8002/api/courses`;
    const resp = await fetch(url, {
      method: "POST",
      body: JSON.stringify({
        name: courseName,
        description: courseDescription,
        teacher_id: courseTeacherID,
        max_enrollment: maxEnrollment,
        num_credits: numCredits,
      })
    });
    if (resp.status === 200 || resp.status === 201) {
      const { id } = await resp.json();
      
      const majorResp = await addCourseToMajors(id, majorIDs)
      if (majorResp.success) {
        return { success: true }
      } else {
        return { success: false, errMessage: majorResp?.errMessage ?? "Error occured while adding course to majors" }
      }
    } else {
      return { success: false, errMessage: (await resp.text())}
    }
  } catch(err) {
    console.error(err)
    return { success: false, errMessage: err instanceof Error ? err.message : "Unknown error occurred" }
  }
}
