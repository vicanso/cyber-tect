import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";

import request from "../../helpers/request";
import { sha256 } from "../../helpers/crypto";
import {
  USERS_ME,
  USERS_LOGIN,
  USERS_LOGINS,
  USERS_ROLES,
  USERS,
  USERS_ID,
} from "../../constants/url";
import { generatePassword } from "../../helpers/util";

const prefix = "user";
const mutationProcessing = `${prefix}.processing`;
const mutationInfo = `${prefix}.info`;

const prefixLogin = `${prefix}.login`;
const mutationLoginListProcessing = `${prefixLogin}.processing`;
const mutationLoginList = `${prefixLogin}.list`;

const prefixRole = `${prefix}.role`;
const mutationRoleListProcessing = `${prefixRole}.processing`;
const mutationRoleList = `${prefixRole}.list`;

const prefixUsers = `${prefix}.users`;
const mutationUserListProcessing = `${prefixUsers}.processing`;
const mutationUserList = `${prefixUsers}.list`;

const prefixUpdate = `${prefix}.update`;
const mutationUpdateProcessing = `${prefixUpdate}.processing`;
const mutationUpdateDone = `${prefixUpdate}.done`;

interface UserInfo {
  processing: boolean;
  date: string;
  account: string;
  trackID: string;
  roles: string[];
  groups: string[];
}
interface UserLoginList {
  processing: boolean;
  count: number;
  items: any[];
}
interface UserRoleList {
  processing: boolean;
  items: any[];
}
interface UserList {
  processing: boolean;
  count: number;
  items: any[];
}
interface UserUpdateInfo {
  processing: boolean;
}

interface UserState {
  info: UserInfo;
  logins: UserLoginList;
  roles: UserRoleList;
  users: UserList;
  updateInfo: UserUpdateInfo;
}

const info: UserInfo = {
  processing: false,
  date: "",
  account: "",
  trackID: "",
  roles: [],
  groups: [],
};

const logins: UserLoginList = {
  processing: false,
  count: -1,
  items: [],
};

const roles: UserRoleList = {
  processing: false,
  items: [],
};

const users: UserList = {
  processing: false,
  count: -1,
  items: [],
};

const updateInfo: UserUpdateInfo = {
  processing: false,
};

const state: UserState = {
  // 用户信息
  info,
  // 用户登录记录
  logins,
  // 用户角色
  roles,
  // 用户列表
  users,
  // 用户信息更新
  updateInfo,
};

const emptyUserInfo: UserInfo = {
  processing: false,
  date: "",
  account: "",
  trackID: "",
  roles: [],
  groups: [],
};

