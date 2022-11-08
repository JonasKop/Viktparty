import { GetServerSideProps } from "next";
import { unstable_getServerSession } from "next-auth";
import { signOut } from "next-auth/react";
import { authOptions } from "../pages/api/auth/[...nextauth]";

export const getSSRToken = async (req: any, res: any): Promise<string> => {
  const session = await unstable_getServerSession(req, res, authOptions as any);
  return (session as any)?.accessToken || "";
};

export const withSSRAuth = (fn: GetServerSideProps) => async (ctx: any) => {
  const token = await getSSRToken(ctx.req, ctx.res);
  if (token) return fn(ctx);
  return { props: {} };
};
