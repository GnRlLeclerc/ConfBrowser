<script lang="ts">
  import ConfPicker from '$components/ConfPicker.svelte';
  import YearPicker from '$components/YearPicker.svelte';
  import { Search } from '@lucide/svelte';
  import { goto } from '$app/navigation';
  import { resolve } from '$app/paths';

  let conference = $state<string | undefined>();
  let year = $state<number | undefined>();
  let disabled = $derived.by(
    () =>
      conference === undefined ||
      year === undefined ||
      !(year.toString().length === 4 && !isNaN(year)),
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
      <ConfPicker bind:conference />
      <YearPicker bind:year />
      <button class="btn" {disabled} type="submit"><Search size="18" />See papers</button>
    </div>
  </form>
</div>
