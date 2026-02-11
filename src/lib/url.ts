import { goto } from '$app/navigation';
import { page } from '$app/state';

/** Set a url search param */
export const navigate = (key: string, value: string | null) => {
  const params = page.url.searchParams;
  if (value === '' || value === null) {
    params.delete(key);
  } else {
    params.set(key, value);
  }

  // eslint-disable-next-line svelte/no-navigation-without-resolve
  goto(`?${params}`, {
    replaceState: false,
    noScroll: true,
    keepFocus: true,
  });
};
