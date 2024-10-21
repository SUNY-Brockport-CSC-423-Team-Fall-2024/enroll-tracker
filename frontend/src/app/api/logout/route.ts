import { cookies } from "next/headers";
import { NextResponse } from "next/server";

export async function POST() {
  try {
    const accessToken = cookies().get("access_token");
    const refreshToken = cookies().get("refresh_token");
    const refreshTokenID = cookies().get("refresh_token_id");
    //TODO: Find better way to address missing cookies
    const resp = await fetch("http://api:443/auth/logout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken?.value}`,
      },
      body: JSON.stringify({
        refresh_token: refreshToken?.value,
        refresh_token_id: refreshTokenID?.value,
      }),
    });
    cookies().delete("access_token");
    cookies().delete("refresh_token");
    cookies().delete("refresh_token_id");

    if (resp.ok) {
      return new NextResponse(JSON.stringify({ msg: "Logged out successfully" }), { status: 200 });
    } else {
      //TODO: Should this return a 200 or 400. Probably 400 but it's better to push a logout
      return new NextResponse(JSON.stringify({ msg: "Unable to logout" }), { status: 200 });
    }
  } catch (err) {
    return new NextResponse(JSON.stringify({ err }), { status: 200 });
  }
}
