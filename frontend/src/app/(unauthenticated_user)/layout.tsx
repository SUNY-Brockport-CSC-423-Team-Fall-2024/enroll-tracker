import UnauthNavbar from "../components/unauth-nav-bar";
import styles from "./styles.module.css"

export default function Layout({children}: {children: React.ReactNode }) {
    return (
        <div className={styles.all_content}>
            <UnauthNavbar />
            {children}
        </div>
    )
}
