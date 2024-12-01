'use client'

import { useState, useEffect } from "react"
import { AsyncResponse, createCourse, getMajors } from "@/app/lib/client/data";
import styles from "./styles.module.css";
import componentStyles from "@/app/components/styles.module.css";
import { handleTextInputChange } from "@/app/lib/client/utils";
import { useRouter } from "next/navigation";
import { useAuth } from "@/app/providers/auth-provider";
import Select from "react-select";

export interface MajorOptions {
  value: number;
  label: string;
}

export default function Page() {
  const router = useRouter();
  const { userID } = useAuth();

  const [courseName, setCourseName] = useState<string>("");
  const [courseDescription, setCourseDescription] = useState<string>("");
  const [maxEnrollment, setMaxEnrollment] = useState<number | undefined>(undefined);
  const [numCredits, setNumCredits] = useState<number | undefined>(undefined);
  const [resultSuccess, setResultSuccess] = useState<boolean | undefined>(undefined);
  const [resultMessage, setResultMessage] = useState<string | undefined>(undefined);
  const [courseMajors, setCourseMajors] = useState<number[]>([]);
  const [availableMajors, setAvailableMajors] = useState<MajorOptions[]>([]);
  
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    setResultMessage(undefined);
    setResultSuccess(undefined);

    if(maxEnrollment === undefined) {
      setResultMessage("Max enrollment is not set.")
      setResultSuccess(false);
      return;
    }
    if(numCredits === undefined) {
      setResultMessage("Number of credits is not set.")
      setResultSuccess(false);
      return;
    }
    if(userID === undefined) {
      setResultMessage("Error occured creating course.")
      setResultSuccess(false);
      return;
    }

    if(courseMajors.length <= 0) {
      setResultMessage("Major not selected for course.")
      setResultSuccess(false);
      return;
    }

    const resp: AsyncResponse = await createCourse(courseName, courseDescription, userID, maxEnrollment, numCredits, courseMajors);

    if (resp.success) {
      setResultMessage("Successfully created the course!");
      setResultSuccess(true);
    } else if (resp.errMessage) {
      setResultMessage("Unable to create course, error occured.")
      setResultSuccess(false);
    } else {
      setResultMessage("An error occured while creating the course.")
      setResultSuccess(false);
    }
  }

  const fetchMajors = async () => {
    const majors = await getMajors()
    let majorOptions: MajorOptions[] = []
    majors.map(major => majorOptions.push({ value: major.id, label: major.name }))
    setAvailableMajors(majorOptions)
  }

  useEffect(() => {
    fetchMajors()
  }, [])

  return (
    <div className={styles.add_course_container}>
      <div className={styles.add_course_form_container}>
        <h2>Course Information</h2>
        <form onSubmit={handleSubmit} id="add_course_form" className={styles.add_course_form}>
          <div className={styles.add_course_form_upper_fields}>
            <label htmlFor="course_name">Course Name</label>
            <input id="course_name" name="course_name" type="text" onChange={(e) => handleTextInputChange(setCourseName, e)} required/>
            <label htmlFor="course_description">Course Description</label>
            <textarea id="course_description" form="add_course_form" name="course_description" rows={5} maxLength={255} onChange={(e) => handleTextInputChange(setCourseDescription, e)} required/>
            <label htmlFor="course_max_enrollment">Max Enrollment</label>
            <input id="course_max_enrollment" name="course_max_enrollment" type="number" min={0} max={99} onChange={(e) => handleTextInputChange(setMaxEnrollment, e, true)} required/>
            <label htmlFor="course_num_credits">Number of Credits</label>
            <input id="course_num_credits" name="course_num_credits" type="number" min={1} max={6} onChange={(e) => handleTextInputChange(setNumCredits, e, true)} required/>
            <label htmlFor="course_majors">Major</label>
            <Select 
              id="course_majors"
              name="course_majors"
              styles={{
                valueContainer: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: 'rgb(249, 233, 220)'
                }),
                indicatorsContainer: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: 'rgb(249, 233, 220)'
                }),
                multiValueLabel: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: 'rgb(25, 146, 162)',
                  color: '#ffffff',
                  padding: '5px',
                  borderRadius: '0px'
                }),
                multiValueRemove: (baseStyles) => ({
                  ...baseStyles,
                  backgroundColor: 'rgb(25, 146, 162)',
                  color: '#ffffff',
                  padding: '5px',
                  borderRadius: '0px'
                })
              }}
              isMulti
              options={availableMajors}
              onChange={(e) => {
                let majorIds: number[] = []
                e.map(option => majorIds.push(option.value))
                setCourseMajors(majorIds)
              }}
              required
            />
          </div>
          <div className={styles.add_course_form_lower_fields_container}>
            <div className={styles.add_course_form_lower_fields}>
              <input className={componentStyles.secondary_button} type="button" id="course_cancel_creation" name="course_cancel_creation" onClick={() => router.push('/courses')} value="Cancel" />
              <input className={componentStyles.primary_button} id="course_submit" name="course_submit" type="submit" value="Confirm"/>
            </div>
          </div>
        </form>
        <div className={styles.result_message_container}>
          {resultSuccess && <p className={styles.result_successful}>{resultMessage}</p>}
          {resultSuccess === false && <p className={styles.result_unsuccessful}>{resultMessage}</p>}
        </div>
      </div>
    </div>
  )
}

