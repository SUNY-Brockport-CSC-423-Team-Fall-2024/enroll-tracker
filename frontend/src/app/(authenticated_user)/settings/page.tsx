import ChangePassword from "@/app/components/change-password";
import styles from "./styles.module.css";
import SettingsTabBar from "@/app/components/settings-tab-bar";

export default function Settings() {
  return (
    <div className={styles.settings_root}>
      <h1>Settings</h1>
      <SettingsTabBar />
    </div>
  );
}
