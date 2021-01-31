import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";

import request from "../../helpers/request";
import {
  DETECTORS_HTTPS,
  DETECTORS_RECEIVERS,
  DETECTORS_HTTPS_UPDATE,
  DETECTORS_DNSES,
  DETECTORS_DNSES_UPDATE,
  DETECTORS_TCPS,
  DETECTORS_TCPS_UPDATE,
  DETECTORS_PINGS,
  DETECTORS_PINGS_UPDATE,
} from "../../constants/url";

const prefix = "detector";

const prefixHTTP = `${prefix}.http`;
const mutationHTTPListProcessing = `${prefixHTTP}.processing`;
const mutationHTTPList = `${prefixHTTP}.list`;
const mutationHTTPAdd = `${prefixHTTP}.add`;
const mutationHTTPUpdate = `${prefixHTTP}.update`;

const prefixDNS = `${prefix}.dns`;
const mutationDNSListProcessing = `${prefixDNS}.processing`;
const mutationDNSList = `${prefixDNS}.list`;
const mutationDNSAdd = `${prefixDNS}.add`;
const mutationDNSUpdate = `${prefixDNS}.update`;

const prefixTCP = `${prefix}.tcp`;
const mutationTCPListProcessing = `${prefixTCP}.processing`;
const mutationTCPList = `${prefixTCP}.list`;
const mutationTCPAdd = `${prefixTCP}.add`;
const mutationTCPUpdate = `${prefixTCP}.update`;

const prefixPing = `${prefix}.ping`;
const mutationPingListProcessing = `${prefixPing}.processing`;
const mutationPingList = `${prefixPing}.list`;
const mutationPingAdd = `${prefixPing}.add`;
const mutationPingUpdate = `${prefixPing}.update`;

const prefixReceiver = `${prefix}.receiver`;
const mutationReceiverListProcessing = `${prefixReceiver}.processing`;
const mutationReceiverList = `${prefixReceiver}.list`;

interface DetectorList {
  processing: boolean;
  count: number;
  items: any[];
}

interface ReceiverList {
  processing: boolean;
  items: any[];
}

interface DetectorState {
  https: DetectorList;
  receivers: ReceiverList;
  dnses: DetectorList;
  tcps: DetectorList;
  pings: DetectorList;
}

