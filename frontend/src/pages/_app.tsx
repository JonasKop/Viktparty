import { SessionProvider, signIn, useSession } from "next-auth/react";
import type { AppProps } from "next/app";
import { ChakraProvider, Stack, Spinner, Box } from "@chakra-ui/react";
import { ReactElement } from "react";
import Head from "next/head";

interface WrapperProps {
  children: ReactElement;
}

function Wrapper({ children }: WrapperProps) {
  return (
    <Box>
      <Head>
        <title>Viktparty</title>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      {children}
    </Box>
  );
}

function Middleware({ pageProps, Component }: AppProps) {
  const session = useSession();

  if (session.status === "loading" || session.status === "unauthenticated") {
    if (session.status === "unauthenticated") {
      signIn("keycloak", { callbackUrl: "/auth/callback" });
    }

    return (
      <Stack
        height="100vh"
        width="100vw"
        alignItems="center"
        justifyContent="center"
      >
        <Spinner size="xl" height="100px" width="100px" />
      </Stack>
    );
  }

  return <Component {...pageProps} />;
}

export default function App({
  pageProps: { session, ...pageProps },
  ...rest
}: AppProps) {
  return (
    <ChakraProvider resetCSS={true}>
      <SessionProvider session={session}>
        <Wrapper>
          <Middleware pageProps={pageProps} {...rest} />
        </Wrapper>
      </SessionProvider>
    </ChakraProvider>
  );
}
