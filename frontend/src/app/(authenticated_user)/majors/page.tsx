//Ammaad Denmark 11/13
//Abstract component that will represent all majors

import styles from "./styles.module.css";
import MajorListItem from "@/app/components/majors-listitem";

export default function Majors() {
  return (
    <div className={styles.majors_root}>
      <h1>Majors</h1>

      <MajorListItem
        majorName="Poultry science"
        description="The study and research of various forms of chicken"
        path="poultryscience.edu"

      ></MajorListItem>
    </div>
  );
}//End of majors function


//Ammaad Denmark, create list item for every single major
//Give dummy data to make sure its working,
//You'll have to create another route.
