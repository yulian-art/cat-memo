import vue from '@vitejs/plugin-vue'
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      // 浏览器里发到 /v1/... 的请求，Vite 会原样转发给 Go 后端
      '/v1': 'http://127.0.0.1:9000',
    },
  },
})