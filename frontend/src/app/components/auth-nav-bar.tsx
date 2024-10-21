"use client";

import { links } from "../lib/data";
import styles from "./styles.module.css";
import Link from "next/link";
import { useAuth } from "@/app/providers/auth-provider";
import { useRouter } from "next/navigation";

interface IAuthNavbar {
  userRole: string | undefined;
}

const AuthNavbar: React.FC<IAuthNavbar> = ({ userRole }) => {
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
    <div className={styles.auth_navbar}>
      {links.map((link) => {
        if (link.allowedRoles.find((role) => role === userRole) === undefined) {
          return null;
        }
        return (
          <div className={styles.auth_navlink} key={link.name}>
            <Link key={link.name} href={link.href}>
              {link.name}
            </Link>
          </div>
        );
      })}
      <div className={styles.auth_navlink}>
        <button onClick={handleLogout}>Logout</button>
      </div>
    </div>
  );
};

export default AuthNavbar;
