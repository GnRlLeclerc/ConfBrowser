import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter({ fallback: 'index.html' }),
    alias: {
      $components: 'src/components',
      $assets: 'src/assets',
    },
  },
};

export default config;
