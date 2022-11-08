const genDefaultHeaders = (accessToken: string) => ({
  Authorization: `Bearer ${accessToken}`,
  Accept: "application/json",
  "Content-Type": "application/json",
});

const serverUrl = process.env.SERVER_URL || "http://localhost:8080";
const baseUrl = typeof window === "undefined" ? serverUrl : "/api/v1";

export const fetchTodaysWeight = async (accessToken: string) =>
  fetch(`${baseUrl}/weight/today`, {
    headers: genDefaultHeaders(accessToken),
  });

export const fetchNewWeight = async (accessToken: string, weight: number) =>
  fetch(`${baseUrl}/weight`, {
    method: "POST",
    headers: genDefaultHeaders(accessToken),
    body: JSON.stringify({ weight }),
  });

export const fetchDeleteTodaysWeight = async (accessToken: string) =>
  fetch(`${baseUrl}/weight/today`, {
    method: "DELETE",
    headers: genDefaultHeaders(accessToken),
  });
