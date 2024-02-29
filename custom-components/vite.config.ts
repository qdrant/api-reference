import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: '../fern/dist',
    rollupOptions: {
      output: {
        entryFileNames: `output.js`,
        assetFileNames: `output.[ext]`,
      }
    }
  }
})
