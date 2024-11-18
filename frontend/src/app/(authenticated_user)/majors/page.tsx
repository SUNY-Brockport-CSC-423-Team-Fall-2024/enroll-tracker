//Ammaad Denmark 11/13
//Abstract component that will represent all majors

import styles from "./styles.module.css";
import MajorListItem from "@/app/components/majors-listitem";

export default function Majors() {
  return (
    <div className={styles.majors_root}>
      <h1>Majors</h1>

      <MajorListItem
        majorName="Computer Science"
        description="Learn more about computers and theory of computation"
        path="computerscience.edu"

      ></MajorListItem>
    </div>
  );
}//End of majors function


//Ammaad Denmark, create list item for every single major
//Give dummy data to make sure its working,
//You'll have to create another route.
