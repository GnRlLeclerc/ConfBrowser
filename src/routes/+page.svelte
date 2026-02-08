<script lang="ts">
  import ConfPicker from '$components/ConfPicker.svelte';
  import YearPicker from '$components/YearPicker.svelte';
  import { Search } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import { resolve } from '$app/paths';
  import { page } from '$app/state';

  let conference = $derived(page.url.searchParams.get('conference'));
  let year = $derived(page.url.searchParams.get('year'));
  let disabled = $derived.by(
    () => conference === null || year === null || !(year.toString().length === 4 && !isNaN(+year)),
  );

  const goToPapers = (event: SubmitEvent) => {
    event.preventDefault();
    if (disabled) return;
    goto(resolve(`/papers/${conference}/${year}`));
  };
</script>

<div class="flex h-full w-full flex-col items-center justify-center gap-6">
  <h1 class="text-4xl font-bold">Conf Browser</h1>

  <form onsubmit={goToPapers}>
    <div class="flex gap-4">
      <ConfPicker
        bind:conference={
          () => conference,
          (v) =>
            // eslint-disable-next-line svelte/no-navigation-without-resolve
            goto(`?conference=${v ?? ''}&year=${year ?? ''}`, {
              replaceState: false,
              noScroll: true,
              keepFocus: true,
            })
        }
      />
      <YearPicker
        bind:year={
          () => (year ? +year : undefined),
          (v) =>
            // eslint-disable-next-line svelte/no-navigation-without-resolve
            goto(`?conference=${conference ?? ''}&year=${v ?? ''}`, {
              replaceState: false,
              noScroll: true,
              keepFocus: true,
            })
        }
      />
      <button class="btn" {disabled} type="submit"><Search size="18" />See papers</button>
    </div>
  </form>
</div>
