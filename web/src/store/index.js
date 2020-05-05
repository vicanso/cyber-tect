import Vue from 'vue'
import Vuex from 'vuex'

import userStore from '@/store/modules/user'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    user: userStore
  }
})
