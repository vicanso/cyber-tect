import request from '@/request'
import {
  USERS_ME
} from '@/urls'

const mutationUserFetching = 'user.fetching'
const mutationUserInfo = 'user.info'

export default {
  state: {
    fetching: false,
    account: '',
    trackID: ''
  },
  mutations: {
    [mutationUserFetching] (state, value) {
      state.fetching = value
    },
    [mutationUserInfo] (state, value) {
      Object.assign(state, value)
    }
  },
  actions: {
    async userFetch ({ commit }) {
      commit(mutationUserFetching, true)
      try {
        const {
          data
        } = await request.get(USERS_ME)
        commit(mutationUserInfo, {
          account: data.account || '',
          trackID: data.trackID
        })
      } finally {
        commit(mutationUserFetching, false)
      }
    }
  }
}
