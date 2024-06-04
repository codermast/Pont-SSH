import { createApp } from 'vue'
import App from './App.vue'

// 注册路由
import router from './router'

let app = createApp(App);
app.use(router)

app.mount('#app')
