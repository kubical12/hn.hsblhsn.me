import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { visualizer } from 'rollup-plugin-visualizer'
import { ViteMinifyPlugin } from 'vite-plugin-minify'

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
  plugins: [
    react(),
    visualizer({
      filename: './frontend/build/stats.html',
    }),
    ViteMinifyPlugin({
      esbuildTarget: 'es2015',
    }),
  ],
  build: {
    outDir: './frontend/build',
    rollupOptions: {
      output: {
        manualChunks: {
          react: ['react', 'react-dom', 'react-router', 'react-router-dom'],
          baseui: ['baseui'],
        },
      },
    },
    chunkSizeWarningLimit: 700,
  },
  server: {
    proxy: proxyConfig,
  },
  preview: {
    proxy: proxyConfig,
  },
  publicDir: './frontend/public',
})
