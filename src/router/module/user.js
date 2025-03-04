import Register from '@/views/register/RegisterForm.vue'
import LoginForm from '@/views/login/LoginForm.vue'
const userRoutes = [
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
  {
    path: '/profile',
    name: 'profile',
    meta: {
      // 路由元信息，用于控制路由访问权限
      auth: true,
    },
    component: () => import('@/views/profile/Profile.vue'),
  },
]

export default userRoutes
