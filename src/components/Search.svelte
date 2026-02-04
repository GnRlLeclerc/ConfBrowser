<script lang="ts">
  // Debounced search input component
  import { Search } from '@lucide/svelte';
  interface Props {
    search: string;
  }

  let { search = $bindable() }: Props = $props();

  let value = $state('');
  let timeout: ReturnType<typeof setTimeout>;

  const debounce = () => {
    if (timeout) clearTimeout(timeout);

    timeout = setTimeout(() => {
      search = value;
    }, 300);
  };
</script>

<label class="input">
  <Search size="18" />
  <input type="text" placeholder="Filter..." bind:value oninput={debounce} />
</label>
