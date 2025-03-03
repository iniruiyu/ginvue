<script lang="ts" setup>
import { reactive, ref } from 'vue'
import router from '@/router'

// 导入刚刚创建的services/storageServices.js 缓存服务
import storageService from '@/services/storageService'
import userService from '@/services/userService'

// 导入Vuex导出的store
// VUE3组合式这么使用
import { useStore } from 'vuex';
const store = useStore();


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
  userService
    .register(user)
    .then((res) => {
      console.log("res-",res)
      // 保存token
      // localStorage.setItem('token',res.data.data.token)
      // storageService.set(storageService.USER_TOKEN, res.data.data.token)
      //console.log(res.data.data.token);
      console.log("then((res1)",res);
      store.commit('userModule/SET_TOKEN', res.data.data.token)

      // 获取用户信息
      // 请求中调用请求，逻辑不清晰，改成链式调用
      return userService.info();  // 在成功的返回里面，再返回一个promise，然后就可以链式调用then
    }).then((res) => {
        // 保存用户信息
        // storageService.set(storageService.USER_INFO, res.data.data.user)
        // 保存 序列化的 用户信息
        // storageService.set(storageService.USER_INFO, JSON.stringify(res.data.user))
        console.log("then((res2)",res);

        store.commit('userModule/SET_USERINFO', res.data.user)  // 出错就盯着返回的RES看，看看是不是返回的数据结构不对

        router.push({ path: '/home' })  // 跳转主页
      })
      .catch((err) => {
      console.log('err',err)
      /*if (err.response.data.msg){
      alert(err.response.data.msg)
      }*/
    })

}
</script>

<template>
  <div class="container">
    <!-- <div>register</div>

  name<input type="text" v-model="user.name" />
  <br />
  手机号<input type="text" v-model="user.telephone" />
  <br />
  <span style="color: red; font-size: small" v-if="ShowtelephoneValidate">手机号无效</span>
  <br />
  密码<input type="password" v-model="user.password" />

  <button @click="register">注册</button> -->

    <el-form :inline="true" :model="user" class="demo-form-inline">
      <el-form-item label="用户名">
        <el-input v-model="user.name" placeholder="username" clearable />
      </el-form-item>

      <br />

      <el-form-item label="手机号">
        <el-input v-model="user.telephone" placeholder="username" clearable />
      </el-form-item>
      <!-- <el-form-item label="手机号">
      <el-select
        v-model="user.telephone"
        placeholder="Telephone"
        clearable
      >
        <el-option label="Zone one" value="shanghai" />
        <el-option label="Zone two" value="beijing" />
      </el-select>
    </el-form-item> -->

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
        <el-button type="primary" @click="register">注册</el-button>
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
