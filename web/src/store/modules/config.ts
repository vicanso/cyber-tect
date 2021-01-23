import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";

import request from "../../helpers/request";
import { CONFIG_ENABLED, CONFIG_DISABLED } from "../../constants/common";
import {
  CONFIGS,
  CONFIGS_ID,
  CONFIGS_CURRENT_VALID,
} from "../../constants/url";

const prefix = "config";

const mutationConfigProcessing = `${prefix}.processing`;
const mutationConfigDone = `${prefix}.done`;

const prefixList = `${prefix}.list`;
const mutationConfigListProcessing = `${prefixList}.processing`;
const mutationConfigList = `${prefixList}.list`;
const mutationConfigListReset = `${prefixList}.reset`;

const prefixCurrentValid = `${prefix}.currentValid`;
const mutationConfigCurrentValidProcessing = `${prefixCurrentValid}.processing`;
const mutationConfigCurrentValidDone = `${prefixCurrentValid}.done`;

interface ConfigList {
  processing: boolean;
  count: number;
  items: any[];
}
interface ConfigInfo {
  processing: boolean;
  data: any;
}
interface ConfigCurrentValid {
  processing: boolean;
  data: string;
}

interface ConfigState {
  configs: ConfigList;
  statuses: any[];
  current: ConfigInfo;
  currentValid: ConfigCurrentValid;
}

const configs: ConfigList = {
  processing: false,
  count: -1,
  items: [],
};

const current: ConfigInfo = {
  processing: false,
  data: null,
};

const currentValid: ConfigCurrentValid = {
  processing: false,
  data: "",
};

const state: ConfigState = {
  configs: configs,
  statuses: [
    {
      label: "启用",
      value: CONFIG_ENABLED,
    },
    {
      label: "禁用",
      value: CONFIG_DISABLED,
    },
  ],
  current,
  currentValid,
};

export const configStore = createStore<ConfigState>({
  state,
  mutations: {
    // 设置正在获取配置列表
    [mutationConfigListProcessing](state: ConfigState, processing: boolean) {
      state.configs.processing = processing;
    },
    // 重置列表数据
    [mutationConfigListReset](state: ConfigState) {
      state.configs.count = -1;
      state.configs.items = [];
    },
    // 设置配置信息
    [mutationConfigList](
      state: ConfigState,
      data: { count: number; configurations: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.configs.count = count;
      }
      state.configs.items = data.configurations || [];
    },
    // 设置当前配置处理中
    [mutationConfigProcessing](state: ConfigState, processing: boolean) {
      state.current.processing = processing;
    },
    // 设置当前配置
    [mutationConfigDone](state: ConfigState, data) {
      state.current.data = data;
    },
    // 设置获取当前配置中
    [mutationConfigCurrentValidProcessing](
      state: ConfigState,
      processing: boolean
    ) {
      state.currentValid.processing = processing;
    },
    // 设置当前配置
    [mutationConfigCurrentValidDone](state: ConfigState, data: any) {
      state.currentValid.data = JSON.stringify(data, null, 2);
    },
  },
  actions: {
    // add 添加配置
    async add(context: { commit: Commit }, config) {
      context.commit(mutationConfigProcessing, true);
      try {
        const { data } = await request.post(CONFIGS, config);
        context.commit(mutationConfigDone, data);
      } finally {
        context.commit(mutationConfigProcessing, false);
      }
    },
    // list 查询配置
    async list(context: { commit: Commit }, params: any) {
      context.commit(mutationConfigListProcessing, true);
      if (!params.offset) {
        context.commit(mutationConfigListReset);
      }
      if (!params.limit) {
        params.limit = 50;
      }
      try {
        const { data } = await request.get(CONFIGS, {
          params,
        });
        context.commit(mutationConfigList, data);
        return data;
      } finally {
        context.commit(mutationConfigListProcessing, false);
      }
    },
    // findByID 通过ID查询配置
    async findByID(context: { commit: Commit }, id: number) {
      context.commit(mutationConfigProcessing, true);
      try {
        const url = CONFIGS_ID.replace(":id", `${id}`);
        const { data } = await request.get(url);
        context.commit(mutationConfigDone, data);
        return data;
      } finally {
        context.commit(mutationConfigProcessing, false);
      }
    },
    // updateByID 通过ID更新配置
    async updateByID(
      context: { commit: Commit },
      params: { id: number; data: any }
    ) {
      context.commit(mutationConfigProcessing, true);
      try {
        const url = CONFIGS_ID.replace(":id", `${params.id}`);
        const { data } = await request.patch(url, params.data);
        context.commit(mutationConfigDone, data);
      } finally {
        context.commit(mutationConfigProcessing, false);
      }
    },
    // getCurrentValid 获取当前有效的配置
    async getCurrentValid(context: { commit: Commit }) {
      context.commit(mutationConfigCurrentValidProcessing, true);
      try {
        const { data } = await request.get(CONFIGS_CURRENT_VALID);
        context.commit(mutationConfigCurrentValidDone, data);
      } finally {
        context.commit(mutationConfigCurrentValidProcessing, false);
      }
    },
  },
});

export function getConfigStore(): Store<ConfigState> {
  return configStore;
}
