"use client";
import { getStudentMajor } from "@/app/lib/client/data";
import { Roles } from "@/app/lib/definitions";
import { useAuth } from "@/app/providers/auth-provider";
import { useEffect, useState } from "react";

export function MajorName() {
  const { username, userRole } = useAuth();
  const [majorName, setMajorName] = useState<string | undefined>(undefined);

  const getMajorName = async () => {
    if (username && userRole === Roles.STUDENT) {
      const major = await getStudentMajor(username);
      if (major) {
        setMajorName(major.name);
      }
    }
  };

  useEffect(() => {
    getMajorName();
  }, [username, userRole]);
  return (
    <>
      {majorName && <p>{majorName}</p>}
      {!majorName && <p>No major.</p>}
    </>
  );
}
