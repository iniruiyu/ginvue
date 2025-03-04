<script lang="ts" setup>
import { reactive, ref } from 'vue'
import router from '@/router'

// VUE3组合式VUEX4这么使用
import { useStore } from 'vuex'
const store = useStore()


const ShowtelephoneValidate = ref(false)

const user = reactive({
  name: '挖到撒火炬大厦',
  password: '111111',
  telephone: '12345678888',
})

function login() {
  //表单验证
  if (user.telephone.length != 11) {
    ShowtelephoneValidate.value = true
    return
  }
  //请求
  store.dispatch('userModule/login',user)
  .then(()=>{
    router.push({ path: '/home' }) // 跳转主页
  })
  .catch((err) => {
    console.log(err);

      alert(err.response.data.msg)
    })
}
</script>

<template>
  <div class="container">
    <el-form :inline="true" :model="user" class="demo-form-inline">
      <!-- <el-form-item label="用户名">
        <el-input v-model="user.name" placeholder="username" clearable />
      </el-form-item>

      <br /> -->

      <el-form-item label="手机号">
        <el-input v-model="user.telephone" placeholder="username" clearable />
      </el-form-item>
      <br />

      <el-form-item label="密 码 :">
        <el-input v-model="user.password" placeholder="username" clearable />
      </el-form-item>
      <!-- <el-form-item label="Activity time">
      <el-date-picker
        v-model="user.password"
        type="date"
        placeholder="Pick a date"
        clearable
      />
    </el-form-item> -->

      <br />

      <el-form-item>
        <el-button type="primary" @click="login">登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style>
.container {
  display: flex;

  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  height: 200px; /* 设置容器高度 */
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
