import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { visualizer } from 'rollup-plugin-visualizer'
const proxyConfig = {
  '/graphql': {
    target: 'http://localhost:8080/',
    changeOrigin: true,
  },
  '/images.jpeg': {
    target: 'http://localhost:8080/',
    changeOrigin: true,
  },
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: './frontend/build',
    rollupOptions: {
      plugins: [visualizer()],
    },
  },
  server: {
    proxy: proxyConfig,
  },
  preview: {
    proxy: proxyConfig,
  },
})
