import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  base: process.env.STATIC || "/",
  plugins: [
    vue({
      template: {
        preprocessOptions: {
          doctype: "html",
        },
      },
    }),
  ],
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          common: [
            "vue",
            "vuex",
            "vue-router",
            "axios",
            "dayjs",
            "localforage",
            "pako",
          ],
          plus: ["element-plus"],
          pluscss: ["element-plus/lib/theme-chalk/index.css"],
        },
      },
    },
  },
  optimizeDeps: {
    include: ["dayjs", "localforage", "pako"],
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:7001",
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});
