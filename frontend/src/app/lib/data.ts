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

export const links: IAuthNavLinks[] = [
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
  {
    name: "Settings",
    href: "/settings",
    allowedRoles: [Roles.STUDENT, Roles.TEACHER, Roles.ADMIN],
  },
];

export const JWTCLOCKTOLERANCE = 10; //in seconds
export const JWTSIGNINGALGORITHM = "RS256";

export const getSigningKeyFromEnv = (): string | undefined => {
  const signingKeyEncoded = process.env.ENROLL_TRACKER_RSA_PUBLIC_KEY;
  const signingKeyDecoded = Buffer.from(signingKeyEncoded ?? "", "base64").toString("utf8");
  const signingKey = signingKeyDecoded.match(/.{1,64}/g)?.join("\n");

  return signingKey;
};

export const getJWTPublicTokenPEMFormatted = (): string => {
  const signingKey = getSigningKeyFromEnv();
  const pemFormattedKey = `-----BEGIN PUBLIC KEY-----\n${signingKey}\n-----END PUBLIC KEY-----`;

  return pemFormattedKey;
};
