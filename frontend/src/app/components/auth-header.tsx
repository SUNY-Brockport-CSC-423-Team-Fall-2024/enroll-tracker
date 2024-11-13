"use client";

import { headerLinks, footerLinks } from "../lib/data";
import styles from "./styles.module.css";
import Link from "next/link";
import { useAuth } from "@/app/providers/auth-provider";
import { useRouter } from "next/navigation";

interface IAuthHeader {
    pageTitle: string;
    userRole: string | undefined;
    userName: string;
    userInits: string;
}

// This is checking that the user is logged in?
const AuthHeader: React.FC<IAuthHeader> = ({ userRole, pageTitle, userName, userInits }) => {
    const router = useRouter();
    

    const { setIsLoggedIn } = useAuth();
  
    const handleLogout = async () => {
      const resp = await fetch("/api/logout", {
        method: "POST",
      });
  
      if (resp.ok) {
        await fetch("/api/login-status", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            isLoggedIn: false,
          }),
        });
        setIsLoggedIn(false);
        router.push("/");
      }
    };
  
    // Generate User initials
    userInits = generateInitials(userName);

    return (
      // Start building component html here
      // Page Title, Profile Box, User Name (And initials)
      <div className={styles.auth_header}>
            <div className={styles.auth_title}>
              <h1>{pageTitle}</h1>
            </div>
        <div className={styles.auth_userBox}>
            <div className={styles.auth_userInit}>
                <p>{userInits}</p>
            </div>
            <div className={styles.auth_userName}>
                <p>{userName}</p>
            </div>
        </div>
      </div>
    );
  };

  // Function to generate initials from username
export const generateInitials = (name: string | undefined): string => {
  if (!name || typeof name !== 'string') {
    return 'NA'; // or return a default value like 'NA' for Not Available
  }

  const nameParts = name.trim().split(" ");
  const initials = nameParts.length >= 2
    ? nameParts[0][0] + nameParts[1][0]
    : nameParts[0][0];
  return initials.toUpperCase();
};

export default AuthHeader;