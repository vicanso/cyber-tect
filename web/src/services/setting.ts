// 用户的配置

import store from "../helpers/store";

const userSettingKey = "userSettings";

interface UserSetting {
  // 主侧边栏是否隐藏
  mainNavShrinking: boolean;
}

let currentUserSetting: UserSetting = {
  mainNavShrinking: false,
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
