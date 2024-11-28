import ChangePassword from "@/app/components/change-password";
import styles from "./styles.module.css";

export default function Settings() {
  return (
    <div className={styles.settings_root}>
      <div className={styles.settings_change_password}>
        <ChangePassword />
      </div>
    </div>
  );
}
