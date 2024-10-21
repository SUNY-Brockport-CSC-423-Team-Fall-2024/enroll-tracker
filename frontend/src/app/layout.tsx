"use client";

import localFont from "next/font/local";
import "./globals.css";
import { TokenRefreshProvider } from "./providers/token-refresh-provider";
import { AuthProvider } from "./providers/auth-provider";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <AuthProvider>
      <TokenRefreshProvider>
        <html lang="en">
          <body className={`${geistSans.variable} ${geistMono.variable}`}>{children}</body>
        </html>
      </TokenRefreshProvider>
    </AuthProvider>
  );
}
