// src/store/index.js
import { createStore } from 'vuex';

import userModule from './module/user'

// 创建 store 实例
const store = createStore({
  state() {
    return {
      count: 0, // 状态数据
    };
  },
  mutations: {
    increment(state) {
      state.count++; // 直接修改状态
    },
    decrement(state) {
      state.count--; // 直接修改状态
    },
  },
  actions: {
    incrementAsync({ commit }) {
      setTimeout(() => {
        commit('increment'); // 调用 mutation 修改状态
      }, 1000);
    },
  },
  getters: {
    doubleCount(state) {
      return state.count * 2; // 返回计算后的状态
    },
  },
  modules:{
    userModule,
  }
});

export default store;
