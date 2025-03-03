<template>
    <el-menu
      v-if="!userInfo"
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1" @click="router.push({ path: '/home' })"
        >Processing Center</el-menu-item
      >
      <el-menu-item index="2" disabled>Info</el-menu-item>
      <div class="flex-grow" />
      <el-menu-item @click="router.replace('/login')" index="3">Login</el-menu-item>
      <el-menu-item @click="router.push('/register')" index="4"> Register</el-menu-item>
      <el-sub-menu index="5">
        <template #title>Workspace</template>
        <el-menu-item index="5-1">item one</el-menu-item>
        <el-menu-item index="5-2">item two</el-menu-item>
        <el-menu-item index="5-3">item three</el-menu-item>
        <el-sub-menu index="5-4">
          <template #title>item four</template>
          <el-menu-item index="5-4-1">item one</el-menu-item>
          <el-menu-item index="5-4-2">item two</el-menu-item>
          <el-menu-item index="5-4-3">item three</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
    </el-menu>

    <div class="h-6" />
    <el-menu
      v-if="userInfo"
      :default-active="activeIndex2"
      class="el-menu-demo"
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      @select="handleSelect"
    >
      <el-menu-item index="1" @click="router.push({ path: '/home' })"
        >Processing Center</el-menu-item
      >
      <el-menu-item index="2" disabled>Info</el-menu-item>
      <div class="flex-grow" />
      <el-menu-item @click="router.replace('/login')" index="3">{{ userInfo.name }}</el-menu-item>
      <el-menu-item @click="router.push('/register')" index="4">个人主页 </el-menu-item>
      <el-menu-item @click="router.push('/register')" index="5">登出 </el-menu-item>
      <el-sub-menu index="5">
        <template #title>Workspace</template>
        <el-menu-item index="5-1">item one</el-menu-item>
        <el-menu-item index="5-2">item two</el-menu-item>
        <el-menu-item index="5-3">item three</el-menu-item>
        <el-sub-menu index="5-4">
          <template #title>item four</template>
          <el-menu-item index="5-4-1">item one</el-menu-item>
          <el-menu-item index="5-4-2">item two</el-menu-item>
          <el-menu-item index="5-4-3">item three</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
    </el-menu>

</template>
<script setup>
import { RouterLink, RouterView, useRoute } from 'vue-router'

//import element from './views/element.vue'
import storageService from '@/services/storageService'

import { ref } from 'vue'
import router from '@/router'
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
})
</script>

<style scoped>
.flex-grow {
  flex-grow: 1;
}
</style>
