import { createStore } from 'vuex'
import storageService from '@/services/storageService'
import userService from '@/services/userService'

const userModule = {
  namespaced: true, // 开启命名空间

  state: () => ({
    // 定义两个状态
    token: storageService.get(storageService.USER_TOKEN) || null, // 如果缓存中没有值，默认为 null
    userInfo: JSON.parse(storageService.get(storageService.USER_INFO)) || null, // 如果缓存中没有值，默认为 null
  }),

  mutations: {
    SET_TOKEN(state, token) {
      // 更新本地缓存
      storageService.set(storageService.USER_TOKEN, token)
      // 更新 state
      state.token = token
    },
    SET_USERINFO(state, user) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(user))
      // 更新 state
      state.userInfo = user // 注意这里应该是 userInfo 而不是 user
    },
  },

  actions: {
    // 可以根据需要添加 actions，例如异步操作
    login({ commit }, {telephone,password}) {
      return new Promise((resolve,rejet)=>{
        userService.login({telephone,password})
        .then((res)=>{
          // 保存token
          commit('SET_TOKEN',res.data.data.token)
          return userService.info()
        }).then(
            (res)=>{
            commit('SET_USERINFO',res.data.user),
            resolve(res)
            }
        )
        .catch((err)=>{
          rejet(err)
        })
      })
    },
    loginOut({ commit }) {
      // 清除 Token 和用户信息
      commit('SET_TOKEN', null)
      commit('SET_USERINFO', null)

      // 清除本地存储
      storageService.remove(storageService.USER_TOKEN)
      storageService.remove(storageService.USER_INFO)
    },
    register({ commit }, { name, telephone, password }) {
      return new Promise((resolve, reject) => {
        userService
          .register({ name, telephone, password })
          .then((res) => {
            console.log('res-', res)
            // 保存token
            //store.commit('userModule/SET_TOKEN', res.data.data.token)
            commit('SET_TOKEN',res.data.data.token)
            return userService.info() // 在成功的返回里面，再返回一个promise，然后就可以链式调用then
          })
          .then((res) => {
            // 保存用户信息
            // store.commit('userModule/SET_USERINFO', res.data.user) // 出错就盯着返回的RES看，看看是不是返回的数据结构不对
            commit('SET_USERINFO',res.data.user)
            // 成功resolve一下
            resolve(res)
          })
          .catch((err) => {
            reject(err)
          })
      })
    },
  },
  getters: {
    isLoggedIn: (state) => !!state.token, // 判断是否已登录
    currentUser: (state) => state.userInfo, // 获取当前用户信息
  },
}

export default userModule
