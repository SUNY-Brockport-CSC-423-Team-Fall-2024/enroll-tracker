"use client";

import { Dispatch, SetStateAction, useState } from "react";
import { clsx } from "clsx";
import styles from "./styles.module.css";
import { passwordRegex } from "../lib/client/data";

export default function ChangePassword() {
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [resultMessage, setResultMessage] = useState("");
  const [resetSuccess, setResetSuccess] = useState<boolean | undefined>(undefined);

  const changePassword = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    //Reset result fields
    setResultMessage("");
    setResetSuccess(undefined);

    //Check passwords entered to see if they meet criteria
    if (!passwordRegex.test(currentPassword)) {
      setResultMessage("Current password doesn't meet password criteria");
      setResetSuccess(false);
      return;
    }
    if (!passwordRegex.test(newPassword)) {
      setResultMessage("New password doesn't meet password criteria");
      setResetSuccess(false);
      return;
    }

    //Make sure confirm password and new password match
    if (confirmPassword !== newPassword) {
      setResultMessage("New password and confirm password do not match.");
      setResetSuccess(false);
      return;
    }

    try {
      const resp = await fetch("/api/change-password", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          current_password: currentPassword,
          new_password: newPassword,
        }),
      });

      if (resp.ok) {
        setResultMessage("Successfully changed your password!");
        setResetSuccess(true);
      } else {
        setResultMessage("Error occurred when changing your password. Please try again");
        setResetSuccess(false);
      }
    } catch (err) {
      setResultMessage("Error occurred when changing your password. Please try again");
      setResetSuccess(false);
    }
  };

  const handleTextInputChange = (
    reactHook: Dispatch<SetStateAction<string>>,
    e: React.FormEvent<HTMLInputElement>,
  ) => {
    reactHook(e.currentTarget.value);
  };

  return (
    <div>
      <form onSubmit={changePassword} className={styles.change_password_form}>
        <div className={styles.change_password_form_field}>
          <label htmlFor="current_password">Current Password</label>
          <input
            type="password"
            id="current_password"
            name="current password"
            onChange={(e) => handleTextInputChange(setCurrentPassword, e)}
            minLength={8}
            maxLength={30}
          />
        </div>
        <div className={styles.change_password_form_field}>
          <label htmlFor="new_password">New Password</label>
          <input
            type="password"
            id="new_password"
            name="new password"
            onChange={(e) => handleTextInputChange(setNewPassword, e)}
            minLength={8}
            maxLength={30}
          />
        </div>
        <div className={styles.change_password_form_field}>
          <label htmlFor="confirm_password">Confirm Password</label>
          <input
            type="password"
            id="confirm_password"
            name="confirm password"
            onChange={(e) => handleTextInputChange(setConfirmPassword, e)}
            minLength={8}
            maxLength={30}
          />
        </div>
        <div className={styles.change_password_result_field}>
          {resultMessage && (
            <p
              className={clsx(
                { [styles.change_password_success_message]: resetSuccess },
                { [styles.change_password_error_message]: !resetSuccess },
              )}
            >
              {resultMessage}
            </p>
          )}
        </div>
        <div className={styles.change_password_form_field}>
          <input id={styles.change_password_submit} type="submit" value="Change Password" />
        </div>
      </form>
    </div>
  );
}
