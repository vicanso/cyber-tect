import store from "../helpers/store";
import request from "../helpers/request";
import { USERS_ACTIONS } from "../constants/url";

const userActionKey = "userActions";
const userActions: any[] = [];

// 定时flush的间隔
const flushInterval = 60 * 1000;
let timer: number;

// UserActionData 用户action信息
interface UserActionData {
  category: string;
  route: string;
  path: string;
  result: number;
  time: number;
}

async function loadFromStore() {
  const data = await store.getItem(userActionKey);
  if (!data) {
    return;
  }
  const arr = JSON.parse(data);
  userActions.push(...arr);
}
loadFromStore();

// 默认的key，除此之外的都添加至extra
const defaultKeys: string[] = ["category", "route", "path", "result", "time"];
async function flush() {
  store.removeItem(userActionKey);
  // 需要将actions转换
  const actions = userActions.slice(0).map((action) => {
    const extra: any = {};
    const newAction: any = {
      extra,
    };
    Object.keys(action).forEach((element) => {
      const value = action[element];
      if (defaultKeys.includes(element)) {
        newAction[element] = value;
      } else {
        extra[element] = value;
      }
    });
    return newAction;
  });
  console.dir(actions);

  userActions.length = 0;
  request.post(USERS_ACTIONS, {
    actions,
  });
}

// 成功
export const SUCCESS = 0;
// 失败
export const FAIL = 1;

// 分类
// 点击（默认分类）
export const CLICK = "click";
// 登录
export const LOGIN = "login";
// 注册
export const REGISTER = "register";

export function addUserAction(data: UserActionData): void {
  // 每次添加新的action时，清空定时器
  clearTimeout(timer);
  userActions.push(data);
  if (userActions.length < 10) {
    store.setItem(userActionKey, JSON.stringify(userActions));
    // 重新启动定时器
    timer = setTimeout(flush, flushInterval);
    return;
  }
  flush();
}
