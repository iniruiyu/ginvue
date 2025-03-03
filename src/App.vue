<script setup>
import { RouterLink, RouterView, useRoute } from 'vue-router'
import NavBar from './views/layout/NavBar.vue';
/*//import element from './views/element.vue'
import storageService from '@/services/storageService'

import { ref } from 'vue'
import router from './router'
const menu_select = ref(true)
//const router = useRoute()

// 用户登录后不显示登录注册
import { computed } from 'vue'

// 导入Vuex导出的store
import { useStore } from 'vuex';
const store = useStore();

const userInfo = computed(() => {
  // 将从缓存中获取修改为，从VUEX中获取
  console.log("computed,",store.state.userModule.userInfo);

  return store.state.userModule.userInfo
})*/
</script>

<template>
  <header>

  </header>
 <NavBar/>

  <RouterView />
</template>


