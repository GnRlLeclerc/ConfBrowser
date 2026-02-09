<script lang="ts">
  import Conference from './Conference.svelte';
  import type { ConfMetadata } from '$lib/types';

  import NeurIPS from '$assets/neurips.svg';
  import ICLR from '$assets/iclr.svg';
  import ICML from '$assets/icml.svg';
  import CVPR from '$assets/cvf.png';

  interface Props {
    conference: string | null;
  }

  let { conference = $bindable() }: Props = $props();

  const conferences: ConfMetadata[] = [
    { name: 'NeurIPS', icon: NeurIPS },
    { name: 'ICLR', icon: ICLR },
    { name: 'ICML', icon: ICML },
    { name: 'CVPR', icon: CVPR },
  ];

  let conf: ConfMetadata | undefined = $derived(
    conferences.find(({ name }) => name === conference),
  );
</script>

<div class="dropdown dropdown-bottom">
  <div tabindex="0" role="button" class="btn m-1 w-42">
    {#if conf !== undefined}
      <Conference {conf} />
    {:else}
      Pick a conference
    {/if}
  </div>
  <ul tabindex="-1" class="dropdown-content menu z-1 w-42 rounded-box bg-base-100 p-2 shadow-sm">
    {#each conferences as conf (conf.name)}
      <li>
        <button
          type="button"
          onclick={() => {
            conference = conf.name;
            (document.activeElement as HTMLElement).blur();
          }}
        >
          <Conference {conf} />
        </button>
      </li>
    {/each}
  </ul>
</div>
