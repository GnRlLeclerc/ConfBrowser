<script lang="ts">
  import { Calendar } from '@lucide/svelte';

  interface Props {
    year: number | undefined;
  }

  let { year = $bindable() }: Props = $props();
</script>

<label class="validator input w-24">
  <Calendar size="22" />
  <input
    class="tabular-nums"
    required
    placeholder="Year"
    pattern="[0-9]*"
    minlength="4"
    maxlength="4"
    bind:value={year}
  />
</label>
