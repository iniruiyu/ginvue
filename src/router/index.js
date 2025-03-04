import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'

import element from '@/views/element.vue'
import userRoutes from './module/user'

import { useStore } from 'vuex'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'element',
      component: element,
    },
    {
      path: '/home',
      name: 'home',
      component: Home,
    },
    ...userRoutes,
  ],
})

router.beforeEach((to, from, next) => {
  // 要去的路由
  if (to.meta.auth) {
    // 判断是否需要登录
    const store = useStore()
    if (store.state.userModule.token) {
      // 这里还需要判断token的有效性，有没有过期，需要后台发放的时候，带上token的有效期
      // 如果Token无效需要请求Token
      next()
    } else {
      router.push({ path: '/login' })
    }
  } else {
    next()
  }
})

export default router
