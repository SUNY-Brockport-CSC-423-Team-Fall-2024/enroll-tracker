import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import * as jose from "jose";
import {
  getJWTPublicTokenPEMFormatted,
  getSigningKeyFromEnv,
  JWTCLOCKTOLERANCE,
  JWTSIGNINGALGORITHM,
  RouteRoles,
} from "./app/lib/data";

async function handleTokenVerification(
  request: NextRequest,
  accessToken: string | undefined,
  route: string,
  roleKey: keyof typeof RouteRoles,
) {
  const signingKey = getSigningKeyFromEnv();
  const pemFormattedKey = getJWTPublicTokenPEMFormatted();

  if (!accessToken || !signingKey) {
    return NextResponse.redirect("/login");  
  }

  try {
    const publicKey = await jose.importSPKI(pemFormattedKey, JWTSIGNINGALGORITHM);
    const { payload } = await jose.jwtVerify(accessToken, publicKey, {
      clockTolerance: JWTCLOCKTOLERANCE,
    });

    // Check if user has the appropriate role for the route
    if (RouteRoles[roleKey].roles.find((role) => role === payload.role) === undefined) {
        return NextResponse.redirect(new URL("/dashboard", request.url));  // Simplified relative URL
    }

    // If token and role are valid, allow access
    return NextResponse.next();
  } catch (err) {
      request.cookies.delete("access_token");
      request.cookies.delete("refresh_token");
      request.cookies.delete("refresh_token_id");
      request.cookies.set("is_logged_in", "false");

    // Token invalid or other errors, redirect to login
    return NextResponse.redirect(new URL("/login", request.url));  // Simplified relative URL
  }
}

export async function middleware(request: NextRequest) {
  let accessToken = request.cookies.get("access_token")?.value;

  // Extract route and role key logic for easier reuse
  const routeMapping: Record<string, keyof typeof RouteRoles> = {
    "/dashboard": "dashboard",
    "/courses": "courses",
    "/majors": "majors",
    "/settings": "settings",
    "/users": "users",
  };

  // Find the matching route and corresponding role key
  for (const route in routeMapping) {
    if (request.nextUrl.pathname.startsWith(route)) {
      return handleTokenVerification(request, accessToken, route, routeMapping[route]);
    }
  }

  // If no route matched, proceed with the request (no verification required)
  return NextResponse.next();
}
