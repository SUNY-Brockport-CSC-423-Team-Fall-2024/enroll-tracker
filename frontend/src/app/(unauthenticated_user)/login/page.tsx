"use client";

import { Dispatch, SetStateAction, useState } from "react";
import styles from "../styles.module.css";
import { useRouter } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";

export default function Login() {
  const router = useRouter();

  const { setIsLoggedIn, getUserRole, setUserRole } = useAuth();

  let [username, setUsername] = useState("");
  let [password, setPassword] = useState("");
  let [errorMessage, setErrorMessage] = useState("");

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const resp = await fetch(`/api/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: username,
          password: password,
        }),
      });

      if (resp.ok) {
        await fetch("/api/login-status", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            isLoggedIn: true,
          }),
        });
        setIsLoggedIn(true);
        setUserRole(await getUserRole());
        router.push("/dashboard");
      }
      // display error message for failed login
      else if (resp.status === 400) {
        setErrorMessage("Invalid login credentials. Please try again.");
      }
    } catch (err) {
      console.error(err);
    }
  };

  const handleTextInputChange = (
    reactHook: Dispatch<SetStateAction<string>>,
    e: React.FormEvent<HTMLInputElement>,
  ) => {
    reactHook(e.currentTarget.value);
  };

  return (
    <div className={styles.main_content}>
      <div className={styles.login_box}>
        <h2>Login</h2>
        <form onSubmit={handleLogin} className={styles.login_form}>
          <div className={styles.login_form_field}>
            <label htmlFor="username">Username</label>
            <input
              id="username_input"
              type="text"
              name="username"
              onChange={(e) => handleTextInputChange(setUsername, e)}
            />
          </div>
          <div className={styles.login_form_field}>
            <label htmlFor="password">Password</label>
            <input
              id="password_input"
              type="password"
              name="password"
              onChange={(e) => handleTextInputChange(setPassword, e)}
            />
          </div>

          {errorMessage && <div className={styles.error_message}>{errorMessage}</div>}

          <div className={styles.login_form_field}>
            <input id={styles.login_submit} type="submit" value="Login" />
          </div>
        </form>
      </div>
    </div>
  );
}
