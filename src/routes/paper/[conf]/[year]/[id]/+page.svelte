<script lang="ts">
  import { createQuery } from '@tanstack/svelte-query';
  import type { PageProps } from './$types';
  import { File, Image } from '@lucide/svelte';
  import { fetchBackend } from '$lib/backend';
  import type { Paper } from '$lib/types';
  import Latex from '$components/Latex.svelte';
  let props: PageProps = $props();

  const query = createQuery(() => ({
    queryKey: ['paper', props.params.conf, props.params.year, props.params.id],
    queryFn: () =>
      fetchBackend<Paper>(`/paper/${props.params.conf}/${props.params.year}/${props.params.id}`),
    retry: false,
  }));
</script>

{#if query.isLoading}
  <div>Loading paper...</div>
{:else if query.isError}
  <div class="text-red-500">Error loading paper</div>
{:else if query.data !== undefined}
  <div class="flex flex-col gap-4">
    <h1 class="text-center text-4xl font-bold"><Latex text={query.data.title} /></h1>
    <h2 class="text-center text-2xl italic">{query.data.authors.join(', ')}</h2>
    <div class="m-auto mt-4 flex w-1/2 flex-col gap-2 text-justify">
      <h3 class="text-center text-xl font-bold">Abstract</h3>
      <p>
        <Latex text={query.data.abstract} />
      </p>
    </div>

    <div class="m-auto mt-4 flex w-1/2 items-center justify-center gap-10">
      <a href={query.data.pdf} rel="external" target="_blank" class="flex items-center gap-2"
        ><File />PDF</a
      >
      {#if query.data.poster}
        <a href={query.data.poster} rel="external" target="_blank" class="flex items-center gap-2"
          ><Image />Poster</a
        >
      {/if}
    </div>
  </div>
{/if}
