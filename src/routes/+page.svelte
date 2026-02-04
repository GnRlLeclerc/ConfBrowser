<script lang="ts">
  import { onMount } from 'svelte';
  import { exit } from '$lib/backend';
  import ConfPicker from '$components/ConfPicker.svelte';
  import YearPicker from '$components/YearPicker.svelte';
  import { Search } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import { resolve } from '$app/paths';

  // Stop the backend once the tab is closed
  onMount(() => {
    window.addEventListener('beforeunload', (event) => {
      exit();
      event.preventDefault();
    });
  });

  let conference = $state<string | undefined>();
  let year = $state<number | undefined>();
  let disabled = $derived(
    () =>
      conference === undefined ||
      year === undefined ||
      !(year.toString().length === 4 && !isNaN(year)),
  );

  const goToPapers = () => {
    if (disabled()) return;
    goto(resolve(`/papers/${conference}/${year}`));
  };
</script>

<div class="flex h-screen w-full flex-col items-center justify-center gap-6">
  <h1 class="text-4xl font-bold">Conf Browser</h1>

  <div class="flex gap-4">
    <ConfPicker bind:conference />
    <YearPicker bind:year />
    <button class="btn" disabled={disabled()} onclick={goToPapers}
      ><Search size="18" />See papers</button
    >
  </div>
</div>
