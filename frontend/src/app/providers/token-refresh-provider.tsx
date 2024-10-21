"use client";

import React, { createContext, useContext, useEffect, useRef } from "react";
import { useAuth } from "./auth-provider";
import { usePathname } from "next/navigation";

const TokenRefreshContext = createContext(null);

export const useTokenRefresh = () => {
  return useContext(TokenRefreshContext);
};

export function TokenRefreshProvider({ children }: { children: React.ReactNode }) {
  const refreshTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const initalTimeoutRef = useRef<NodeJS.Timeout | null>(null);
  const { isLoggedIn, setIsLoggedIn, checkLoginStatus } = useAuth();
  const hasInitRef = useRef(false);
  const pathname = usePathname();

  const refreshToken = async () => {
    try {
      const resp = await fetch("/api/token-refresh", {
        method: "GET",
        credentials: "include",
      });

      if (!resp.ok) {
        throw new Error("Failed to refresh token");
      }
    } catch (err) {
      console.error("Couldn't refresh access token");
    }
  };

  useEffect(() => {
    async () => {
      setIsLoggedIn((await checkLoginStatus()) ?? false);
      if (isLoggedIn && !hasInitRef.current && pathname !== "/refresh-session") {
        hasInitRef.current = true;
        initalTimeoutRef.current = setTimeout(() => {
          //Call refresh token
          refreshToken();
          //Start continuous refresh
          refreshTimeoutRef.current = setInterval(
            () => {
              refreshToken();
            },
            1000 * 60 * 5,
          ); //Continuous refreshes every 5 min
        }, 0); // Initial refresh. On page load always get a new access token.
      }

      if (!isLoggedIn) {
        if (initalTimeoutRef.current) {
          clearTimeout(initalTimeoutRef.current);
        }

        if (refreshTimeoutRef.current) {
          clearInterval(refreshTimeoutRef.current);
        }
        hasInitRef.current = false;
      }
    };

    //Clean up on component unmount
    return () => {
      if (initalTimeoutRef.current) {
        clearTimeout(initalTimeoutRef.current);
      }

      if (refreshTimeoutRef.current) {
        clearInterval(refreshTimeoutRef.current);
      }
    };
  }, [isLoggedIn]);

  return <TokenRefreshContext.Provider value={null}>{children}</TokenRefreshContext.Provider>;
}
