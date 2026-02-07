<script lang="ts">
  import katex from 'katex';
  import 'katex/dist/katex.min.css';
  import sanitizeHtml from 'sanitize-html';

  interface Props {
    text: string;
  }
  const props: Props = $props();

  // Sanitize text before Katex transformation to prevent XSS attacks.
  const html = $derived(
    sanitizeHtml(props.text).replaceAll(/(\$[^$]+\$)/g, (match) =>
      katex.renderToString(match.slice(1, -1), { throwOnError: false }),
    ),
  );
</script>

<!-- eslint-disable-next-line svelte/no-at-html-tags -->
{@html html}
