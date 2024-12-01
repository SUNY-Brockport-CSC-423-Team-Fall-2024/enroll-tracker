"use client";

import { useEffect, useState } from "react";
import { useAuthHeader } from "@/app/providers/auth-header-provider";
import { useParams, useRouter } from "next/navigation";
import { Course, Major } from "@/app/lib/definitions";
import { getCourseMajors, updateCourse, getMajors, deleteCourse } from "@/app/lib/client/data";
import styles from "./styles.module.css";
import componentStyles from "@/app/components/styles.module.css";
import { handleTextInputChange } from "@/app/lib/client/utils";
import { useAuth } from "@/app/providers/auth-provider";
import { AsyncResponse } from "@/app/lib/client/data";
import Select from "react-select";
import { MajorOptions } from "../../add-course/page";

export interface SelectOption {
  label: string;
  value: string;
}

export default function EditCourse() {
  const { userID } = useAuth();
  const { courseID } = useParams();
  const router = useRouter();

  const { setPageTitle } = useAuthHeader();
  const [course, setCourse] = useState<Course | null>(null);
  const [courseMajorsOptions, setCourseMajorsOptions] = useState<MajorOptions[]>([]);
  const [courseMajorIds, setCourseMajorIds] = useState<number[]>([]);
  const [origCourseMajorIds, setOrigCourseMajorIds] = useState<number[]>([]);
  const [availableMajors, setAvailableMajors] = useState<MajorOptions[]>([]);
  const [courseName, setCourseName] = useState<string>("");
  const [courseDescription, setCourseDescription] = useState<string>("");
  const [maxEnrollment, setMaxEnrollment] = useState<number | undefined>(undefined);
  const [numCredits, setNumCredits] = useState<number | undefined>(undefined);
  const [courseStatus, setCourseStatus] = useState<SelectOption | undefined>();
  const [resultSuccess, setResultSuccess] = useState<boolean | undefined>(undefined);
  const [resultMessage, setResultMessage] = useState<string | undefined>(undefined);

  const fetchAvailableMajors = async () => {
    const majors = await getMajors();
    let majorOptions: MajorOptions[] = [];
    majors.map((major) => majorOptions.push({ value: major.id, label: major.name }));
    setAvailableMajors(majorOptions);
  };

  const fetchCourse = async () => {
    try {
      const response = await fetch(`http://localhost:8002/api/courses/${courseID}`);
      if (!response.ok) throw new Error("Failed to fetch course details");

      const data: Course = await response.json();
      const majors: Major[] = await getCourseMajors(data.id);
      setCourse(data);
      setCourseStatus(
        data.status === "active"
          ? { label: "Active", value: "active" }
          : { label: "Inactive", value: "inactive" },
      );
      setCourseName(data.name);
      setCourseDescription(data.description);
      setMaxEnrollment(data.max_enrollment);
      setNumCredits(data.num_credits);
      setPageTitle(data.name);

      //Set orig course ids
      let majorOptions: MajorOptions[] = [];
      let majorIds: number[] = [];
      majors.map((major) => {
        majorOptions.push({ value: major.id, label: major.name });
        majorIds.push(major.id);
      });
      setCourseMajorsOptions(majorOptions);
      setCourseMajorIds(majorIds);
      setOrigCourseMajorIds(majorIds);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    setResultMessage(undefined);
    setResultSuccess(undefined);

    if (maxEnrollment === undefined) {
      setResultMessage("Max enrollment is not set.");
      setResultSuccess(false);
      return;
    }
    if (numCredits === undefined) {
      setResultMessage("Number of credits is not set.");
      setResultSuccess(false);
      return;
    }
    if (userID === undefined) {
      setResultMessage("Error occured updating course.");
      setResultSuccess(false);
      return;
    }
    if (courseStatus === undefined) {
      setResultMessage("Error occured updating course.");
      setResultSuccess(false);
      return;
    }
    if (courseMajorIds.length <= 0) {
      setResultMessage("Majors not selected for course.");
      setResultSuccess(false);
      return;
    }
    if (course === null) {
      setResultMessage("Unexpected error occured. Try again later.");
      setResultSuccess(false);
      return;
    }

    const resp: AsyncResponse = await updateCourse(
      course,
      origCourseMajorIds,
      courseMajorIds,
      course.id,
      courseName,
      courseDescription,
      maxEnrollment,
      numCredits,
      courseStatus.value,
    );

    if (resp.success) {
      setResultMessage("Successfully updated the course!");
      setResultSuccess(true);
    } else if (resp.errMessage) {
      setResultMessage("Unable to update course, error occured.");
      setResultSuccess(false);
    } else {
      setResultMessage("An error occured while updating the course.");
      setResultSuccess(false);
    }
  };

  useEffect(() => {
    fetchCourse();
    fetchAvailableMajors();
  }, []);

  return (
    <div className={styles.edit_course_container}>
      <div className={styles.edit_course_form_container}>
        <h2>Course Information</h2>
        <form onSubmit={handleSubmit} id="edit_course_form" className={styles.edit_course_form}>
          <div className={styles.edit_course_form_upper_fields}>
            <label htmlFor="course_name">Course Name</label>
            <input
              id="course_name"
              name="course_name"
              type="text"
              defaultValue={course !== null ? course.name : ""}
              onChange={(e) => handleTextInputChange(setCourseName, e)}
              required
            />
            <label htmlFor="course_description">Course Description</label>
            <textarea
              id="course_description"
              form="edit_course_form"
              name="course_description"
              defaultValue={course !== null ? course.description : ""}
              rows={5}
              maxLength={255}
              onChange={(e) => handleTextInputChange(setCourseDescription, e)}
              required
            />
            <label htmlFor="course_max_enrollment">Max Enrollment</label>
            <input
              id="course_max_enrollment"
              name="course_max_enrollment"
              type="number"
              defaultValue={course !== null ? course.max_enrollment : ""}
              min={0}
              max={99}
              onChange={(e) => handleTextInputChange(setMaxEnrollment, e, true)}
              required
            />
            <label htmlFor="course_num_credits">Number of Credits</label>
            <input
              id="course_num_credits"
              name="course_num_credits"
              type="number"
              defaultValue={course !== null ? course.num_credits : ""}
              min={1}
              max={6}
              onChange={(e) => handleTextInputChange(setNumCredits, e, true)}
              required
            />
            <label htmlFor="course_status">Status</label>
            <Select
              id="course_status"
              name="course_status"
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
              value={courseStatus}
              onChange={(e) => {
                if (e !== null) {
                  setCourseStatus({ label: e.label, value: e.value });
                }
              }}
              required
            />
            <label htmlFor="course_majors">Major</label>
            <Select
              id="course_majors"
              name="course_majors"
              styles={{
                valueContainer: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: "rgb(249, 233, 220)",
                }),
                indicatorsContainer: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: "rgb(249, 233, 220)",
                }),
                multiValueLabel: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: "rgb(25, 146, 162)",
                  color: "#ffffff",
                  padding: "5px",
                  borderRadius: "0px",
                }),
                multiValueRemove: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: "rgb(25, 146, 162)",
                  color: "#ffffff",
                  padding: "5px",
                  borderRadius: "0px",
                }),
              }}
              isMulti
              options={availableMajors}
              value={courseMajorsOptions}
              onChange={(e) => {
                let majorIds: number[] = [];
                let majorOptions: MajorOptions[] = [];
                e.map((option) => {
                  majorIds.push(option.value);
                  majorOptions.push({ label: option.label, value: option.value });
                });
                setCourseMajorIds(majorIds);
                setCourseMajorsOptions(majorOptions);
              }}
              required
            />
          </div>
          <div className={styles.edit_course_form_lower_fields_container}>
            <div className={styles.edit_course_form_lower_fields}>
              <input
                className={componentStyles.secondary_button}
                type="button"
                id="course_cancel_edit"
                name="course_cancel_edit"
                onClick={() => router.push("/courses")}
                value="Cancel"
              />
              <input
                className={componentStyles.primary_button}
                id="course_submit"
                name="course_submit"
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
