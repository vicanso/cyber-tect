// 用户的配置

import store from "../helpers/store";

const userSettingKey = "userSettings";

interface UserSetting {
  // 主侧边栏是否隐藏
  mainNavShrinking: boolean;
  // mainDetectorResultCount 首页展示的检测结果数
  mainDetectorResultCount: number;
  // mainDetectorRefreshInterval 首页展示定时刷新间隔
  mainDetectorRefreshInterval: number;
  // mainDetectorOnlyFailure 首页仅展示失败记录
  mainDetectorOnlyFailure: boolean;
  // mainDetectorTimeRange 首页查询的时间区间
  mainDetectorTimeRange: string;
}

let currentUserSetting: UserSetting = {
  mainNavShrinking: false,
  mainDetectorResultCount: 100,
  mainDetectorRefreshInterval: 0,
  mainDetectorOnlyFailure: false,
  mainDetectorTimeRange: "",
};

export async function loadSetting(): void {
  const data = await store.getItem(userSettingKey);
  if (!data) {
    return;
  }
  currentUserSetting = JSON.parse(data);
}

export function getSetting(): UserSetting {
  return currentUserSetting;
}

export async function saveSetting(setting: UserSetting): void {
  await store.setItem(userSettingKey, JSON.stringify(setting));
  currentUserSetting = setting;
}
