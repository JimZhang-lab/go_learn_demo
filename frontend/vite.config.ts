/*
 * @Author: JimZhang
 * @Date: 2025-07-22 17:29:27
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 23:34:16
 * @FilePath: /server/frontend/vite.config.ts
 * @Description:
 *
 */
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import tailwindcss from "@tailwindcss/vite";

import {resolve} from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [tailwindcss(), vue()],
  // 路径别名
  resolve: {
    alias: {
      // 绝对路径
      "@": resolve(__dirname, "src"),
    },
    // 扩展名，导入文件时可以省略
    extensions: ['.js', '.ts', '.vue', '.json', '.tsx', '.mjs']
  },
  server: {
    // 端口
    port: 8080,
    // 启动时是否自动打开浏览器
    open: false,
    // 允许跨域
    cors: true,
    // 允许代理
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
});
