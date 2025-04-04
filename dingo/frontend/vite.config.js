import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import path from "path";
// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    hmr: true // 啟用熱重載
  },
  resolve: {
    alias: {
      $lib: path.resolve("./src/lib"),
    },
  },
});