export const userStore = createStore<UserState>({
  state,
  mutations: {
    // 设置处理中
    [mutationProcessing](state: UserState, processing: boolean) {
      state.info.processing = processing;
    },
    // 设置用户信息
    [mutationInfo](state: UserState, info: UserInfo) {
      Object.assign(state.info, info);
    },
    // 设置正在查询用户登录记录
    [mutationLoginListProcessing](state: UserState, processing: boolean) {
      state.logins.processing = processing;
    },
    // 设置查询登录记录结果
    [mutationLoginList](
      state: UserState,
      data: { count: number; userLogins: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.logins.count = count;
      }
      data.userLogins?.forEach((element) => {
        // element.createdAtDesc = formatDate(element.createdAt);
        const locations: string[] = [];
        ["country", "province", "city"].forEach((key) => {
          if (element[key]) {
            locations.push(element[key]);
          }
        });
        element.location = locations.join(" ") || "--";
      });
      state.logins.items = data.userLogins || [];
    },
    // 设置正在查询用户角色
    [mutationRoleListProcessing](state: UserState, processing: boolean) {
      state.roles.processing = processing;
    },
    // 设置用户角色
    [mutationRoleList](state: UserState, data: { userRoles: any[] }) {
      state.roles.items = data.userRoles || [];
    },
    // 设置正在查询用户
    [mutationUserListProcessing](state: UserState, processing: boolean) {
      state.users.processing = processing;
    },
    // 设置用户列表
    [mutationUserList](
      state: UserState,
      data: { users: any[]; count: number }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.users.count = count;
      }
      state.users.items = data.users || [];
    },
    // 设置正在更新用户信息
    [mutationUpdateProcessing](state: UserState, processing: boolean) {
      state.updateInfo.processing = processing;
    },
    // 更新用户信息
    [mutationUpdateDone](state: UserState, data: { id: number; data: any }) {
      const users = state.users.items.slice(0);
      users.forEach((item) => {
        if (item.id === data.id) {
          Object.assign(item, data.data);
        }
      });
      state.users.items = users;
    },
  },
  actions: {
    // fetch 获取用户信息
    async fetch(context: { commit: Commit }) {
      context.commit(mutationProcessing, true);
      try {
        const { data } = await request.get(USERS_ME);
        context.commit(mutationInfo, <UserInfo>data);
      } finally {
        context.commit(mutationProcessing, false);
      }
    },
    // login 用户登录
    async login(
      context: { commit: Commit },
      params: { account: string; password: string; captcha: string }
    ) {
      context.commit(mutationProcessing, true);
      try {
        const res = await request.get(USERS_LOGIN);
        const { token } = res.data;
        const { data } = await request.post(
          USERS_LOGIN,
          {
            account: params.account,
            password: sha256(generatePassword(params.password) + token),
          },
          {
            headers: {
              // 图形验证码
              "X-Captcha": params.captcha,
            },
          }
        );
        context.commit(mutationInfo, <UserInfo>data);
      } finally {
        context.commit(mutationProcessing, false);
      }
    },
    // register 用户注册
    async register(
      context: { commit: Commit },
      params: { account: string; password: string; captcha: string }
    ) {
      context.commit(mutationProcessing, true);
      try {
        await request.post(
          USERS_ME,
          {
            account: params.account,
            password: generatePassword(params.password),
          },
          {
            headers: {
              // 图形验证码
              "X-Captcha": params.captcha,
            },
          }
        );
      } finally {
        context.commit(mutationProcessing, false);
      }
    },
    // logout 退出登录
    async logout(context: { commit: Commit }) {
      context.commit(mutationProcessing, true);
      try {
        await request.delete(USERS_ME);
        context.commit(mutationInfo, emptyUserInfo);
      } finally {
        context.commit(mutationProcessing, false);
      }
    },
    // listLogin 查询用户登录记录
    async listLogin(context: { commit: Commit }, params) {
      context.commit(mutationLoginListProcessing, true);
      try {
        const { data } = await request.get(USERS_LOGINS, {
          params,
        });
        context.commit(mutationLoginList, data);
      } finally {
        context.commit(mutationLoginListProcessing, false);
      }
    },
    // listRole 查询用户角色
    async listRole(context: { commit: Commit }) {
      if (state.roles.items.length !== 0) {
        return;
      }
      context.commit(mutationRoleListProcessing, true);
      try {
        const { data } = await request.get(USERS_ROLES);
        context.commit(mutationRoleList, data);
        return data;
      } finally {
        context.commit(mutationRoleListProcessing, false);
      }
    },
    // list 查询用户
    async list(context: { commit: Commit }, params) {
      context.commit(mutationUserListProcessing, true);
      try {
        const { data } = await request.get(USERS, {
          params,
        });
        context.commit(mutationUserList, data);
      } finally {
        context.commit(mutationUserListProcessing, false);
      }
    },
    // findByID 通过ID查询用户
    async findByID(_, params: { id: number }) {
      const { data } = await request.get(
        USERS_ID.replace(":id", `${params.id}`)
      );
      return data;
    },
    // updateByID 通过ID更新用户信息
    async updateByID(
      context: { commit: Commit },
      params: { id: number; data: any }
    ) {
      context.commit(mutationUpdateProcessing, true);
      const data = Object.assign({}, params.data);
      try {
        // 如果groups未设置，则清空
        ["groups"].forEach((key) => {
          if (data[key] && data[key].length === 0) {
            delete data[key];
          }
        });
        await request.patch(USERS_ID.replace(":id", `${params.id}`), data);
        context.commit(mutationUpdateDone, {
          id: params.id,
          data,
        });
      } finally {
        context.commit(mutationUpdateProcessing, false);
      }
    },
  },
});

// getUserStore 获取用户store
export function getUserStore(): Store<UserState> {
  return userStore;
}
