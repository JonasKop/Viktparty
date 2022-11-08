import { Box, Text, Input, Button, Stack } from "@chakra-ui/react";
import React, { ReactElement, useEffect, useState } from "react";
import {
  fetchDeleteTodaysWeight,
  fetchNewWeight,
  fetchTodaysWeight,
} from "../api";
import { Link } from "@chakra-ui/react";
import { getSSRToken, withSSRAuth } from "../hooks/withSSRAuth";
import { getFrontendToken } from "../hooks/auth";
import { getSession, signOut } from "next-auth/react";

interface HomeProps {
  errorCode: number;
  weight: number;
}

function Home({ weight: defaultWeight, errorCode }: HomeProps) {
  const [weight, setWeight] = useState<string>(defaultWeight?.toString() || "");
  const [checkedIn, setCheckedIn] = useState<boolean | undefined>(
    !!defaultWeight
  );

  const handleSubmit: React.MouseEventHandler<HTMLElement> = async (e) => {
    e.preventDefault();
    if (weight) {
      const token = await getFrontendToken();
      const resp = await fetchNewWeight(token, parseFloat(weight));
      if (resp.status < 400) {
        setCheckedIn(true);
      }
    }
  };

  const handleDelete = async () => {
    const token = await getFrontendToken();
    const resp = await fetchDeleteTodaysWeight(token);
    if (resp.status < 400) {
      setCheckedIn(false);
      setWeight("");
    }
  };

  if (errorCode) {
    return (
      <Stack
        height="100vh"
        width="100vw"
        alignItems="center"
        justifyContent="center"
      >
        <Text>Error: Could not connect to backend: {errorCode}</Text>
      </Stack>
    );
  }

  return (
    <Box>
      <Box
        display="grid"
        height="100vh"
        justifyContent="center"
        gridTemplateRows="4fr 4fr 1fr"
      >
        <Text
          display="grid"
          pb="10"
          alignSelf="end"
          justifyContent="center"
          fontSize="6xl"
        >
          Viktparty
        </Text>
        {!checkedIn && (
          <Box>
            <form>
              <Input
                placeholder="Vikt"
                type="number"
                value={(weight || "").toString()}
                onChange={(e) => setWeight(e.target.value)}
                mb={2}
              />
              <Button type="submit" colorScheme="teal" onClick={handleSubmit}>
                Spara
              </Button>
            </form>
          </Box>
        )}
        {checkedIn && (
          <Box display="grid" gridAutoFlow="column" gridGap={2}>
            <Text fontSize="2xl">Vikt: {weight} kg</Text>
            <Button colorScheme="teal" onClick={handleDelete}>
              Ångra
            </Button>
          </Box>
        )}
        <Box display="grid" justifyItems="center">
          <Link href="https://grafana.home.jonassjodin.com/">Se historik</Link>
        </Box>
      </Box>
    </Box>
  );
}

interface PageWrapperProps {
  children: ReactElement;
}

const PageWrapper = (Component: any) =>
  function PageWrapper(props: any) {
    useEffect(() => {
      (async () => {
        const session = await getSession();
        if ((session as any)?.error === "RefreshAccessTokenError") signOut();
      })();
    }, []);

    return <Component {...props} />;
  };

export default PageWrapper(Home);

export const getServerSideProps = withSSRAuth(async ({ req, res }: any) => {
  const token = await getSSRToken(req, res);
  const resp = await fetchTodaysWeight(token);
  if (resp.status === 400) {
    return { props: { weight: null } };
  }
  if (resp.status >= 400) {
    return { props: { errorCode: resp.status } };
  }
  const jsonBody = await resp.json();

  return { props: { weight: jsonBody.weight } };
});
