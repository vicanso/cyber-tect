import { DeepReadonly, reactive, readonly } from "vue";
import {
  USERS_ME,
  USERS_LOGIN,
  USERS_INNER_LOGIN,
  USERS,
  USERS_LOGINS,
  USERS_ID,
} from "../constants/url";
// eslint-disable-next-line
// @ts-ignore
import { sha256 } from "../helpers/crypto";
import request from "../helpers/request";

const hash = "JT";

function generatePassword(pass: string): string {
  return sha256(hash + sha256(pass + hash));
}

// 用户信息
interface UserInfo {
  processing: boolean;
  date: string;
  account: string;
  groups: string[];
  roles: string[];
}
const info: UserInfo = reactive({
  processing: false,
  date: "",
  account: "",
  groups: [],
  roles: [],
});

// 用户账户信息
interface UserAccount {
  id: number;
  account: string;
  groups: string[];
  roles: string[];
  email: string;
  status: number;
  statusDesc: string;
}
// 用户账户列表
interface UserAccounts {
  processing: boolean;
  count: number;
  items: UserAccount[];
}
const users: UserAccounts = reactive({
  processing: false,
  count: -1,
  items: [],
});

const accountStatusDesc = ["启用", "禁用"];
function fillUserAccountInfo(data: UserAccount) {
  data.statusDesc = accountStatusDesc[data.status - 1] || "未知";
}
function fillUserInfo(data: UserInfo) {
  info.account = data.account;
  info.date = data.date;
  info.roles = data.roles || [];
  info.groups = data.groups || [];
}

// 用户登录信息
interface UserLoginRecord {
  account: string;
  userAgent?: string;
  ip?: string;
  trackID?: string;
  sessionID?: string;
  xForwardedFor?: string;
  country?: string;
  province?: string;
  city?: string;
  isp?: string;
  updatedAt?: string;
  createdAt?: string;
  location?: string;
}
// 用户登录列表
interface UserLoginRecords {
  processing: boolean;
  count: number;
  items: UserLoginRecord[];
}

const logins: UserLoginRecords = reactive({
  processing: false,
  count: -1,
  items: [],
});

function fillUserLoginRecord(data: UserLoginRecord) {
  const arr: string[] = [];
  if (data.province) {
    arr.push(data.province);
  }
  if (data.city) {
    arr.push(data.city);
  }
  data.location = arr.join("");
}

// userFetchInfo 拉取用户信息
export async function userFetchInfo(): Promise<void> {
  // TODO 是否需要针对并发调用出错
  if (info.processing) {
    return;
  }
  try {
    info.processing = true;
    const { data } = await request.get(USERS_ME);
    fillUserInfo(<UserInfo>data);
  } finally {
    info.processing = false;
  }
}

// userLogout 退出登录
export async function userLogout(): Promise<void> {
  if (info.processing) {
    return;
  }
  try {
    info.processing = true;
    await request.delete(USERS_ME);
    fillUserInfo({
      account: "",
      roles: [],
      groups: [],
      date: "",
      processing: false,
    });
  } finally {
    info.processing = false;
  }
}

// userRegister 用户注册
export async function userRegister(params: {
  account: string;
  password: string;
  captcha: string;
}): Promise<void> {
  if (info.processing) {
    return;
  }
  try {
    // 如果密码小于6位或者纯数字
    if (params.password.length < 6 || /^\d+$/.exec(params.password)) {
      throw new Error("密码过于简单，请使用数字加字母且长度大于6位");
    }
    info.processing = true;
    await request.post(
      USERS_ME,
      {
        account: params.account,
        password: generatePassword(params.password),
      },
      {
        headers: {
          "X-Captcha": params.captcha,
        },
      }
    );
  } finally {
    info.processing = false;
  }
}

// userLogin 用户登录
export async function userLogin(params: {
  account: string;
  password: string;
  captcha: string;
}): Promise<void> {
  if (info.processing) {
    return;
  }
  try {
    info.processing = true;
    const resp = await request.get(USERS_LOGIN);
    const { token } = resp.data;
    const { data } = await request.post(
      USERS_INNER_LOGIN,
      {
        account: params.account,
        password: sha256(generatePassword(params.password) + token),
      },
      {
        headers: {
          "X-Captcha": params.captcha,
        },
      }
    );
    fillUserInfo(<UserInfo>data);
  } finally {
    info.processing = false;
  }
}

// userList 查询用户
export async function userList(params: {
  keyword?: string;
  limit: number;
  offset: number;
  role?: string;
  status?: string;
  order?: string;
}): Promise<void> {
  if (users.processing) {
    return;
  }
  try {
    users.processing = true;
    const { data } = await request.get(USERS, {
      params,
    });
    const count = data.count || 0;
    if (count >= 0) {
      users.count = count;
    }
    users.items = data.users || [];
    users.items.forEach(fillUserAccountInfo);
  } finally {
    users.processing = false;
  }
}

// userListClear 清空用户记录
export function userListClear(): void {
  users.count = -1;
  users.items.length = 0;
}

// userListLogin 查询用户登录记录
export async function userListLogin(params: {
  account?: string;
  begin: string;
  end: string;
  limit: number;
  offset: number;
  order?: string;
}): Promise<void> {
  if (logins.processing) {
    return;
  }
  try {
    logins.processing = true;
    const { data } = await request.get(USERS_LOGINS, {
      params,
    });
    const count = data.count || 0;
    if (count >= 0) {
      logins.count = count;
    }
    logins.items = data.userLogins || [];
    logins.items.forEach(fillUserLoginRecord);
  } finally {
    logins.processing = false;
  }
}

// userLoginClear 清空登录记录
export function userLoginClear(): void {
  logins.count = -1;
  logins.items.length = 0;
}

// userUpdateByID 通过ID更新用户
export async function userUpdateByID(params: {
  id: number;
  data: Record<string, unknown>;
}): Promise<void> {
  await request.patch(USERS_ID.replace(":id", `${params.id}`), params.data);
}

// 仅读用户state
interface ReadonlyUserState {
  info: DeepReadonly<UserInfo>;
  users: DeepReadonly<UserAccounts>;
  logins: DeepReadonly<UserLoginRecords>;
}

const state = {
  info: readonly(info),
  users: readonly(users),
  logins: readonly(logins),
};

// useUserState 使用用户state
export default function useUserState(): ReadonlyUserState {
  return state;
}
