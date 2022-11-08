import { getSession } from "next-auth/react";

export const getFrontendToken = async (): Promise<string> => {
  const session = await getSession();
  return (session as any)?.accessToken || "";
};
