"use client";
import AuthNavbar from "../components/auth-nav-bar";
import AuthHeader from "../components/auth-header";
import styles from "./styles.module.css";
import AuthHeader from "../components/auth-header";
import { useAuth } from "../providers/auth-provider";
import { useEffect } from "react";
import { AuthHeaderProvider } from "../providers/auth-header-provider";

export default function Layout({ children }: { children: React.ReactNode }) {
  const { username, userRole } = useAuth();

  useEffect(() => {}, [username, userRole]);
  return (
    <div className={styles.auth_root}>
      <div className={styles.auth_navbar}>
        <AuthNavbar />
      </div>
      <div className={styles.auth_content}>
        <AuthHeaderProvider>
          <AuthHeader />
          {children}
        </AuthHeaderProvider>
      </div>
    </div>
  );
};
