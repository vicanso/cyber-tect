import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";
import request from "../../helpers/request";
import {
  FLUXES_TRACKERS,
  FLUXES_HTTP_ERRORS,
  FLUXES_TAG_VALUES,
} from "../../constants/url";

const prefix = "flux";

const prefixTracker = `${prefix}.tracker`;
const mutationTrackerListProcessing = `${prefixTracker}.processing`;
const mutationTrackerList = `${prefixTracker}.list`;

const prefixHTTPError = `${prefix}.httpError`;
const mutationHTTPErrorListProcessing = `${prefixHTTPError}.processing`;
const mutationHTTPErrorList = `${prefixHTTPError}.list`;

const prefixAction = `${prefix}.action`;
const mutationActionListProcessing = `${prefixAction}.processing`;
const mutationActionList = `${prefixAction}.list`;

const prefixHTTPErrorCategory = `${prefix}.httpError.category`;
const mutationHTTPErrorCategoryListProcessing = `${prefixHTTPErrorCategory}.processing`;
const mutationHTTPErrorCategoryList = `${prefixHTTPErrorCategory}.list`;

interface TrackerList {
  processing: boolean;
  items: any[];
}
interface HTTPErrorList {
  processing: boolean;
  items: any[];
}
interface ActionList {
  processing: boolean;
  items: any[];
}
interface HTTPErrorCategoryList {
  processing: boolean;
  items: any[];
}

interface FluxState {
  trackers: TrackerList;
  httpErrors: HTTPErrorList;
  actions: ActionList;
  httpErrorCategories: HTTPErrorCategoryList;
}

const trackers: TrackerList = {
  processing: false,
  items: [],
};

const httpErrors: HTTPErrorList = {
  processing: false,
  items: [],
};

const actions: ActionList = {
  processing: false,
  items: [],
};
const httpErrorCategories: HTTPErrorCategoryList = {
  processing: false,
  items: [],
};

const state: FluxState = {
  trackers,
  httpErrors,
  actions,
  httpErrorCategories,
};

function fluxItemsSort(items: any[]) {
  return (items || [])
    .sort((item1, item2) => {
      if (item1._time < item2._time) {
        return -1;
      }
      if (item1._time > item2._time) {
        return 1;
      }
      return 0;
    })
    .reverse();
}

export const fluxStore = createStore<FluxState>({
  state,
  mutations: {
    // 设置正在查询用户行为
    [mutationTrackerListProcessing](state: FluxState, processing: boolean) {
      state.trackers.processing = processing;
    },
    // 设置用户行为记录
    [mutationTrackerList](state: FluxState, data: { trackers: any[] }) {
      state.trackers.items = fluxItemsSort(data.trackers);
    },
    // 设置正在查询出错列表
    [mutationHTTPErrorListProcessing](state: FluxState, processing: boolean) {
      state.httpErrors.processing = processing;
    },
    // 设置出错列表
    [mutationHTTPErrorList](state: FluxState, data: { httpErrors: any[] }) {
      state.httpErrors.items = fluxItemsSort(data.httpErrors);
    },
    // 设置正在查询action列表
    [mutationActionListProcessing](state: FluxState, processing: boolean) {
      state.actions.processing = processing;
    },
    // 设置action列表
    [mutationActionList](state: FluxState, data: { values: any[] }) {
      state.actions.items = data.values || [];
    },
    // 设置正在查询http error category列表
    [mutationHTTPErrorCategoryListProcessing](
      state: FluxState,
      processing: boolean
    ) {
      state.httpErrorCategories.processing = processing;
    },
    // 设置http error category列表
    [mutationHTTPErrorCategoryList](state: FluxState, data: { values: any[] }) {
      state.httpErrorCategories.items = data.values || [];
    },
  },
  actions: {
    async listTracker(context: { commit: Commit }, params) {
      context.commit(mutationTrackerListProcessing, true);
      try {
        const { data } = await request.get(FLUXES_TRACKERS, {
          params,
        });
        context.commit(mutationTrackerList, data);
      } finally {
        context.commit(mutationTrackerListProcessing, false);
      }
    },
    async listHTTPError(context: { commit: Commit }, params) {
      context.commit(mutationHTTPErrorListProcessing, true);
      try {
        const { data } = await request.get(FLUXES_HTTP_ERRORS, {
          params,
        });
        context.commit(mutationHTTPErrorList, data);
      } finally {
        context.commit(mutationHTTPErrorListProcessing, false);
      }
    },
    async listActions(context: { commit: Commit }) {
      if (state.actions.items?.length !== 0) {
        return;
      }
      context.commit(mutationActionListProcessing, true);
      try {
        const url = FLUXES_TAG_VALUES.replace(
          ":measurement",
          "userTracker"
        ).replace(":tag", "action");
        const { data } = await request.get(url);
        context.commit(mutationActionList, data);
      } finally {
        context.commit(mutationActionListProcessing, false);
      }
    },
    async listHTTPErrorCategories(context: { commit: Commit }) {
      if (state.httpErrorCategories.items?.length !== 0) {
        return;
      }
      context.commit(mutationHTTPErrorCategoryListProcessing, true);
      try {
        const url = FLUXES_TAG_VALUES.replace(
          ":measurement",
          "httpError"
        ).replace(":tag", "category");
        const { data } = await request.get(url);
        context.commit(mutationHTTPErrorCategoryList, data);
      } finally {
        context.commit(mutationHTTPErrorCategoryListProcessing, false);
      }
    },
  },
});

// getFluxStore 获取flux store
export function getFluxStore(): Store<FluxState> {
  return fluxStore;
}
