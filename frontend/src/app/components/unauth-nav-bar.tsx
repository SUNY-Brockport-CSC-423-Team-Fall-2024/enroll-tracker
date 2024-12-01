import styles from "./styles.module.css";
import Link from "next/link";

export default function UnauthNavbar() {
  return (
    <div className={styles.unauth_navbar}>
      <Link href="/">
        <h1>EnrollTracker</h1>
      </Link>
      <div></div>
    </div>
  );
}
