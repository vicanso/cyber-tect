import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";

import request from "../../helpers/request";
import {
  DETECTORS_HTTPS_RESULTS,
  DETECTORS_HTTPS_RESULTS_DETAIL,
  DETECTORS_DNSES_RESULTS,
  DETECTORS_DNSES_RESULTS_DEATIL,
  DETECTORS_TCPS_RESULTS,
  DETECTORS_TCPS_RESULTS_DETAIL,
  DETECTORS_PINGS_RESULTS,
  DETECTORS_PINGS_RESULTS_DETAIL,
} from "../../constants/url";

const prefix = "detectorResult";

const prefixHTTP = `${prefix}.http`;
const mutationHTTPListProcessing = `${prefixHTTP}.processing`;
const mutationHTTPList = `${prefixHTTP}.list`;

const prefixDNS = `${prefix}.dns`;
const mutationDNSListProcessing = `${prefixDNS}.processing`;
const mutationDNSList = `${prefixDNS}.list`;

const prefixTCP = `${prefix}.tcp`;
const mutationTCPListProcessing = `${prefixTCP}.processing`;
const mutationTCPList = `${prefixTCP}.list`;

const prefixPing = `${prefix}.ping`;
const mutationPingListProcessing = `${prefixPing}.processing`;
const mutationPingList = `${prefixPing}.list`;

interface DetectorResultList {
  processing: boolean;
  count: number;
  items: any[];
}

interface DetectorResultState {
  https: DetectorResultList;
  dnses: DetectorResultList;
  tcps: DetectorResultList;
  pings: DetectorResultList;
}

const https: DetectorResultList = {
  processing: false,
  count: -1,
  items: [],
};
const dnses: DetectorResultList = {
  processing: false,
  count: -1,
  items: [],
};
const tcps: DetectorResultList = {
  processing: false,
  count: -1,
  items: [],
};
const pings: DetectorResultList = {
  processing: false,
  count: -1,
  items: [],
};

const state: DetectorResultState = {
  https,
  dnses,
  tcps,
  pings,
};

function fillName(items: any[], detectors: any[]) {
  if (!detectors || detectors.length === 0) {
    return;
  }
  const names: any = {};
  detectors.forEach((detector) => {
    names[detector.id] = detector.name;
  });
  items.forEach((item) => {
    item.name = names[item.task];
  });
}

export const detectorResultStore = createStore<DetectorResultState>({
  state,
  mutations: {
    // 设置正在获取http检测结果
    [mutationHTTPListProcessing](
      state: DetectorResultState,
      processing: boolean
    ) {
      state.https.processing = processing;
    },
    // 设置http检测列表
    [mutationHTTPList](
      state: DetectorResultState,
      data: { count: number; httpResults: any[]; httpDetectors: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.https.count = count;
      }
      const items = data.httpResults || [];
      fillName(items, data.httpDetectors);
      state.https.items = items;
    },
    // 设置正在获取dns检测结果
    [mutationDNSListProcessing](
      state: DetectorResultState,
      processing: boolean
    ) {
      state.dnses.processing = processing;
    },
    // 设置dns检测列表
    [mutationDNSList](
      state: DetectorResultState,
      data: { count: number; dnsResults: any[]; dnsDetectors: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.dnses.count = count;
      }
      const items = data.dnsResults || [];
      fillName(items, data.dnsDetectors);
      state.dnses.items = items;
    },
    // 设置正在获取tcp检测结果
    [mutationTCPListProcessing](
      state: DetectorResultState,
      processing: boolean
    ) {
      state.tcps.processing = processing;
    },
    // 设置tcp检测结果列表
    [mutationTCPList](
      state: DetectorResultState,
      data: { count: number; tcpResults: any[]; tcpDetectors: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.tcps.count = count;
      }
      const items = data.tcpResults || [];
      fillName(items, data.tcpDetectors);
      state.tcps.items = items;
    },
    // 设置正在获取ping检测结果
    [mutationPingListProcessing](
      state: DetectorResultState,
      processing: boolean
    ) {
      state.pings.processing = processing;
    },
    // 设置ping检测结果列表
    [mutationPingList](
      state: DetectorResultState,
      data: { count: number; pingResults: any[]; pingDetectors: any[] }
    ) {
      const count = data.count || 0;
      if (count >= 0) {
        state.pings.count = count;
      }
      const items = data.pingResults || [];
      fillName(items, data.pingDetectors);
      state.pings.items = items;
    },
  },
  actions: {
    // listHTTP 查询http检测结果
    async listHTTP(context: { commit: Commit }, params: any) {
      context.commit(mutationHTTPListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_HTTPS_RESULTS, {
          params,
        });
        context.commit(mutationHTTPList, data);
      } finally {
        context.commit(mutationHTTPListProcessing, false);
      }
    },
    // getHTTP 查询http检测详情
    async getHTTP(context: { commit: Commit }, id) {
      context.commit(mutationHTTPListProcessing, true);
      try {
        const url = DETECTORS_HTTPS_RESULTS_DETAIL.replace(":id", id);
        const { data } = await request.get(url);
        return data;
      } finally {
        context.commit(mutationHTTPListProcessing, false);
      }
    },
    // listDNS 查询dns检测结果
    async listDNS(context: { commit: Commit }, params: any) {
      context.commit(mutationDNSListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_DNSES_RESULTS, {
          params,
        });
        context.commit(mutationDNSList, data);
      } finally {
        context.commit(mutationDNSListProcessing, false);
      }
    },
    // getDNS 查询dns检测结果详情
    async getDNS(context: { commit: Commit }, id) {
      context.commit(mutationDNSListProcessing, true);
      try {
        const url = DETECTORS_DNSES_RESULTS_DEATIL.replace(":id", id);
        const { data } = await request.get(url);
        return data;
      } finally {
        context.commit(mutationDNSListProcessing, false);
      }
    },
    // listTCP 查询tcp检测结果
    async listTCP(context: { commit: Commit }, params: any) {
      context.commit(mutationTCPListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_TCPS_RESULTS, {
          params,
        });
        context.commit(mutationTCPList, data);
      } finally {
        context.commit(mutationTCPListProcessing, false);
      }
    },
    // getTCP 查询tcp检测结果详情
    async getTCP(context: { commit: Commit }, id) {
      context.commit(mutationTCPListProcessing, true);
      try {
        const url = DETECTORS_TCPS_RESULTS_DETAIL.replace(":id", id);
        const { data } = await request.get(url);
        return data;
      } finally {
        context.commit(mutationTCPListProcessing, true);
      }
    },
    // listPing 查询ping检测结果
    async listPing(context: { commit: Commit }, params: any) {
      context.commit(mutationPingListProcessing, true);
      try {
        const { data } = await request.get(DETECTORS_PINGS_RESULTS, {
          params,
        });
        context.commit(mutationPingList, data);
      } finally {
        context.commit(mutationPingListProcessing, false);
      }
    },
    // getPing 获取ping检测结果详情
    async getPing(context: { commit: Commit }, id) {
      context.commit(mutationPingListProcessing, true);
      try {
        const url = DETECTORS_PINGS_RESULTS_DETAIL.replace(":id", id);
        const { data } = await request.get(url);
        return data;
      } finally {
        context.commit(mutationPingListProcessing, false);
      }
    },
  },
});

// getDetectorResultStore get detector result store
export function getDetectorResultStore(): Store<DetectorResultState> {
  return detectorResultStore;
}
