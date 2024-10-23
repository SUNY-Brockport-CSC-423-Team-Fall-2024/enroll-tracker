import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
  const accessToken = cookies().get("access_token")?.value;

  if (accessToken === undefined) {
    return new NextResponse(JSON.stringify({ msg: "Error ocurred while changing password" }), {
      status: 400,
    });
  }

  try {
    const { current_password, new_password } = await req.json();

    const resp = await fetch("http://api:443/auth/change-password", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
      body: JSON.stringify({
        current_password: current_password,
        new_password: new_password,
      }),
    });

    if (resp.ok) {
      return new NextResponse(JSON.stringify({ msg: "Password successfully changed" }), {
        status: 200,
      });
    } else {
      throw new Error("Error ocurred while changing password");
    }
  } catch (err) {
    return new NextResponse(JSON.stringify({ msg: "Error ocurred while changing password" }), {
      status: 400,
    });
  }
}
