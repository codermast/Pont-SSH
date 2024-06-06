import { createApp } from 'vue'
// 引入createPinia，用于创建pinia
import { createPinia } from 'pinia'

import App from './App.vue'

// 注册路由
import router from './router'

// 创建pinia
const pinia = createPinia()


let app = createApp(App);
app.use(pinia)
app.use(router)
app.mount('#app')
