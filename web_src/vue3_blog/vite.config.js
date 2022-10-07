import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import requireMs from "@vitejs/plugin-vue";
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import * as path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  base: '/',
  resolve: {
    // 配置路径别名
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
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
  plugins: [
      vue(),
      AutoImport({
        imports: [
            'vue',
          {
            'naive-ui':[
              'useDialog',
              'useMessage',
              'useNotification',
              'useLoadingBar'
            ]
          }
        ]
      }),
      Components({
        resolvers: [NaiveUiResolver()]
      })
  ],
    css: {
      preprocessorOptions: {
          scss: {
              additionalData: '@import "../assets/scss/globalvar.scss"; @import "../assets/scss/globalMixin.scss"'
          }
      }
    }
})
