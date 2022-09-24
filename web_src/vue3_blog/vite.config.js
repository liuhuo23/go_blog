import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import requireMs from "@vitejs/plugin-vue";
import * as path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  alias: {
    // 键必须以斜线开始斜线结束
    '/@/': path.resolve(__dirname, './src')
  },
  server:{
    hostname: 'localhost',
    port: 8081,
    open: true,
    http: false,
    ssr: false,
    base: './',
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:9999',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
      '/1api/noapi':{
        target: 'http://127.0.0.1:9999',
        changeOrigin:true,
        rewrite: (path) => path.replace(/^\/1api\/noapi/, '')
      }
    },
  },
  plugins: [vue()],
    css: {
      preprocessorOptions: {
          scss: {
              additionalData: '@import "../assets/scss/globalvar.scss"; @import "../assets/scss/globalMixin.scss"'
          }
      }
    }
})
