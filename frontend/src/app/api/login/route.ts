import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
  try {
    const data = await req.json();

    const { username, password } = data;

    const loginResp = await fetch("http://api:443/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    });

    const loginRespData = await loginResp.json();
    const { access_token, refresh_token, refresh_token_id } = loginRespData;

    let response = new NextResponse("", { status: 200 });

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
    return new NextResponse(JSON.stringify(err), { status: 400 });
  }
}
