<script lang="ts">
  import type { ConfMetadata } from '$lib/types';

  interface Props {
    conf: ConfMetadata;
  }
  const { conf }: Props = $props();
</script>

<div>
  <img src={conf.icon} alt={conf.name} class="mr-2 inline h-6 w-6" />
  {conf.name}
</div>
