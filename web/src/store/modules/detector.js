import request from "@/request";
import { DETECTORS, DETECTOR, DETECT_RESULTS } from "@/constants/url";
import { formatDuration, formatDate, delay } from "@/helpers/util";

const mutationDetectorProcessing = "detector.processing";
const mutationDetectorReset = "detector.reset";
const mutationDetectorList = "detector.list";
const mutationDetectorChangeCurrent = "detector.changeCurrent";

const mutationDetectorUpdate = "detector.update";
const mutationDetectorResultProcessing = "detector.result.processing";
const mutationDetectorResultList = "detector.result.list";
const mutationDetectorResultReset = "detector.result.reset";

const mutationDetectorTaskFilterProcessing = "detector.task.filter.processing";
const mutationDetectorTaskFilter = "detector.task.filter";
const mutationDetectorTaskFilterReset = "detector.task.filterReset";

const mutationDetectorMimeList = "detector.mime.list";
const mutationDetectorMimeProcessing = "detector.mime.processing";

const statusDescList = ["未知", "启用", "禁用"];
const resultDescList = ["未知", "成功", "失败"];

const state = {
  resultSucess: 1,
  resultFail: 2,
  processing: false,
  currentCategory: "",
  currentDetectors: null,
  updateDetector: null,
  // 我所创建的dtector
  mime: {
    http: {
      processing: false,
      detectors: null,
    },
    dns: {
      processing: false,
      detectors: null,
    },
    tcp: {
      processing: false,
      detectors: null,
    },
    ping: {
      processing: false,
      detectors: null,
    },
  },
  http: {
    count: -1,
    detectors: null,
  },
  dns: {
    count: -1,
    detectors: null,
  },
  ping: {
    count: -1,
    detectors: null,
  },
  tcp: {
    count: -1,
    detectors: null,
  },
  httpListResult: {
    count: -1,
    results: null,
    processing: false,
  },
  dnsListResult: {
    count: -1,
    results: null,
    processing: false,
  },
  pingListResult: {
    count: -1,
    results: null,
    processing: false,
  },
  tcpListResult: {
    count: -1,
    results: null,
    processing: false,
  },
  filterProcessing: false,
  filterTasks: {
    tasks: null,
  },
};
export default {
  state,
  mutations: {
    [mutationDetectorProcessing](state, value) {
      state.processing = value;
    },
    // 重置detector数据
    [mutationDetectorReset](state, category) {
      state[category].count = -1;
      state[category].detectors = null;
    },
    // 获取detector配置列表
    [mutationDetectorList](state, { category, data }) {
      state.currentCategory = category;
      const arr = (state[category].detectors || []).slice(0);
      if (data.count >= 0) {
        state[category].count = data.count;
      }
      state[category].detectors = arr.concat(data.detectors || []);
    },
    // 设置当前更新的detector
    [mutationDetectorChangeCurrent](state, { category, limit, offset }) {
      state.currentCategory = category;
      if (!limit) {
        state.currentDetectors = null;
      } else {
        state.currentDetectors = state[category].detectors.slice(
          offset,
          offset + limit
        );
      }
    },
    // 设置detector更新数据
    [mutationDetectorUpdate](state, data) {
      state.updateDetector = data;
    },
    // 设置是否正在获取detector结果
    [mutationDetectorResultProcessing](state, { category, processing }) {
      state[`${category}ListResult`].processing = processing;
    },
    // 重置detector结果列表
    [mutationDetectorResultReset](state, { category }) {
      const listResult = state[`${category}ListResult`];
      listResult.count = 0;
      listResult.results = null;
    },
    // 获取detector结果列表
    [mutationDetectorResultList](state, { category, data }) {
      const listResult = state[`${category}ListResult`];
      if (data.count >= 0) {
        listResult.count = data.count;
      }
      data.results.forEach((item) => {
        if (!item.addrs) {
          item.addrs = [];
        }
        if (item.addrs.length === 0 && item.addr) {
          item.addrs.push(item.addr);
        }
        if (item.updatedAt) {
          item.updatedAtDesc = formatDate(item.updatedAt);
        }
        if (item.result) {
          item.resultDesc = resultDescList[item.result];
        }
        item.durationDesc = formatDuration(item.duration);
      });
      listResult.results = data.results;
    },
    // 设置是否正在获取task filter
    [mutationDetectorTaskFilterProcessing](state, value) {
      state.filterProcessing = value;
    },
    // 重置task filter数据
    [mutationDetectorTaskFilterReset](state) {
      state.filterTasks.tasks = null;
    },
    // 设置task filter数据
    [mutationDetectorTaskFilter](state, data) {
      state.filterTasks.tasks = data;
    },
    // 设置是否正在获取我的detector
    [mutationDetectorMimeProcessing](state, { category, processing }) {
      state.mime[category].processing = processing;
    },
    // 设置我的detector数据
    [mutationDetectorMimeList](state, { category, data }) {
      state.mime[category].detectors = data;
    },
  },
  actions: {
    async addDetector({ commit }, { category, detector }) {
      commit(mutationDetectorProcessing, true);
      try {
        await request.post(DETECTORS.replace(":category", category), detector);
      } finally {
        commit(mutationDetectorProcessing, false);
      }
    },
    async listDetector({ commit }, { category, params }) {
      commit(mutationDetectorProcessing, true);
      // 如果没有指定offset或者为0，则重置
      if (params && !params.offset) {
        commit(mutationDetectorReset, category);
        commit(mutationDetectorChangeCurrent, {
          category,
          limit: 0,
        });
      }
      try {
        const { data } = await request.get(
          DETECTORS.replace(":category", category),
          {
            params,
          }
        );
        if (!data.count) {
          data.count = 0;
        }
        if (data.detectors) {
          data.detectors.forEach((item) => {
            item.statusDesc = statusDescList[item.status];
            item.updatedAtDesc = formatDate(item.updatedAt);
          });
        }
        // 保存完整的列表，暂时未使用
        commit(mutationDetectorList, {
          category,
          data,
        });
        commit(mutationDetectorChangeCurrent, {
          category,
          limit: params.limit,
          offset: params.offset,
        });
      } finally {
        commit(mutationDetectorProcessing, false);
      }
    },
    async updateDetector({ commit }, { id, category, data }) {
      const url = DETECTOR.replace(":category", category).replace(":id", id);
      commit(mutationDetectorProcessing, true);
      try {
        await request.patch(url, data);
      } finally {
        commit(mutationDetectorProcessing, false);
      }
    },
    async clearUpdateDetector({ commit }) {
      commit(mutationDetectorUpdate, null);
    },
    async getUpdateDetector({ commit }, { category, id }) {
      commit(mutationDetectorUpdate, null);
      commit(mutationDetectorProcessing, true);
      // 从缓存中获取
      if (state.currentCategory === category && state.currentDetectors) {
        let found = null;
        state.currentDetectors.forEach((item) => {
          if (item.id === id) {
            found = item;
          }
        });
        if (found) {
          await delay(30);
          commit(mutationDetectorUpdate, found);
          commit(mutationDetectorProcessing, false);
          return;
        }
      }

      const url = DETECTOR.replace(":category", category).replace(":id", id);
      try {
        const { data } = await request.get(url);
        commit(mutationDetectorUpdate, data);
      } finally {
        commit(mutationDetectorProcessing, false);
      }
    },
    async listMimeDetector({ commit }, { category, params }) {
      // 如果已加载过，直接返回
      if (state.mime[category].detectors) {
        return;
      }
      commit(mutationDetectorMimeProcessing, {
        category,
        processing: true,
      });
      try {
        const { data } = await request.get(
          DETECTORS.replace(":category", category),
          {
            params,
          }
        );
        commit(mutationDetectorMimeList, {
          category,
          data: data.detectors,
        });
      } finally {
        commit(mutationDetectorMimeProcessing, {
          category,
          processing: false,
        });
      }
    },
    async listDetectorResult({ commit }, { category, params }) {
      commit(mutationDetectorResultProcessing, {
        category,
        processing: true,
      });
      try {
        const { data } = await request.get(
          DETECT_RESULTS.replace(":category", category),
          {
            params,
          }
        );
        if (!data.count) {
          data.count = 0;
        }
        commit(mutationDetectorResultList, {
          category,
          data,
        });
      } finally {
        commit(mutationDetectorResultProcessing, {
          category,
          processing: false,
        });
      }
    },
    resetDetectorResults({ commit }, { category }) {
      commit(mutationDetectorResultReset, {
        category,
      });
    },
    async resetDetectorTaskFilter({ commit }) {
      commit(mutationDetectorTaskFilterReset);
    },
    async filterDetectorTask({ commit }, { category, params }) {
      commit(mutationDetectorTaskFilterProcessing, true);
      try {
        const { data } = await request.get(
          DETECTORS.replace(":category", category),
          {
            params,
          }
        );
        commit(mutationDetectorTaskFilter, data.detectors);
      } finally {
        commit(mutationDetectorTaskFilterProcessing, false);
      }
    },
  },
};
