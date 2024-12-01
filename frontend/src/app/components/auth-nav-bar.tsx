"use client";

import { headerLinks, footerLinks } from "../lib/definitions";
import styles from "./styles.module.css";
import Link from "next/link";
import { useAuth } from "@/app/providers/auth-provider";
import { usePathname, useRouter } from "next/navigation";
import clsx from "clsx";

const AuthNavbar: React.FC = () => {
  const router = useRouter();
  const pathname = usePathname();

  const { setIsLoggedIn, setUsername, setUserID, setAuthID, setUserRole, userRole } = useAuth();

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
      setUsername(undefined);
      setUserID(undefined);
      setAuthID(undefined);
      setUserRole(undefined);
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
          const IconComponent = link.icon;
          return (
            <div
              className={clsx(
                styles.auth_navlink,
                pathname === link.href && styles.auth_navlink_selected,
              )}
              key={link.name}
            >
              <div className={styles.auth_navlink_icon}>
                <IconComponent stroke={link.name === "Dashboard" ? "#FFFFFF" : "none"} />
              </div>
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
          const IconComponent = link.icon;
          return (
            <div
              className={clsx(
                styles.auth_navlink,
                pathname === link.href && styles.auth_navlink_selected,
              )}
              key={link.name}
            >
              <div className={styles.auth_navlink_icon}>
                <IconComponent fill="none" stroke="#FFFFFF" />
              </div>
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
