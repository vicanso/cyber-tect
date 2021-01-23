import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";
import request from "../../helpers/request";
import {
  COMMONS_CAPTCHA,
  COMMONS_ROUTERS,
  COMMONS_STATUSES,
  COMMONS_RANDOM_KEYS,
} from "../../constants/url";

const prefix = "common";

const prefixStatus = `${prefix}.status`;
const mutationStatusListProcessing = `${prefixStatus}.processing`;
const mutationStatusList = `${prefixStatus}.list`;

const prefixRouter = `${prefix}.router`;
const mutationRouterListProcessing = `${prefixRouter}.processing`;
const mutationRouterList = `${prefixRouter}.list`;

const prefixRandomKey = `${prefix}.randomKey`;
const mutationRandomKeyProcessing = `${prefixRandomKey}.processing`;
const mutationRandomKeyDone = `${prefixRandomKey}.done`;

interface RouterInfo {
  processing: boolean;
}
interface CommonStatusList {
  processing: boolean;
  items: any[];
}
interface CommonRouterList {
  processing: boolean;
  items: any[];
}
interface CommonRandomKeyList {
  processing: boolean;
  items: any[];
}
interface CommonState {
  routerInfo: RouterInfo;
  statuses: CommonStatusList;
  routers: CommonRouterList;
  randomKeys: CommonRandomKeyList;
}
interface CaptchaData {
  data: string;
  expiredAt: string;
  id: string;
  type: string;
}

const routerInfo: RouterInfo = {
  processing: false,
};
const statuses: CommonStatusList = {
  processing: false,
  items: [],
};
const routers: CommonRouterList = {
  processing: false,
  items: [],
};
const randomKeys: CommonRandomKeyList = {
  processing: false,
  items: [],
};
const state: CommonState = {
  routerInfo,
  statuses,
  routers,
  randomKeys,
};

export const commonStore = createStore<CommonState>({
  state,
  mutations: {
    // 设置状态查询中
    [mutationStatusListProcessing](state: CommonState, processing: boolean) {
      state.statuses.processing = processing;
    },
    // 设置状态列表
    [mutationStatusList](state: CommonState, data: { statuses: any[] }) {
      state.statuses.items = data.statuses || [];
    },
    // 设置路由列表查询中
    [mutationRouterListProcessing](state: CommonState, processing: boolean) {
      state.routers.processing = processing;
    },
    // 设置路由列表
    [mutationRouterList](state: CommonState, data: { routers: any[] }) {
      state.routers.items = data.routers || [];
    },
    // 设置正在获取随机字符串
    [mutationRandomKeyProcessing](state: CommonState, processing: boolean) {
      state.randomKeys.processing = processing;
    },
    // 设置获取的随机字符串
    [mutationRandomKeyDone](state: CommonState, data: { keys: any[] }) {
      state.randomKeys.items = data.keys || [];
    },
  },
  actions: {
    // getCaptcha 获取图形验证码
    async getCaptcha() {
      const { data } = await request.get(COMMONS_CAPTCHA);
      return data as CaptchaData;
    },
    // listStatus 获取状态列表
    async listStatus(context: { commit: Commit }) {
      if (state.statuses.items.length !== 0) {
        return;
      }
      context.commit(mutationStatusListProcessing, true);
      try {
        const { data } = await request.get(COMMONS_STATUSES);
        context.commit(mutationStatusList, data);
      } finally {
        context.commit(mutationStatusListProcessing, false);
      }
    },
    // listRouter 获取路由列表
    async listRouter(context: { commit: Commit }) {
      if (state.routers.items.length !== 0) {
        return;
      }
      context.commit(mutationRouterListProcessing, true);
      try {
        const { data } = await request.get(COMMONS_ROUTERS);
        context.commit(mutationRouterList, data);
      } finally {
        context.commit(mutationRouterListProcessing, false);
      }
    },
    // listRandomKey 获取随机字符串
    async listRandomKey(context: { commit: Commit }) {
      context.commit(mutationRandomKeyProcessing, true);
      try {
        const { data } = await request.get(COMMONS_RANDOM_KEYS, {
          params: {
            size: 10,
            n: 5,
          },
        });
        context.commit(mutationRandomKeyDone, data);
      } finally {
        context.commit(mutationRandomKeyProcessing, false);
      }
    },
  },
});

// getCommonStore 获取公共store
export function getCommonStore(): Store<CommonState> {
  return commonStore;
}
