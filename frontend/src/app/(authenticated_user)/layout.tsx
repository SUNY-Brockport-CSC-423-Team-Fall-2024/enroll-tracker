import AuthNavbar from "../components/auth-nav-bar";
import AuthHeader from "../components/auth-header";
import styles from "./styles.module.css";
import { cookies } from "next/headers";
import { JWTCLOCKTOLERANCE, JWTSIGNINGALGORITHM, getJWTPublicTokenPEMFormatted } from "../lib/data";
import * as jose from "jose";

export default async function Layout({ children }: { children: React.ReactNode }) {
  let accessToken = cookies().get("access_token");
  let userRole: string | undefined = undefined;

  if (accessToken !== undefined) {
    try {
      const publicKey = await jose.importSPKI(getJWTPublicTokenPEMFormatted(), JWTSIGNINGALGORITHM);
      const { payload } = await jose.jwtVerify(accessToken.value, publicKey, {
        clockTolerance: JWTCLOCKTOLERANCE,
      });
      if (typeof payload.role === "string") {
        userRole = payload.role;
      } else {
        throw new Error("no user role");
      }
    } catch (err) {
      console.error("can verify user role");
    }
  }

  return (
    <div className={styles.auth_root}>
      <div className={styles.auth_navbar}>
        <AuthNavbar userRole={userRole} />
      </div>
      <div className={styles.auth_header}>
        <AuthHeader userRole={userRole} />
      </div>
      <div className={styles.auth_content}>{children}</div>
    </div>
  );
}
