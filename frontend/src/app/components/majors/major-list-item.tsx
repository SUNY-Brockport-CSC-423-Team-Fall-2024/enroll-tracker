'use client'

import { ButtonType, Major } from "@/app/lib/definitions";
import styles from "../styles.module.css";
import Button from "../button";

interface MajorItemItemProps {
  major: Major;
  studentsMajorID: number | undefined;
}

export default function MajorListItem({major, studentsMajorID}: MajorItemItemProps) {
  return (
    <div className={styles.majors_list_item_container}>
      <div className={styles.majors_list_item_content}>
        <p className={styles.majors_list_item_title}><strong>{major.name}</strong></p>
        <p><strong>Description:</strong> {major.description}</p>
      </div>
      <div className={styles.majors_list_item_action_button}>
        {studentsMajorID !== undefined && major.id === studentsMajorID && (
          <Button btnTitle="Your Major" btnType={ButtonType.PRIMARY} href={`/majors/${major.id}`} />
        )}
        {studentsMajorID !== undefined && major.id !== studentsMajorID && (
          <Button btnTitle="Learn More" btnType={ButtonType.SECONDARY} href={`/majors/${major.id}`} />
        )}
        {studentsMajorID === undefined && (
          <Button btnTitle="View Major" btnType={ButtonType.PRIMARY} href={`/majors/${major.id}`} />
        )}
      </div>
    </div>
  );
}
