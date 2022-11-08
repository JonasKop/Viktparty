import KeycloakProvider from "next-auth/providers/keycloak";

const clientId = process.env.KEYCLOAK_CLIENT_ID!;
const clientSecret = process.env.KEYCLOAK_CLIENT_SECRET!;
const issuer = process.env.KEYCLOAK_ISSUER!;

import NextAuth from "next-auth";
/**
 * Takes a token, and returns a new token with updated
 * `accessToken` and `accessTokenExpires`. If an error occurs,
 * returns the old token and an error property
 */
async function refreshAccessToken(token: string) {
  try {
    const searchParams = new URLSearchParams({
      client_id: clientId,
      client_secret: clientSecret,
      grant_type: "refresh_token",
      refresh_token: (token as any).refreshToken,
    });

    const response = await fetch(`${issuer}/protocol/openid-connect/token`, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      method: "POST",
      body: searchParams,
    });

    const refreshedTokens = await response.json();

    if (!response.ok) {
      throw refreshedTokens;
    }

    return {
      ...(token as any),
      accessToken: refreshedTokens.access_token,
      accessTokenExpires: Date.now() + refreshedTokens.expires_at * 1000,
      refreshToken:
        refreshedTokens.refresh_token ?? (token as any).refreshToken, // Fall back to old refresh token
    };
  } catch (error) {
    console.log(error);

    return {
      ...(token as any),
      error: "RefreshAccessTokenError",
    };
  }
}

function parseJwt(token: string) {
  return JSON.parse(Buffer.from(token.split(".")[1], "base64").toString());
}

export const authOptions = {
  providers: [
    KeycloakProvider({
      clientId,
      clientSecret,
      issuer,
    }),
  ],
  callbacks: {
    async jwt({ token, user, account }: any) {
      // Initial sign in
      if (account && user) {
        return {
          accessToken: account.access_token,
          accessTokenExpires: Date.now() + account.expires_at * 1000,
          refreshToken: account.refresh_token,
          user,
        };
      }

      const jwt = parseJwt(token.accessToken);
      token.accessTokenExpires = jwt.exp * 1000;
      // Return previous token if the access token has not expired yet
      if (Date.now() < token.accessTokenExpires - 10000) {
        const secondsLeft = Math.floor(
          (token.accessTokenExpires - Date.now()) / 1000
        );
        console.log(`Token has not expired, time left: ${secondsLeft}s`);
        return token;
      }
      console.log("Token has expired: refreshing");
      // Access token has expired, try to update it
      return refreshAccessToken(token);
    },
    async session({ session, token }: any) {
      session.user = token.user;
      session.accessToken = token.accessToken;
      session.error = token.error;

      return session;
    },
  },
};

export default NextAuth(authOptions as any);
