<script lang="ts">
  import { createQuery } from '@tanstack/svelte-query';
  import type { PageProps } from './$types';
  import { fetchBackend } from '$lib/backend';
  import type { PaperMetadata } from '$lib/types';
  let props: PageProps = $props();

  const query = createQuery(() => ({
    queryKey: ['papers', props.params.conf, props.params.year],
    queryFn: () =>
      fetchBackend<PaperMetadata[]>(`/papers/${props.params.conf}/${props.params.year}`),
  }));
</script>

<div>Hello to the page</div>
{#if query.isLoading}
  <div>Loading papers...</div>
{:else if query.isError}
  <div>Error loading papers: {query.error.message}</div>
{:else if query.data !== undefined}
  <div>
    <h2>Papers for {props.params.conf} {props.params.year}</h2>
    <ul>
      {#each query.data as paper (paper.id)}
        <li>{paper.id}: {paper.title}</li>
      {/each}
    </ul>
  </div>
{/if}
