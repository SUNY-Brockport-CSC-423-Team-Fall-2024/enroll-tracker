import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

export async function GET() {
  const isLoggedIn = cookies().get("is_logged_in");

  if (isLoggedIn === undefined) {
    return new NextResponse(JSON.stringify({ isLoggedIn: false }), { status: 200 });
  }

  return new NextResponse(JSON.stringify({ isLoggedIn: isLoggedIn.value }), { status: 200 });
}

export async function POST(req: NextRequest) {
  try {
    const data = await req.json();
    const { isLoggedIn } = data;
    let response = new NextResponse(JSON.stringify({ isLoggedIn: isLoggedIn }), { status: 200 });
    response.cookies.set("is_logged_in", isLoggedIn, {
      httpOnly: true,
      sameSite: "lax",
      secure: false,
      path: "/",
    });
    return response;
  } catch (err) {
    return new NextResponse(JSON.stringify({ isLoggedIn: false }), { status: 200 });
  }
}
