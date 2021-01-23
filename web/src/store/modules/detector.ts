import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";

import request from "../../helpers/request";
import {
  DETECTORS_HTTPS,
  DETECTORS_RECEIVERS,
  DETECTORS_HTTPS_UPDATE,
} from "../../constants/url";

const prefix = "detector";

const prefixHTTP = `${prefix}.http`;
const mutationHTTPListProcessing = `${prefixHTTP}.processing`;
const mutationHTTPList = `${prefixHTTP}.list`;
const mutationHTTPAdd = `${prefixHTTP}.add`;
const mutationHTTPUpdate = `${prefixHTTP}.update`;

const prefixReceiver = `${prefix}.receiver`;
const mutationReceiverListProcessing = `${prefixReceiver}.processing`;
const mutationReceiverList = `${prefixReceiver}.list`;

interface HTTPList {
  processing: boolean;
  count: number;
  items: any[];
}

interface ReceiverList {
  processing: boolean;
  items: any[];
}

interface DetectorState {
  https: HTTPList;
  receivers: ReceiverList;
}

const https: HTTPList = {
  processing: false,
  count: -1,
  items: [],
};

const receivers: ReceiverList = {
  processing: false,
  items: [],
};

const state: DetectorState = {
  https,
  receivers,
};

export const detectorStore = createStore<DetectorState>({
  state,
  mutations: {
    // 设置正在获取http配置
    [mutationHTTPListProcessing](state: DetectorState, processing: boolean) {
      state.https.processing = processing;
    },
    // 设置http配置列表
    [mutationHTTPList](
      state: DetectorState,
      data: { count: number; https: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.https.count = count;
      }
      state.https.items = data.https || [];
    },
    // 添加新的http配置
    [mutationHTTPAdd](state: DetectorState, data: any) {
      const items = state.https.items.slice(0);
      items.unshift(data);
      state.https.items = items;
    },
    // 更新http配置
    [mutationHTTPUpdate](state: DetectorState, data: any) {
      const items = state.https.items.slice(0);
      items.forEach((element) => {
        if (element.id === data.id) {
          Object.assign(element, data);
        }
      });
      state.https.items = items;
    },
    // 设置正在获取接收者列表
    [mutationReceiverListProcessing](
      state: DetectorState,
      processing: boolean
    ) {
      state.receivers.processing = processing;
    },
    // 设置接收者列表
    [mutationReceiverList](state: DetectorState, data: { receivers: any[] }) {
      state.receivers.items = data.receivers || [];
    },
  },
  actions: {
    // listReceiver 获取接收者列表
    async listReceiver(context: { commit: Commit }) {
      if (state.receivers.items?.length) {
        return;
      }
      context.commit(mutationReceiverListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_RECEIVERS);
        context.commit(mutationReceiverList, data);
      } finally {
        context.commit(mutationReceiverListProcessing, false);
      }
    },
    // listHTTP 查询http配置
    async listHTTP(context: { commit: Commit }, params: any) {
      context.commit(mutationHTTPListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_HTTPS, {
          params,
        });
        context.commit(mutationHTTPList, data);
      } finally {
        context.commit(mutationHTTPListProcessing, false);
      }
    },
    // addHTTP 添加http配置
    async addHTTP(context: { commit: Commit }, params: any) {
      context.commit(mutationHTTPListProcessing, true);
      try {
        const { data } = await request.post(DETECTORS_HTTPS, params);
        context.commit(mutationHTTPAdd, data);
      } finally {
        context.commit(mutationHTTPListProcessing, false);
      }
    },
    // updateHTTPByID 通过ID更新http检测配置
    async updateHTTPByID(
      context: { commit: Commit },
      params: { id: number; data: any }
    ) {
      context.commit(mutationHTTPListProcessing, true);
      try {
        const url = DETECTORS_HTTPS_UPDATE.replace(":id", `${params.id}`);
        const { data } = await request.patch(url, params.data);
        context.commit(mutationHTTPUpdate, data);
      } finally {
        context.commit(mutationHTTPListProcessing, false);
      }
    },
  },
});

// getDetectorStore get detector store
export function getDetectorStore(): Store<DetectorState> {
  return detectorStore;
}
