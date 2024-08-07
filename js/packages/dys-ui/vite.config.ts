import { URL, fileURLToPath } from 'node:url';

import vue from '@vitejs/plugin-vue';
import autoprefixer from 'autoprefixer';
import tailwind from 'tailwindcss';
// import twNesting from 'tailwindcss/nesting';
import { defineConfig } from 'vite';

import { peerDependencies } from './package.json';

export default defineConfig({
  css: {
    postcss: {
      plugins: [
        // twNesting(),
        tailwind(),
        autoprefixer(),
      ],
    },
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  build: {
    lib: {
      // file that contains our components exported
      entry: fileURLToPath(new URL('./src/index.ts', import.meta.url)),
      formats: ['es'],
      fileName: 'dys-ui',
    },
    rollupOptions: {
      // externalize deps that shouldn't be bundled into the library
      external: [...Object.keys(peerDependencies)],
    },
  },
  server: {
    port: 8788,
  },
});
