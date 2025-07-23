/*
 * @Author: JimZhang
 * @Date: 2025-07-22 17:29:27
 * @LastEditors: 很拉风的James
 * @LastEditTime: 2025-07-22 23:34:56
 * @FilePath: /server/frontend/src/main.ts
 * @Description: 
 * 
 */
import { createApp } from 'vue'
import '@/style.css'
import App from '@/App.vue'
import {initApp} from '@/config/init'

// 对各个浏览器的基础样式进行抹平
import 'normalize.css/normalize.css'

(async () => {

    // 初始化系统基础配置信息(保证所有的模块的基础数据加载后才加载 ui)
    // 1. 全局变量初始化 (app), 语言包(lpk)，Ajax，Tools 的定义
    // 2. 异步加载基础模块的配置信息
    // 3. 异步加载业务模块，并完成基本的初始化

    await initApp()

    // 初始化 ui
    const uiApp = createApp(App)

    // 注册全局组件
    

    // 向根组件绑定全局对象
    uiApp.config.globalProperties.app = window.app
    uiApp.config.globalProperties.Tools = window.Tools
    

    // 初始化状态管理与路由，并渲染根组件
    uiApp.mount('#app')
})()
