import ChangePassword from "@/app/components/change-password";
import styles from "./styles.module.css";

export default function Settings() {
  return (
    <div className={styles.settings_root}>
      <h1>Settings</h1>
      <div className={styles.settings_change_password}>
        <ChangePassword />
      </div>
    </div>
  );
}
