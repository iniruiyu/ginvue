import { createRouter, createWebHistory } from 'vue-router'
import Register from '@/views/register/RegisterForm.vue'
import LoginForm from '@/views/login/LoginForm.vue'

import element from '@/views/element.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'element',
      component: element,
    },
    {
      path: '/login',
      name: 'login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: LoginForm,
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
    },
  ],
})

export default router
