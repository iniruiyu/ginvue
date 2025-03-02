<script setup>
import axios from 'axios';
import { reactive, ref } from 'vue';

const ShowtelephoneValidate = ref(false)

const user = reactive({
  name:"45645",
  password:"111111",
  telephone:"12345678888"
})



function register(){

  //表单验证
  if(user.telephone.length != 11){
    ShowtelephoneValidate.value = true
    return
  }
  //请求
  const api ='http://localhost:8080/api/auth/register'

  axios.post(api,{...user})
  .then(res =>{
    console.log(res);

    // 保存token
    // 跳转主页
  }).catch(err=>
    console.log('err',err.response.data.msg)
  )

}
</script>
<template>

<div>
  register
</div>

name<input type="text" v-model="user.name">
<br>
手机号<input type="text" v-model="user.telephone">
<br>
<span style="color: red;font-size: small;" v-if="ShowtelephoneValidate">手机号无效</span>
<br>
密码<input type="password" v-model="user.password">

<button @click="register">注册</button>



</template>


