<script lang="ts">
  import katex from 'katex';
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

<svelte:head>
  <link
    rel="stylesheet"
    href="https://cdn.jsdelivr.net/npm/katex@0.16.28/dist/katex.min.css"
    integrity="sha384-Wsr4Nh3yrvMf2KCebJchRJoVo1gTU6kcP05uRSh5NV3sj9+a8IomuJoQzf3sMq4T"
    crossorigin="anonymous"
  />

  <!-- The loading of KaTeX is deferred to speed up page rendering -->
  <script
    defer
    src="https://cdn.jsdelivr.net/npm/katex@0.16.28/dist/katex.min.js"
    integrity="sha384-+W9OcrYK2/bD7BmUAk+xeFAyKp0QjyRQUCxeU31dfyTt/FrPsUgaBTLLkVf33qWt"
    crossorigin="anonymous"
  ></script>
</svelte:head>

<!-- eslint-disable-next-line svelte/no-at-html-tags -->
{@html html}
