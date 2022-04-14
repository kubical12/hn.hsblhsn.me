import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
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
  },
  server: {
    proxy: proxyConfig,
  },
  preview: {
    proxy: proxyConfig,
  },
})
