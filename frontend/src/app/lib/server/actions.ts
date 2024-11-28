import { cookies } from "next/headers";
import { JWTSIGNINGALGORITHM, JWTCLOCKTOLERANCE } from "../definitions";
import * as jose from "jose";

function getSigningKeyFromEnv(): string | undefined {
  const signingKeyEncoded = process.env.ENROLL_TRACKER_RSA_PUBLIC_KEY;
  const signingKeyDecoded = Buffer.from(signingKeyEncoded ?? "", "base64").toString("utf8");
  const signingKey = signingKeyDecoded.match(/.{1,64}/g)?.join("\n");

  return signingKey;
}

function getJWTPublicTokenPEMFormatted(): string {
  const signingKey = getSigningKeyFromEnv();
  const pemFormattedKey = `-----BEGIN PUBLIC KEY-----\n${signingKey}\n-----END PUBLIC KEY-----`;

  return pemFormattedKey;
}

async function userStuff(): Promise<{ [key: string]: any } | null> {
  "use server";
  let accessToken = (await cookies()).get("access_token");
  if (accessToken !== undefined) {
    let pemFormattedKey = getJWTPublicTokenPEMFormatted();
    try {
      const publicKey = await jose.importSPKI(pemFormattedKey, JWTSIGNINGALGORITHM);
      const { payload } = await jose.jwtVerify(accessToken.value, publicKey, {
        clockTolerance: JWTCLOCKTOLERANCE,
      });
      return payload;
    } catch (err) {
      console.error(err);
      return null;
    }
  }
  return null;
}

async function currentUser(): Promise<{ [key: string]: any } | null> {
  "use server";
  const user = await userStuff();
  console.log(user);
  if (user) {
    //Sub is username and user_id is user's id
    const { sub, user_id, role } = user;
    return { username: sub, user_id: user_id, role: role };
  }
  return null;
}

export { currentUser, getSigningKeyFromEnv, getJWTPublicTokenPEMFormatted };
