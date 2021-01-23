import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  build: {
    base: "/static",
  },
  plugins: [vue({
    template: {
      preprocessOptions: {
        doctype: "html",
      },
    },
  })],
  optimizeDeps: {
    include: [
      "dayjs",
      "localforage",
      "pako",
    ],
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:7001",
        rewrite: path => path.replace(/^\/api/, ''),
      },
    }
  },
});
