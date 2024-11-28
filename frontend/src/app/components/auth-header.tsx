'use client'

import { useState, useEffect } from "react";
import styles from "./styles.module.css"
import { getUser } from "../lib/client/data";
import { usePathname, useRouter } from "next/navigation";
import { useAuth } from "../providers/auth-provider";

export default function AuthHeader() {
    const [pageTitle, setPageTitle] = useState<string | null>(null)
    const [userFirstName, setUserFirstName] = useState<string | null>(null)
    const [userLastName, setUserLastName] = useState<string | null>(null)
    const [loading, setLoading] = useState<boolean>(true);
    const pathname = usePathname();
    const router = useRouter();
    const { username, userRole } = useAuth();

    const determinePath = async() => {
        try {
            const path = pathname.split('/')
            switch(path[1]) {
                case "courses":
                    setPageTitle("Courses")
                case "users":
                    setPageTitle("Manage Users")
                    break;
                case "majors":
                    setPageTitle("Majors")
                    break;
                case "settings":
                    setPageTitle("Settings")
                    break;
                default:
                    if (username === undefined || userRole === undefined) {
                        setPageTitle(`Welcome`)
                    } else {
                        const { first_name, last_name } = await getUser(username, userRole)
                        setPageTitle(`Hi, ${first_name}`)
                        setUserFirstName(first_name)
                        setUserLastName(last_name)
                    }
                }
            } catch(err) {
                console.error(err)
                setPageTitle(`Welcome`)
            } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        determinePath()
    }, [pathname, username])
    return (
      <header className={styles.auth_header}>
        <h1>{loading ? "Loading..." : pageTitle}</h1>
        <div className={styles.profile_button} onClick={() => router.push('/settings')}>{userFirstName} {userLastName}</div>
      </header>
    )
}
