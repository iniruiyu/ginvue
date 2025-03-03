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
  headers: {
    Authorization: 'Bearer ' + storageService.get(storageService.USER_TOKEN),
  },
})

// 导出axios实例
export default service
