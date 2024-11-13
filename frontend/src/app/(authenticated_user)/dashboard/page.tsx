import styles from "./styles.module.css";
import PageSlider from "../../../app/components/page-slider";

export default function Dashboard() {
  return (
    <div className={styles.dashboard_root}>
      <PageSlider firstPage={"Test1"} secondPage={"Test2"}/>
    </div>
  );
}
