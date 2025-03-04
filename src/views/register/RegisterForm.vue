<script lang="ts" setup>
import { reactive, ref } from 'vue'
import router from '@/router'

// 导入刚刚创建的services/storageServices.js 缓存服务
import storageService from '@/services/storageService'

// 导入Vuex导出的store
// VUE3组合式这么使用
import { useStore } from 'vuex'
const store = useStore()
// 命名空间用法
// import { createNamespacedHelpers } from 'vuex';
// const mapMutaions = createNamespacedHelpers('userModule')

const ShowtelephoneValidate = ref(false)

const user = reactive({
  name: '45645',
  password: '111111',
  telephone: '12345678888',
})

function register() {
  //表单验证
  if (user.telephone.length != 11) {
    ShowtelephoneValidate.value = true
    return
  }
  //请求
  //const api = 'http://localhost:8080/api/auth/register'
  //axios.post(api,{...user})
  store
    .dispatch('userModule/register', user)
    .then(() => {
      router.push({ path: '/home' }) // 跳转主页
    })
    .catch((err) => {
      alert(err.response.data.msg)
    })
}
</script>

<template>
  <div class="container">
    <el-form :inline="true" :model="user" class="demo-form-inline">
      <h2>注册账号</h2>
      <el-form-item label="用户名">
        <el-input v-model="user.name" placeholder="username" clearable />
      </el-form-item>

      <br />

      <el-form-item label="手机号">
        <el-input v-model="user.telephone" placeholder="username" clearable />
      </el-form-item>

      <br />

      <el-form-item label="密 码 :">
        <el-input v-model="user.password" placeholder="username" clearable />
      </el-form-item>

      <br />

      <el-form-item>
        <el-button type="primary" @click="register">注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style>
.container {
  margin-top: auto;
  display: flex;

  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  height: 366px; /* 设置容器高度 */
  /*border: 1px solid black; /* 可视化边界 */
  background-color: #eeeeee;
}
.demo-form-inline .el-input {
  --el-input-width: 220px;
}

.demo-form-inline .el-select {
  --el-select-width: 220px;
}
</style>
