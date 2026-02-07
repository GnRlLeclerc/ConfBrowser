<script lang="ts">
  import { createQuery } from '@tanstack/svelte-query';
  import type { PageProps } from './$types';
  import { fetchBackend } from '$lib/backend';
  import type { PaperMetadata } from '$lib/types';
  import Search from '$components/Search.svelte';
  import Latex from '$components/Latex.svelte';
  let props: PageProps = $props();

  const query = createQuery(() => ({
    queryKey: ['papers', props.params.conf, props.params.year],
    queryFn: () =>
      fetchBackend<PaperMetadata[]>(`/papers/${props.params.conf}/${props.params.year}`),
  }));

  let search = $state('');
  let filtered = $derived(
    search === ''
      ? query.data
      : query.data?.filter(({ title }) =>
          search
            .toLowerCase()
            .split(' ')
            .every((substr) => title.toLowerCase().includes(substr)),
        ),
  );
</script>

<div class="flex h-full flex-col">
  <h1 class="text-center text-4xl font-bold">{props.params.conf} {props.params.year}</h1>
  {#if query.isLoading}
    <div>Loading papers...</div>
  {:else if query.isError}
    <div>Error loading papers: {query.error.message}</div>
  {:else if filtered !== undefined}
    <div>
      <!-- Wrap in a div to prevent the input from losing its height padding because of flex layout -->
      <Search bind:search />
    </div>
    <div class="mt-4 flex overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
      <table class="flex flex-col">
        <thead class="table w-full table-fixed">
          <tr>
            <th class="w-24">ID</th>
            <th>Title</th>
          </tr>
        </thead>
        <tbody class="block w-full overflow-y-auto">
          <!-- NOTE: temporary trick to limit number of rendered papers for performance reasons (katex). -->
            <tr class="table w-full rounded-none hover:bg-base-200">
              <td class="w-24">{paper.id}</td>
              <td><Latex text={paper.title} /></td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div>
