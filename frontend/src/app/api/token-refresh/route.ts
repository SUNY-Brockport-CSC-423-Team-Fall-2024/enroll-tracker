import { cookies } from "next/headers";
import { NextResponse } from "next/server";

export async function GET() {
  console.log("called");
  try {
    const oldRefreshToken = cookies().get("refresh_token");
    const oldRefreshTokenID = cookies().get("refresh_token_id");
    //TODO: Find better way to address missing cookies
    const resp = await fetch("http://api:443/auth/token-refresh", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        refresh_token: oldRefreshToken?.value,
        refresh_token_id: oldRefreshTokenID?.value,
      }),
    });

    const data = await resp.json();
    const { access_token, refresh_token, refresh_token_id } = data;

    let response = new NextResponse(JSON.stringify({ msg: "Refreshed token" }), { status: 200 });

    response.cookies.set("access_token", access_token, {
      httpOnly: true,
      sameSite: "lax",
      secure: false,
      path: "/",
    });
    response.cookies.set("refresh_token", refresh_token, {
      httpOnly: true,
      sameSite: "lax",
      secure: false,
      path: "/",
    });
    response.cookies.set("refresh_token_id", refresh_token_id, {
      httpOnly: true,
      sameSite: "lax",
      secure: false,
      path: "/",
    });

    return response;
  } catch (err) {
    return new NextResponse(JSON.stringify({ msg: err }), { status: 400 });
  }
}
