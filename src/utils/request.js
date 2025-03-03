import storageService from '@/services/storageService'

import axios from 'axios'

//console.log(import.meta.env.VITE_APP_BASE_URL)

// 创建axios实例
const service = axios.create({
  //baseURL: 'http://localhost:8080',
  //process是nodejs维护的全局变量
  // vite弃用 改写成import.meta.env.VITE_APP_BASE_URL
  baseURL: import.meta.env.VITE_APP_BASE_URL,
  timeout: 1000 * 5,
})

// Add a request interceptor  添加一个请求拦截器
service.interceptors.request.use(
  function (config) {
    // Do something before request is sent
    // 在发送请求之前做些什么,我们要修改Header

    Object.assign(config.headers, {
      Authorization: 'Bearer ' + storageService.get(storageService.USER_TOKEN),
    })

    return config
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error)
  },
)

// 导出axios实例
export default service
