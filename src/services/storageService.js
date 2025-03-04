// 本地缓存服务
const PREFIX = 'ginessential_'

// user模块
const USER_PREFIX = `${PREFIX}user_`
const USER_TOKEN = `${USER_PREFIX}token`
const USER_INFO = `${USER_PREFIX}info`

//储存
const set = (key, data) => {
  // 替换缓存方案的具体实现
  // 比如说更换cookie，其他的地方不需要修改
  localStorage.setItem(key, data)
}

//读取
const get = (key) => {
  return localStorage.getItem(key)
}

//删除
const remove = (key) => {
  return localStorage.removeItem(key)
}

// 导出
export default {
  set,
  get,
  USER_TOKEN,
  USER_INFO,
  remove
}
