import { createStore } from 'vuex';
import storageService from '@/services/storageService';

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
      storageService.set(storageService.USER_TOKEN, token);
      // 更新 state
      state.token = token;
    },
    SET_USERINFO(state, user) {
      // 更新本地缓存
      storageService.set(storageService.USER_INFO, JSON.stringify(user));
      // 更新 state
      state.userInfo = user; // 注意这里应该是 userInfo 而不是 user
    },
  },

  actions: {
    // 可以根据需要添加 actions，例如异步操作
    login({ commit }, payload) {
      const { token, userInfo } = payload;
      commit('SET_TOKEN', token);
      commit('SET_USERINFO', userInfo);
    },
    logout({ commit }) {
      // 清除 Token 和用户信息
      commit('SET_TOKEN', null);
      commit('SET_USERINFO', null);

      // 清除本地存储
      storageService.remove(storageService.USER_TOKEN);
      storageService.remove(storageService.USER_INFO);
    },
  },

  getters: {
    isLoggedIn: (state) => !!state.token, // 判断是否已登录
    currentUser: (state) => state.userInfo, // 获取当前用户信息
  },
};

export default userModule;
