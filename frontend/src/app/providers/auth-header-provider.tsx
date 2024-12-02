import { createContext, useContext, useState } from "react";

interface AuthHeaderContextProps {
  pageTitle: string;
  setPageTitle: (value: string) => void;
}

const AuthHeaderContext = createContext<AuthHeaderContextProps | undefined>(undefined);

export const useAuthHeader = () => {
  const context = useContext(AuthHeaderContext);
  if (!context) {
    throw new Error("useAuthHeader must be used with auth header provider");
  }
  return context;
};

export function AuthHeaderProvider({ children }: { children: React.ReactNode }) {
  const [pageTitle, setPageTitle] = useState<string>("");

  return (
    <AuthHeaderContext.Provider
      value={{
        pageTitle,
        setPageTitle,
      }}
    >
      {children}
    </AuthHeaderContext.Provider>
  );
}
