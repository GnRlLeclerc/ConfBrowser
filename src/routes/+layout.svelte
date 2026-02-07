<script lang="ts">
  import '../app.css';
  import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
  const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
  <div class="h-screen p-6">
    <slot />
  </div>
</QueryClientProvider>
