"use client";

import { headerLinks, footerLinks } from "../lib/definitions";
import styles from "./styles.module.css";
import Link from "next/link";
import { useAuth } from "@/app/providers/auth-provider";
import { useRouter } from "next/navigation";

const AuthNavbar: React.FC = () => {
  const router = useRouter();

  const { setIsLoggedIn, setUsername, setUserID, setUserRole, userRole } = useAuth();

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
      router.push("/");
      setIsLoggedIn(false);
      setUsername(undefined)
      setUserID(undefined)
      setUserRole(undefined)
    }
  };

  return (
    <div className={styles.auth_navbar}>
      <h1>
        Enroll
        <br />
        Tracker
      </h1>
      <div className={styles.auth_navlinks_header}>
        {headerLinks.map((link) => {
          if (link.allowedRoles.find((role) => role === userRole) === undefined) {
            return null;
          }
          return (
            <div className={styles.auth_navlink} key={link.name}>
              <div className={styles.auth_navlink_icon}></div>
              <Link key={link.name} href={link.href}>
                {link.name}
              </Link>
            </div>
          );
        })}
      </div>
      <div className={styles.auth_navlinks_footer}>
        {footerLinks.map((link) => {
          if (link.allowedRoles.find((role) => role === userRole) === undefined) {
            return null;
          }
          return (
            <div className={styles.auth_navlink} key={link.name}>
              <div className={styles.auth_navlink_icon}></div>
              <Link key={link.name} href={link.href}>
                {link.name}
              </Link>
            </div>
          );
        })}
        <div className={styles.auth_navlink}>
          <button className={styles.auth_logout_button} onClick={handleLogout}>
            Logout
          </button>
        </div>
      </div>
    </div>
  );
};

export default AuthNavbar;
