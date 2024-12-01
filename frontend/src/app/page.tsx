import Button from "./components/button";
import { ButtonType } from "./lib/definitions";
import styles from "./page.module.css";
import Link from "next/link";

export default function Home() {
  return (
    <div className={styles.page}>
      <header className={styles.header}>
        <Link href="/">
          <h1 className={styles.landing_page_title}>EnrollmentTracker</h1>
        </Link>
      </header>
      <main className={styles.main}>
        <div className={styles.content_section}>
          <p className="subtitle">
            Streamline student enrollments and course management with ease.
          </p>
          <h2>Features</h2>
          <ul className="features-list">
            <li>
              <strong>For Students:</strong> Declare majors, enroll in courses, and drop courses.
            </li>
            <li>
              <strong>For Teachers:</strong> Create and manage courses effortlessly.
            </li>
            <li>
              <strong>For Administrators:</strong> Manage users and create majors.
            </li>
          </ul>
        </div>
        <div className={styles.card_container}>
          <Link href="/login" className={styles.card}>
            <Button btnTitle="Login" btnType={ButtonType.PRIMARY} />
          </Link>
        </div>
      </main>
      <footer className="footer">
        <p>&copy; 2024 Enrollment Tracker Team.</p>
      </footer>
    </div>
  );
}
