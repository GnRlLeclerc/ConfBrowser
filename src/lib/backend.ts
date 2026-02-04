/** Backend URL in dev vs production mode */
export const BASE_URL =
  import.meta.env.MODE === 'development' ? 'http://localhost:8080/backend' : '/backend';

/** Close the backend (only in production mode to avoid disturbing dev workflow) */
export const exit = () => {
  if (import.meta.env.MODE !== 'development') {
    navigator.sendBeacon(BASE_URL + '/exit');
  }
};

/** Fetch data from the backend */
export const fetchBackend = async <T>(endpoint: string): Promise<T> => {
  const data = await fetch(BASE_URL + endpoint);
  return await data.json();
};
