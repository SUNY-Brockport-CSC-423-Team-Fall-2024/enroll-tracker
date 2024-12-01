"use client";

import { useState, useEffect } from "react";
import styles from "./styles.module.css";
import { getUser } from "../lib/client/data";
import { usePathname, useRouter } from "next/navigation";
import { useAuth } from "../providers/auth-provider";
import { useAuthHeader } from "../providers/auth-header-provider";

export default function AuthHeader() {
  const { pageTitle, setPageTitle } = useAuthHeader();
  const [userFirstName, setUserFirstName] = useState<string | null>(null);
  const [userLastName, setUserLastName] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const pathname = usePathname();
  const router = useRouter();
  const { username, userRole } = useAuth();

  const determinePath = async () => {
    try {
      const path = pathname.split("/");
      switch (path[1]) {
        case "courses":
          if (path.length > 2) {
            switch (path[2]) {
              case "add-course":
                setPageTitle("Add Course");
                break;
            }
          } else {
            setPageTitle("Courses");
          }
          break;
        case "users":
          setPageTitle("Manage Users");
          break;
        case "majors":
          if (path.length > 2) {
            switch(path[2]) {
              case "add-major":
                setPageTitle("Add Major");
                break;
            }
          } else {
            setPageTitle("Majors");
          }
          break;
        case "settings":
          setPageTitle("Settings");
          break;
        default:
          if (username === undefined || userRole === undefined) {
            setUserFirstName(`User`);
            setPageTitle(`Welcome`);
          } else {
            const { first_name } = await getUser(username, userRole);
            setPageTitle(`Hi, ${first_name}`);
          }
      }
    } catch (err) {
      console.error(err);
      setPageTitle(`Welcome`);
    } finally {
      if (username === undefined || userRole === undefined) {
        setUserFirstName(`User`);
      } else {
        const { first_name, last_name } = await getUser(username, userRole);
        setUserFirstName(first_name);
        setUserLastName(last_name);
      }
      setLoading(false);
    }
  };

  useEffect(() => {
    determinePath();
  }, [pathname, username]);
  return (
    <header className={styles.auth_header}>
      <h1>{loading ? "Loading..." : pageTitle}</h1>
      <div className={styles.profile_button} onClick={() => router.push("/settings")}>
        {userFirstName} {userLastName}
      </div>
    </header>
  );
}
