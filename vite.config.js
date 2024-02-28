import { defineConfig } from "vite";

export default defineConfig({
  build: {
    outDir: "./static/dist",
    manifest: true,
    rollupOptions: {
      input: "static/src/main.js",
    },
  },
});
