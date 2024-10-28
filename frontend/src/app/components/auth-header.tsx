"use client";

import { headerLinks, footerLinks } from "../lib/data";
import styles from "./styles.module.css";
import Link from "next/link";
import { useAuth } from "@/app/providers/auth-provider";
import { useRouter } from "next/navigation";

interface IAuthHeader {
    userRole: string | undefined;
}

// This is checking that the user is logged in?
const AuthHeader: React.FC<IAuthHeader> = ({ userRole }) => {
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
  
    return (
      // Start building component html here
      // Page Title, Profile Box, User Name (And initials)
      <div className="authheader">
        <div className="pageTitle"></div>
        <div className="userBox">
            <div className="userInit"></div>
            <div className="userName"></div>
        </div>
      </div>
    );
  };

export default AuthHeader;