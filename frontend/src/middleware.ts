import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import * as jose from "jose";
import { getJWTPublicTokenPEMFormatted, getSigningKeyFromEnv } from "./app/lib/server/actions";
import { isRole, JWTCLOCKTOLERANCE, JWTSIGNINGALGORITHM, Roles, RouteRole, RouteRoles } from "./app/lib/definitions";

async function handleTokenVerification(
  request: NextRequest,
  accessToken: string | undefined,
  allowedRoles: Roles[],
) {
  const signingKey = getSigningKeyFromEnv();
  const pemFormattedKey = getJWTPublicTokenPEMFormatted();

  if (!accessToken || !signingKey) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  try {
    const publicKey = await jose.importSPKI(pemFormattedKey, JWTSIGNINGALGORITHM);
    const { payload } = await jose.jwtVerify(accessToken, publicKey, {
      clockTolerance: JWTCLOCKTOLERANCE,
    });

    const userRole = payload.role;

    // Validate role
    if (!isRole(userRole) || !allowedRoles.includes(userRole)) {
      return NextResponse.redirect(new URL("/dashboard", request.url));
    }

    // If token and role are valid, allow access
    return NextResponse.next();
  } catch (err) {
    request.cookies.delete("access_token");
    request.cookies.delete("refresh_token");
    request.cookies.delete("refresh_token_id");
    request.cookies.set("is_logged_in", "false");

    // Token invalid or other errors, redirect to login
    return NextResponse.redirect(new URL("/login", request.url));
  }
}

export async function middleware(request: NextRequest) {
  const accessToken = request.cookies.get("access_token")?.value;
  const path = request.nextUrl.pathname;

  const matchedRoute = matchRoute(path, RouteRoles);

  if (!matchedRoute) {
    return NextResponse.next();
  }

  return handleTokenVerification(request, accessToken, matchedRoute.roles);
}

const matchRoute = (path: string, routeRoles: RouteRole[]): RouteRole | undefined => {
  for (const route of routeRoles) {
    if (route.subRoutes) {
      const subRouteMatch = matchRoute(path, route.subRoutes);
      if (subRouteMatch) return subRouteMatch;
    }

    if (route.path.test(path)) {
      return route;
    }
  }

  return undefined;
};
