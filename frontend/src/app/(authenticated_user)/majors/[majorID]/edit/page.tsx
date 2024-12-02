"use client";

import { useEffect, useState } from "react";
import { useAuthHeader } from "@/app/providers/auth-header-provider";
import { useParams, useRouter } from "next/navigation";
import { Major } from "@/app/lib/definitions";
import { updateMajor } from "@/app/lib/client/data";
import styles from "./styles.module.css";
import componentStyles from "@/app/components/styles.module.css";
import { handleTextInputChange } from "@/app/lib/client/utils";
import { useAuth } from "@/app/providers/auth-provider";
import { AsyncResponse } from "@/app/lib/client/data";
import Select from "react-select";

export interface SelectOption {
  label: string;
  value: string;
}

export default function EditMajor() {
  const { userID } = useAuth();
  const { majorID } = useParams();
  const router = useRouter();

  const { setPageTitle } = useAuthHeader();
  const [major, setMajor] = useState<Major | null>(null);
  const [majorName, setMajorName] = useState<string>("");
  const [majorDescription, setMajorDescription] = useState<string>("");
  const [majorStatus, setMajorStatus] = useState<SelectOption | undefined>();
  const [resultSuccess, setResultSuccess] = useState<boolean | undefined>(undefined);
  const [resultMessage, setResultMessage] = useState<string | undefined>(undefined);

  const fetchMajor = async () => {
    try {
      const response = await fetch(`http://localhost:8002/api/majors/${majorID}`);
      if (!response.ok) throw new Error("Failed to fetch major details");

      const data: Major = await response.json();
      setMajor(data);
      setMajorStatus(
        data.status === "active"
          ? { label: "Active", value: "active" }
          : { label: "Inactive", value: "inactive" },
      );
      setMajorName(data.name);
      setMajorDescription(data.description);
      setPageTitle(data.name);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    setResultMessage(undefined);
    setResultSuccess(undefined);

    if (userID === undefined) {
      setResultMessage("Error occured updating major.");
      setResultSuccess(false);
      return;
    }
    if (majorStatus === undefined) {
      setResultMessage("Error occured updating major.");
      setResultSuccess(false);
      return;
    }
    if (major === null) {
      setResultMessage("Unexpected error occured. Try again later.");
      setResultSuccess(false);
      return;
    }

    const resp: AsyncResponse = await updateMajor(
      major,
      major.id,
      majorName,
      majorDescription,
      majorStatus.value,
    );

    if (resp.success) {
      setResultMessage("Successfully updated the major!");
      setResultSuccess(true);
    } else if (resp.errMessage) {
      setResultMessage("Unable to update major, error occured.");
      setResultSuccess(false);
    } else {
      setResultMessage("An error occured while updating the major.");
      setResultSuccess(false);
    }
  };

  useEffect(() => {
    fetchMajor();
  }, []);

  return (
    <div className={styles.edit_major_container}>
      <div className={styles.edit_major_form_container}>
        <h2>Major Information</h2>
        <form onSubmit={handleSubmit} id="edit_major_form" className={styles.edit_major_form}>
          <div className={styles.edit_major_form_upper_fields}>
            <label htmlFor="major_name">Name</label>
            <input
              id="major_name"
              name="major_name"
              type="text"
              defaultValue={major !== null ? major.name : ""}
              onChange={(e) => handleTextInputChange(setMajorName, e)}
              required
              disabled
            />
            <label htmlFor="major_description">Description</label>
            <textarea
              id="major_description"
              form="edit_major_form"
              name="major_description"
              defaultValue={major !== null ? major.description : ""}
              rows={5}
              maxLength={255}
              onChange={(e) => handleTextInputChange(setMajorDescription, e)}
              required
            />
            <label htmlFor="major_status">Status</label>
            <Select
              id="major_status"
              name="major_status"
              styles={{
                valueContainer: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: "rgb(249, 233, 220)",
                }),
                indicatorsContainer: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: "rgb(249, 233, 220)",
                }),
              }}
              options={[
                {
                  label: "Active",
                  value: "active",
                },
                {
                  label: "Inactive",
                  value: "inactive",
                },
              ]}
              value={majorStatus}
              onChange={(e) => {
                if (e !== null) {
                  setMajorStatus({ label: e.label, value: e.value });
                }
              }}
              required
            />
          </div>
          <div className={styles.edit_major_form_lower_fields_container}>
            <div className={styles.edit_major_form_lower_fields}>
              <input
                className={componentStyles.secondary_button}
                type="button"
                id="major_cancel_edit"
                name="major_cancel_edit"
                onClick={() => {
                  if (major) {
                    router.push(`/majors/${major.id}`);
                  } else {
                    router.push(`/majors`);
                  }
                }}
                value="Cancel"
              />
              <input
                className={componentStyles.primary_button}
                id="major_submit"
                name="major_submit"
                type="submit"
                value="Save Changes"
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
