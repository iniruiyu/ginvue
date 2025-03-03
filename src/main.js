import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// 引入 Element Plus 的样式和组件
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

//导入elementPlus中的所有图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'


// vuex
import store from './stores'; // 引入 store



const app = createApp(App)

// 使用 Element Plus
app.use(ElementPlus)

//您需要从 @element-plus/icons-vue 中导入所有图标并进行全局注册。
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 使用 Vuex store
app.use(store);

app.use(createPinia())
app.use(router)

app.mount('#app')