const https: DetectorList = {
  processing: false,
  count: -1,
  items: [],
};
const dnses: DetectorList = {
  processing: false,
  count: -1,
  items: [],
};
const tcps: DetectorList = {
  processing: false,
  count: -1,
  items: [],
};
const pings: DetectorList = {
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
  dnses,
  tcps,
  pings,
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
    // 设置正在获取dns配置
    [mutationDNSListProcessing](state: DetectorState, processing: boolean) {
      state.dnses.processing = processing;
    },
    // 设置dns配置列表
    [mutationDNSList](
      state: DetectorState,
      data: { count: number; dnses: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.dnses.count = count;
      }
      state.dnses.items = data.dnses || [];
    },
    // 添加dns配置
    [mutationDNSAdd](state: DetectorState, data: any) {
      const items = state.dnses.items.slice(0);
      items.unshift(data);
      state.dnses.items = items;
    },
    // 更新dns配置
    [mutationDNSUpdate](state: DetectorState, data: any) {
      const items = state.dnses.items.slice(0);
      items.forEach((element) => {
        if (element.id === data.id) {
          Object.assign(element, data);
        }
      });
      state.dnses.items = items;
    },
    // 设置正在获取tcp状态
    [mutationTCPListProcessing](state: DetectorState, processing: boolean) {
      state.tcps.processing = processing;
    },
    // 设置tcp配置列表
    [mutationTCPList](
      state: DetectorState,
      data: { count: number; tcps: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.tcps.count = count;
      }
      state.tcps.items = data.tcps || [];
    },
    // 添加tcp配置
    [mutationTCPAdd](state: DetectorState, data: any) {
      const items = state.tcps.items.slice(0);
      items.unshift(data);
      state.tcps.items = items;
    },
    // 更新tcp配置
    [mutationTCPUpdate](state: DetectorState, data: any) {
      const items = state.tcps.items.slice(0);
      items.forEach((element) => {
        if (element.id === data.id) {
          Object.assign(element, data);
        }
      });
      state.tcps.items = items;
    },
    // 设置正在获取ping状态
    [mutationPingListProcessing](state: DetectorState, processing: boolean) {
      state.pings.processing = processing;
    },
    // 添加ping配置
    [mutationPingAdd](state: DetectorState, data: any) {
      const items = state.pings.items.slice(0);
      items.unshift(data);
      state.pings.items = items;
    },
    // 设置ping配置列表
    [mutationPingList](
      state: DetectorState,
      data: { count: number; pings: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.pings.count = count;
      }
      state.pings.items = data.pings || [];
    },
    // 更新ping配置
    [mutationPingUpdate](state: DetectorState, data: any) {
      const items = state.pings.items.slice(0);
      items.forEach((element) => {
        if (element.id === data.id) {
          Object.assign(element, data);
        }
      });
      state.pings.items = items;
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
    // listDNS 查询dns配置
    async listDNS(context: { commit: Commit }, params: any) {
      context.commit(mutationDNSListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_DNSES, {
          params,
        });
        context.commit(mutationDNSList, data);
      } finally {
        context.commit(mutationDNSListProcessing, false);
      }
    },
    // addDNS 添加dns配置
    async addDNS(context: { commit: Commit }, params: any) {
      context.commit(mutationDNSListProcessing, true);
      try {
        const { data } = await request.post(DETECTORS_DNSES, params);
        context.commit(mutationDNSAdd, data);
      } finally {
        context.commit(mutationDNSListProcessing, false);
      }
    },
    // updateDNSByID 通过ID更新dns检测配置
    async updateDNSByID(
      context: { commit: Commit },
      params: { id: number; data: any }
    ) {
      context.commit(mutationDNSListProcessing, true);
      try {
        const url = DETECTORS_DNSES_UPDATE.replace(":id", `${params.id}`);
        const { data } = await request.patch(url, params.data);
        context.commit(mutationDNSUpdate, data);
      } finally {
        context.commit(mutationDNSListProcessing, false);
      }
    },
    // listTCP 查询tcp配置
    async listTCP(context: { commit: Commit }, params: any) {
      context.commit(mutationTCPListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_TCPS, {
          params,
        });
        context.commit(mutationTCPList, data);
      } finally {
        context.commit(mutationTCPListProcessing, false);
      }
    },
    // addTCP 添加tcp配置
    async addTCP(context: { commit: Commit }, params: any) {
      context.commit(mutationTCPListProcessing, true);
      try {
        const { data } = await request.post(DETECTORS_TCPS, params);
        context.commit(mutationTCPAdd, data);
      } finally {
        context.commit(mutationTCPListProcessing, false);
      }
    },
    // updateTCPByID 更新tcp配置
    async updateTCPByID(
      context: { commit: Commit },
      params: { id: number; data: any }
    ) {
      context.commit(mutationTCPListProcessing, true);
      try {
        const url = DETECTORS_TCPS_UPDATE.replace(":id", `${params.id}`);
        const { data } = await request.patch(url, params.data);
        context.commit(mutationTCPUpdate, data);
      } finally {
        context.commit(mutationTCPListProcessing, false);
      }
    },
    // listPing 查询ping配置
    async listPing(context: { commit: Commit }, params: any) {
      context.commit(mutationPingListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_PINGS, {
          params,
        });
        context.commit(mutationPingList, data);
      } finally {
        context.commit(mutationPingListProcessing, false);
      }
    },
    // addPing 添加ping配置
    async addPing(context: { commit: Commit }, params: any) {
      context.commit(mutationPingListProcessing, true);
      try {
        const { data } = await request.post(DETECTORS_PINGS, params);
        context.commit(mutationPingAdd, data);
      } finally {
        context.commit(mutationPingListProcessing, false);
      }
    },
    // updatePingByID 更新ping配置
    async updatePingByID(
      context: { commit: Commit },
      params: { id: number; data: any }
    ) {
      context.commit(mutationPingListProcessing, true);
      try {
        const url = DETECTORS_PINGS_UPDATE.replace(":id", `${params.id}`);
        const { data } = await request.patch(url, params.data);
        context.commit(mutationPingUpdate, data);
      } finally {
        context.commit(mutationPingListProcessing, false);
      }
    },
  },
});

// getDetectorStore get detector store
export function getDetectorStore(): Store<DetectorState> {
  return detectorStore;
}
