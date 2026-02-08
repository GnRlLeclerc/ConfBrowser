<script lang="ts">
  import '../app.css';
  import { onMount } from 'svelte';
  import { exit, fetchBackend } from '$lib/backend';
  import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
  const queryClient = new QueryClient();

  onMount(() => {
    fetchBackend('/mount'); // Tell the backend a new tab is open

    // Tell the backend the tab just closed
    window.addEventListener('beforeunload', (event) => {
      exit();
      event.preventDefault();
    });
  });
</script>

<QueryClientProvider client={queryClient}>
  <div class="h-screen p-6">
    <slot />
  </div>
</QueryClientProvider>
