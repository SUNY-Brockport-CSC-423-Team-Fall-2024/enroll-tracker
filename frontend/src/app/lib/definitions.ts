export enum Roles {
  STUDENT = "student",
  TEACHER = "teacher",
  ADMIN = "admin",
}

export const RouteRoles = {
  dashboard: {
    roles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
  courses: {
    roles: [Roles.STUDENT, Roles.TEACHER],
  },
  majors: {
    roles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
  settings: {
    roles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
  users: {
    roles: [Roles.ADMIN],
  },
};

export interface IAuthNavLinks {
  name: string;
  href: string;
  allowedRoles: string[];
}

export const headerLinks: IAuthNavLinks[] = [
  {
    name: "Dashboard",
    href: "/dashboard",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
  {
    name: "Courses",
    href: "/courses",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER],
  },
  {
    name: "Majors",
    href: "/majors",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
  {
    name: "Users",
    href: "/users",
    allowedRoles: [Roles.ADMIN],
  },
];

export const footerLinks: IAuthNavLinks[] = [
  {
    name: "Settings",
    href: "/settings",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
];

export const JWTCLOCKTOLERANCE = 10; //in seconds
export const JWTSIGNINGALGORITHM = "RS256";

export interface User {
    username: string,
    id: number,
    first_name: string,
    last_name: string,
    auth_id: number,
    phone_number: string,
    email: string,
    created_at: Date,
    updated_at: Date,
}

export interface Student extends User {
    major_id?: number,
}
export interface Teacher extends User {
    office: string,
}
export interface Administrator extends User {
    office: string,
}

export type Major = {
    id: number,
    name: string,
    description: string,
    status: string,
    last_updated: Date,
    created_at: Date,
}

export type Course = {
    id: number,
    name: string,
    description: string,
    teacher_id: number,
    max_enrollment: number,
    num_credits: number,
    status: string,
    last_updated: Date,
    created_at: Date,
}

export type StudentCourse = {
    course_id: number,
    course_name: string,
    course_description: string,
    teacher_id: number,
    max_enrollment: number,
    num_credits: number,
    status: string,
    last_updated: Date,
    created_at: Date,
    is_enrolled: boolean,
    enrolled_date: Date,
    unenrolled_date: Date | null,
}

export type TeacherCourse = {
    course_id: number,
    course_name: string,
    course_description: string,
    teacher_id: number,
    current_enrollment: number,
    max_enrollment: number,
    num_credits: number,
    status: string,
    last_updated: Date,
    created_at: Date,
}

export interface ITable {
    headers: ITableHeader[],
    rows: ITableRow[]
}

export interface ITableHeader {
    title: string,
}

export interface ITableRow {
    content: TableRowContent[],
    clickable: boolean,
    href?: string
}

export type TableRowContent = string | number;

export enum ButtonType {
    PRIMARY = "primary",
    SECONDARY = "secondary"
}
