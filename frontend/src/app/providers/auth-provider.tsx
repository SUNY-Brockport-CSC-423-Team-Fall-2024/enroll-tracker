import { createContext, useContext, useEffect, useState } from "react";

interface AuthContextProps {
  isLoggedIn: boolean;
  setIsLoggedIn: (value: boolean) => void;
  userRole: string | undefined;
  setUserRole: (role: string | undefined) => void;
  username: string | undefined;
  setUsername: (username: string | undefined) => void;
    getUserStuff: () => Promise<{[key:string]:any} | undefined>;
  checkLoginStatus: () => Promise<boolean>;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used with auth provider");
  }
  return context;
};

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [userRole, setUserRole] = useState<string | undefined>(undefined);
  const [username, setUsername] = useState<string | undefined>(undefined);

  const refreshToken = async () => {
    const resp = await fetch("/api/token-refresh", {
      method: "GET",
    });
    return resp.ok;
  };
  const checkLoginStatus = async (): Promise<boolean> => {
    const resp = await fetch("/api/login-status", {
      method: "GET",
    });
    const { isLoggedIn } = await resp.json();
    return isLoggedIn;
  };
  const getUserStuff = async (): Promise<{[key:string]: any} | undefined> => {
    const resp = await fetch("/api/user-stuff", {
      method: "GET",
    });
    const { role, username } = await resp.json();
    return { "role": role, "username": username };
  };

  useEffect(() => {
    const initializeAuth = async () => {
      try {
        // First, refresh the token
        const refreshSuccess = await refreshToken();
        if (refreshSuccess) {
          // If the token refresh succeeds, check login status
          const loggedIn = await checkLoginStatus();
          setIsLoggedIn(loggedIn);

          // If the user is logged in, fetch their role
          if (loggedIn) {
            const data = await getUserStuff();
            setUserRole(data?.role);
            setUsername(data?.username)
          }
        } else {
          // Token refresh failed, mark the user as logged out
          setIsLoggedIn(false);
        }
      } catch (error) {
        console.error("Error initializing auth:", error);
        setIsLoggedIn(false); // Handle any errors by logging out the user
      }
    };

    initializeAuth(); // Run the async initialization
  }, []);

  return (
    <AuthContext.Provider
      value={{ isLoggedIn, setIsLoggedIn, userRole, setUserRole, username, setUsername, getUserStuff,checkLoginStatus }}
    >
      {children}
    </AuthContext.Provider>
  );
}
