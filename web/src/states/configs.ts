import { DeepReadonly, reactive, readonly } from "vue";

import request from "../helpers/request";

import { CONFIGS, CONFIGS_ID, CONFIGS_CURRENT_VALID } from "../constants/url";

export enum ConfigCategory {
  MockTime = "mockTime",
  BlockIP = "blockIP",
  SignedKey = "signedKey",
  RouterConcurrency = "routerConcurrency",
  SessionInterceptor = "sessionInterceptor",
  RequestConcurrency = "requestConcurrency",
  Router = "router",
  Email = "email",
  HTTPServerInterceptor = "httpServerInterceptor",
}

export enum ConfigStatus {
  Enabled = 1,
  Disabled,
}

// 配置信息
export interface Config {
  key: string;
  id: number;
  createdAt: string;
  updatedAt: string;
  status: number;
  name: string;
  category: string;
  owner: string;
  data: string;
  startedAt: string;
  endedAt: string;
  description?: string;
}

interface Configs {
  processing: boolean;
  current?: Config;
  items: Config[];
  count: number;
}
const configs: Configs = reactive({
  processing: false,
  items: [],
  count: -1,
});

function fillConfigInfo(data: Config) {
  data.key = `${data.id}`;
}

// configAdd
export async function configAdd(params: {
  name: string;
  status: number;
  category: string;
  startedAt: string;
  endedAt: string;
  data: string;
}): Promise<Config> {
  const { data } = await request.post(CONFIGS, params);
  return <Config>data;
}

// 获取mock time的配置
export async function configGetMockTime(): Promise<Config> {
  const { data } = await request.get(CONFIGS, {
    params: {
      category: ConfigCategory.MockTime,
      name: ConfigCategory.MockTime,
      limit: 1,
    },
  });
  const items = data.configurations || [];
  if (items.length === 0) {
    return <Config>{};
  }
  return <Config>items[0];
}

// configFindByID 通过ID查询config
export async function configFindByID(id: number): Promise<Config> {
  const url = CONFIGS_ID.replace(":id", `${id}`);
  const { data } = await request.get(url);
  return <Config>data;
}

// configUpdateByID 通过ID更新config
export async function configUpdateByID(params: {
  id: number;
  data: Record<string, unknown>;
}): Promise<void> {
  const url = CONFIGS_ID.replace(":id", `${params.id}`);
  await request.patch(url, params.data);
}

// configList 查询配置列表
export async function configList(params: {
  name?: string;
  category?: string;
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (configs.processing) {
    return;
  }
  if (!params.limit) {
    params.limit = 50;
  }
  try {
    configs.processing = true;
    const { data } = await request.get(CONFIGS, {
      params,
    });
    const count = data.count || 0;
    if (count >= 0) {
      configs.count = count;
    }
    configs.items = data.configurations || [];
    configs.items.forEach(fillConfigInfo);
  } finally {
    configs.processing = false;
  }
}

export function configListClear(): void {
  configs.items = [];
  configs.count = -1;
}

export async function configGetCurrentValid(): Promise<
  Record<string, unknown>
> {
  const { data } = await request.get(CONFIGS_CURRENT_VALID);
  return data;
}

// 仅读配置state
interface ReadonlyConfigState {
  configs: DeepReadonly<Configs>;
}

const state = {
  configs: readonly(configs),
};
export default function useConfigState(): ReadonlyConfigState {
  return state;
}
