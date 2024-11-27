import AuthNavbar from "../components/auth-nav-bar";
import styles from "./styles.module.css";
import { cookies } from "next/headers";
import { getJWTPublicTokenPEMFormatted } from "../lib/server/actions";
import { JWTCLOCKTOLERANCE, JWTSIGNINGALGORITHM } from "@/app/lib/definitions";
import * as jose from "jose";
import AuthHeader from "../components/auth-header";

export default async function Layout({ children }: { children: React.ReactNode }) {
  let accessToken = cookies().get("access_token");
  let userRole: string = '';
  let username: string = '';

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
      if (typeof payload.sub === "string") {
        username = payload.sub;
      } else {
        throw new Error("no username");
      }
    } catch (err) {
      console.error("can verify user info");
    }
  }

  return (
    <div className={styles.auth_root}>
      <div className={styles.auth_navbar}>
        <AuthNavbar userRole={userRole} />
      </div>
      <div className={styles.auth_content}>
            <AuthHeader username={username} userRole={userRole}/>
            {children}
      </div>
    </div>
  );
}
