import { cookies } from "next/headers";
import { NextResponse } from "next/server";
import {
  getJWTPublicTokenPEMFormatted,
  JWTSIGNINGALGORITHM,
  JWTCLOCKTOLERANCE,
} from "@/app/lib/data";
import * as jose from "jose";

export async function GET() {
  let accessToken = cookies().get("access_token");
  if (accessToken !== undefined) {
    let pemFormattedKey = getJWTPublicTokenPEMFormatted();
    try {
      const publicKey = await jose.importSPKI(pemFormattedKey, JWTSIGNINGALGORITHM);
      const { payload } = await jose.jwtVerify(accessToken.value, publicKey, {
        clockTolerance: JWTCLOCKTOLERANCE,
      });

      return new NextResponse(JSON.stringify({ role: payload.role, username: payload.sub }), { status: 200 });
    } catch (err) {
      return new NextResponse(JSON.stringify({ role: undefined }), { status: 200 });
    }
  }

  return new NextResponse(JSON.stringify({ role: undefined }), { status: 400 });
}
