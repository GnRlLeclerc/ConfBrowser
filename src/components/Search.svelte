<script lang="ts">
  // Debounced search input component
  import { Search } from '@lucide/svelte';
  interface Props {
    search: string;
  }

  let { search = $bindable() }: Props = $props();

  // Makes the input value take the search value when navigating to a page
  // that stores the search query in the URL
  let value = $derived(search);
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
