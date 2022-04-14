import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
const proxyConfig = {
  '/graphql': {
    target: 'https://hn.hsblhsn.me/',
    changeOrigin: true,
  },
  '/images.jpeg': {
    target: 'https://hn.hsblhsn.me/',
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
