import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

const proxyConfig = {
  '/api': {
    target: 'https://hn.hsblhsn.me/',
    changeOrigin: true,
  },
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: './embedded/dist',
  },
  server: {
    proxy: proxyConfig,
  },
  preview: {
    proxy: proxyConfig,
  },
})
