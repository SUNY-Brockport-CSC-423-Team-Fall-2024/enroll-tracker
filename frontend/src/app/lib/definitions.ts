import BooksIcon, { IconProps } from "../components/icons/books";
import CapIcon from "../components/icons/cap";
import DashboardIcon from "../components/icons/dashboard";
import GearIcon from "../components/icons/gear";
import PeopleIcon from "../components/icons/people";

export enum Roles {
  STUDENT = "student",
  TEACHER = "teacher",
  ADMIN = "admin",
}

export function isRole(value: unknown): value is Roles {
  return Object.values(Roles).includes(value as Roles);
}

export interface RouteRole {
  path: RegExp;
  roles: Roles[];
  subRoutes?: RouteRole[];
}

export const RouteRoles: RouteRole[] = [
  {
    path: /^\/dashboard/,
    roles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
  {
    path: /^\/courses/,
    roles: [Roles.STUDENT, Roles.TEACHER],
    subRoutes: [
      {
        path: /^\/courses.*\/edit/,
        roles: [Roles.TEACHER],
      },
      {
        path: /^\/courses.*\/add-course/,
        roles: [Roles.TEACHER],
      },
    ],
  },
  {
    path: /^\/majors/,
    roles: [Roles.STUDENT, Roles.ADMIN],
    subRoutes: [
      {
        path: /^\/majors.*\/edit/,
        roles: [Roles.ADMIN],
      },
      {
        path: /^\/majors.*\/add-major/,
        roles: [Roles.ADMIN],
      },
    ],
  },
  {
    path: /^\/users/,
    roles: [Roles.ADMIN],
  },
  {
    path: /^\/settings/,
    roles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
];

export interface IAuthNavLinks {
  name: string;
  href: string;
  allowedRoles: string[];
  icon: React.FC<IconProps>;
}

export const headerLinks: IAuthNavLinks[] = [
  {
    name: "Dashboard",
    href: "/dashboard",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
    icon: DashboardIcon
  },
  {
    name: "Courses",
    href: "/courses",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER],
    icon: BooksIcon
  },
  {
    name: "Majors",
    href: "/majors",
    allowedRoles: [Roles.STUDENT, Roles.ADMIN],
    icon: CapIcon
  },
  {
    name: "Users",
    href: "/users",
    allowedRoles: [Roles.ADMIN],
    icon: PeopleIcon
  },
];

export const footerLinks: IAuthNavLinks[] = [
  {
    name: "Settings",
    href: "/settings",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
    icon: GearIcon
  },
];

export const JWTCLOCKTOLERANCE = 10; //in seconds
export const JWTSIGNINGALGORITHM = "RS256";

export interface User {
  username: string;
  id: number;
  first_name: string;
  last_name: string;
  auth_id: number;
  phone_number: string;
  email: string;
  created_at: Date;
  updated_at: Date;
}

export interface Student extends User {
  major_id?: number;
}
export interface Teacher extends User {
  office: string;
}
export interface Administrator extends User {
  office: string;
}

export type Major = {
  id: number;
  name: string;
  description: string;
  status: string;
  last_updated: Date;
  created_at: Date;
};

export type Course = {
  id: number;
  name: string;
  description: string;
  teacher_id: number;
  current_enrollment: number;
  max_enrollment: number;
  num_credits: number;
  status: string;
  last_updated: Date;
  created_at: Date;
};

export type StudentCourse = {
  course_id: number;
  course_name: string;
  course_description: string;
  teacher_id: number;
  max_enrollment: number;
  num_credits: number;
  status: string;
  last_updated: Date;
  created_at: Date;
  is_enrolled: boolean;
  enrolled_date: Date;
  unenrolled_date: Date | null;
};

export type CoursesStudent = {
  student_username: string;
  student_id: number;
  first_name: string;
  last_name: string;
  auth_id: number;
  major_id: number;
  phone_number: string;
  email: string;
  created_at: Date;
  updated_at: Date;
  is_enrolled: boolean;
  enrolled_date: Date;
  unenrolled_date: Date | null;
};

export type TeacherCourse = {
  course_id: number;
  course_name: string;
  course_description: string;
  teacher_id: number;
  current_enrollment: number;
  max_enrollment: number;
  num_credits: number;
  status: string;
  last_updated: Date;
  created_at: Date;
};

export interface ITable {
  headers: ITableHeader[];
  rows: ITableRow[];
}

export interface ITableHeader {
  title: string;
}

export interface ITableRow {
  content: TableRowContent[];
  clickable: boolean;
  href?: string;
  callback?: () => void;
}

export type TableRowContent = string | number;

export enum ButtonType {
  PRIMARY = "primary",
  SECONDARY = "secondary",
}
