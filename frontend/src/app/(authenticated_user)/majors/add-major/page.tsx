"use client";

import { useState } from "react";
import { AsyncResponse, createMajor } from "@/app/lib/client/data";
import styles from "./styles.module.css";
import componentStyles from "@/app/components/styles.module.css";
import { handleTextInputChange } from "@/app/lib/client/utils";
import { useRouter } from "next/navigation";

export default function Page() {
  const router = useRouter();

  const [majorName, setCourseName] = useState<string>("");
  const [majorDescription, setCourseDescription] = useState<string>("");
  const [resultSuccess, setResultSuccess] = useState<boolean | undefined>(undefined);
  const [resultMessage, setResultMessage] = useState<string | undefined>(undefined);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    setResultMessage(undefined);
    setResultSuccess(undefined);

    const resp: AsyncResponse = await createMajor(majorName, majorDescription);

    if (resp.success) {
      setResultMessage("Successfully created the major!");
      setResultSuccess(true);
    } else if (resp.errMessage) {
      setResultMessage("Unable to create major, error occured.");
      setResultSuccess(false);
    } else {
      setResultMessage("An error occured while creating the major.");
      setResultSuccess(false);
    }
  };

  return (
    <div className={styles.add_major_container}>
      <div className={styles.add_major_form_container}>
        <h2>Major Information</h2>
        <form onSubmit={handleSubmit} id="add_major_form" className={styles.add_major_form}>
          <div className={styles.add_major_form_upper_fields}>
            <label htmlFor="major_name">Name</label>
            <input
              id="major_name"
              name="major_name"
              type="text"
              onChange={(e) => handleTextInputChange(setCourseName, e)}
              required
            />
            <label htmlFor="major_description">Description</label>
            <textarea
              id="major_description"
              form="add_major_form"
              name="major_description"
              rows={5}
              maxLength={255}
              onChange={(e) => handleTextInputChange(setCourseDescription, e)}
              required
            />
          </div>
          <div className={styles.add_major_form_lower_fields_container}>
            <div className={styles.add_major_form_lower_fields}>
              <input
                className={componentStyles.secondary_button}
                type="button"
                id="major_cancel_creation"
                name="major_cancel_creation"
                onClick={() => router.push("/majors")}
                value="Cancel"
              />
              <input
                className={componentStyles.primary_button}
                id="major_submit"
                name="major_submit"
                type="submit"
                value="Confirm"
              />
            </div>
          </div>
        </form>
        <div className={styles.result_message_container}>
          {resultSuccess && <p className={styles.result_successful}>{resultMessage}</p>}
          {resultSuccess === false && <p className={styles.result_unsuccessful}>{resultMessage}</p>}
        </div>
      </div>
    </div>
  );
}
