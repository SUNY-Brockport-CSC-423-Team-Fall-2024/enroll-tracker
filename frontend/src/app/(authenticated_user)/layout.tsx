'use client'
import AuthNavbar from "../components/auth-nav-bar";
import styles from "./styles.module.css";
import AuthHeader from "../components/auth-header";
import { useAuth } from "../providers/auth-provider";
import { useEffect } from "react";

export default function Layout({ children }: { children: React.ReactNode }) {
  const { username, userRole } = useAuth();

  useEffect(() => {
    
  }, [username, userRole])
  return (
    <div className={styles.auth_root}>
      <div className={styles.auth_navbar}>
        <AuthNavbar />
      </div>
      <div className={styles.auth_content}>
            <AuthHeader />
            {children}
      </div>
    </div>
  );
}
