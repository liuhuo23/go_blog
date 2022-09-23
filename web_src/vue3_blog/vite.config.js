import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import requireMs from "@vitejs/plugin-vue";
import path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  alias: {
    // 键必须以斜线开始斜线结束
    '/@/': path.resolve(__dirname, './src')
  },
  hostname: 'localhost',
  port: 8081,
  open: true,
  http: false,
  ssr: false,
  base: './',
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true,
      rewrite: path => path.replace(/^\/api/, '')
    },
  },
  plugins: [vue()]
})
