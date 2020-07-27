import request from "@/request";
import { USERS_ME, USERS_LOGIN, USERS_LOGOUT, USERS } from "@/constants/url";
import { generatePassword } from "@/helpers/util";
import { sha256 } from "@/helpers/crypto";

const mutationUserFetching = "user.fetching";
const mutationUserInfo = "user.info";
const mutationUsersFetching = "users.fetching";
const mutationUsers = "users";

function commitUserInfo(commit, data) {
  commit(mutationUserInfo, {
    account: data.account || "",
    trackID: data.trackID || "",
    email: data.email || "",
  });
}

export default {
  state: {
    // fetching 默认设置为fetching
    fetching: true,
    fetchingUsers: false,
    account: "",
    trackID: "",
    id: 0,
    email: "",
    users: null,
  },
  mutations: {
    [mutationUserFetching](state, value) {
      state.fetching = value;
    },
    [mutationUserInfo](state, value) {
      Object.assign(state, value);
    },
    [mutationUsersFetching](state, value) {
      state.fetchingUsers = value;
    },
    [mutationUsers](state, users) {
      state.users = users;
    },
  },
  actions: {
    async fetchUser({ commit }) {
      commit(mutationUserFetching, true);
      try {
        const { data } = await request.get(USERS_ME);
        commitUserInfo(commit, data);
      } finally {
        commit(mutationUserFetching, false);
      }
    },
    async login({ commit }, { account, password, captcha }) {
      commit(mutationUserFetching, true);
      try {
        const res = await request.get(USERS_LOGIN);
        const { token } = res.data;
        const { data } = await request.post(
          USERS_LOGIN,
          {
            account,
            password: sha256(generatePassword(password) + token),
          },
          {
            headers: {
              "X-Captcha": captcha,
            },
          }
        );
        commitUserInfo(commit, data);
      } finally {
        commit(mutationUserFetching, false);
      }
    },
    async register(_, { account, password, captcha }) {
      await request.post(
        USERS_ME,
        {
          account,
          password: generatePassword(password),
        },
        {
          headers: {
            "X-Captcha": captcha,
          },
        }
      );
    },
    async logout({ commit }) {
      commit(mutationUserFetching, true);
      try {
        await request.delete(USERS_LOGOUT);
        commitUserInfo(commit, {});
      } finally {
        commit(mutationUserFetching, false);
      }
    },
    async updateUser({ commit }, { email }) {
      await request.patch(USERS_ME, {
        email,
      });
      commit(mutationUserInfo, {
        email,
      });
    },
    async listUser({ commit }, params) {
      commit(mutationUsersFetching, true);
      try {
        const { data } = await request.get(USERS, {
          params,
        });
        commit(mutationUsers, data.users);
      } finally {
        commit(mutationUsersFetching, false);
      }
    },
  },
};
