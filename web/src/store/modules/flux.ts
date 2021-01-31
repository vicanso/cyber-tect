import { createStore, Store } from "vuex";
import { Commit } from "vuex/types";
import request from "../../helpers/request";
import {
  FLUXES_TRACKERS,
  FLUXES_HTTP_ERRORS,
  FLUXES_TAG_VALUES,
  FLUXES_ACTIONS,
} from "../../constants/url";

const prefix = "flux";

const prefixTracker = `${prefix}.tracker`;
const mutationTrackerListProcessing = `${prefixTracker}.processing`;
const mutationTrackerList = `${prefixTracker}.list`;

const prefixHTTPError = `${prefix}.httpError`;
const mutationHTTPErrorListProcessing = `${prefixHTTPError}.processing`;
const mutationHTTPErrorList = `${prefixHTTPError}.list`;

const prefixTrackerAction = `${prefix}.trackerAction`;
const mutationTrackerActionListProcessing = `${prefixTrackerAction}.processing`;
const mutationTrackerActionList = `${prefixTrackerAction}.list`;

const prefixHTTPErrorCategory = `${prefix}.httpError.category`;
const mutationHTTPErrorCategoryListProcessing = `${prefixHTTPErrorCategory}.processing`;
const mutationHTTPErrorCategoryList = `${prefixHTTPErrorCategory}.list`;

// 客户端上传的user action类型
const prefixUserActionCategory = `${prefix}.userActionCategory`;
const mutationUserActionCategoryListProcessing = `${prefixUserActionCategory}.processing`;
const mutationUserActionCategoryList = `${prefixUserActionCategory}.list`;

const prefixUserAction = `${prefix}.userAction`;
const mutationUserActionListProcessing = `${prefixUserAction}.processing`;
const mutationUserActionList = `${prefixUserAction}.list`;

interface TrackerList {
  processing: boolean;
  items: any[];
}
interface HTTPErrorList {
  processing: boolean;
  items: any[];
}
interface TrackerActionList {
  processing: boolean;
  items: any[];
}
interface HTTPErrorCategoryList {
  processing: boolean;
  items: any[];
}
interface UserActionCategoryList {
  processing: boolean;
  items: any[];
}
interface UserActionList {
  processing: boolean;
  items: any[];
}
interface FluxState {
  trackers: TrackerList;
  httpErrors: HTTPErrorList;
  trackerActions: TrackerActionList;
  httpErrorCategories: HTTPErrorCategoryList;
  userActionCategories: UserActionCategoryList;
  userActions: UserActionList;
}

const trackers: TrackerList = {
  processing: false,
  items: [],
};

const httpErrors: HTTPErrorList = {
  processing: false,
  items: [],
};

const trackerActions: TrackerActionList = {
  processing: false,
  items: [],
};
const httpErrorCategories: HTTPErrorCategoryList = {
  processing: false,
  items: [],
};
const userActionCategories: UserActionCategoryList = {
  processing: false,
  items: [],
};
const userActions: UserActionList = {
  processing: false,
  items: [],
};

const state: FluxState = {
  trackers,
  httpErrors,
  trackerActions,
  httpErrorCategories,
  userActionCategories,
  userActions,
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
    [mutationTrackerActionListProcessing](
      state: FluxState,
      processing: boolean
    ) {
      state.trackerActions.processing = processing;
    },
    // 设置action列表
    [mutationTrackerActionList](state: FluxState, data: { values: any[] }) {
      state.trackerActions.items = data.values || [];
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
    // 设置正在查询分类
    [mutationUserActionCategoryListProcessing](
      state: FluxState,
      processing: boolean
    ) {
      state.userActionCategories.processing = processing;
    },
    // 设置user action分类
    [mutationUserActionCategoryList](
      state: FluxState,
      data: { values: any[] }
    ) {
      state.userActionCategories.items = data.values || [];
    },
    // 设置正在查询客户端行为
    [mutationUserActionListProcessing](state: FluxState, processing: boolean) {
      state.userActions.processing = processing;
    },
    // 设置客户端行为记录
    [mutationUserActionList](state: FluxState, data: { actions: any[] }) {
      state.userActions.items = data.actions || [];
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
    async listTrackerActions(context: { commit: Commit }) {
      if (state.trackerActions.items?.length !== 0) {
        return;
      }
      context.commit(mutationTrackerActionListProcessing, true);
      try {
        const url = FLUXES_TAG_VALUES.replace(
          ":measurement",
          "userTracker"
        ).replace(":tag", "action");
        const { data } = await request.get(url);
        context.commit(mutationTrackerActionList, data);
      } finally {
        context.commit(mutationTrackerActionListProcessing, false);
      }
    },
    async listHTTPErrorCategory(context: { commit: Commit }) {
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
    async listUserActionCategory(context: { commit: Commit }) {
      if (state.userActionCategories.items?.length !== 0) {
        return;
      }
      context.commit(mutationUserActionCategoryListProcessing, true);
      try {
        const url = FLUXES_TAG_VALUES.replace(
          ":measurement",
          "userAction"
        ).replace(":tag", "category");
        const { data } = await request.get(url);
        context.commit(mutationUserActionCategoryList, data);
      } finally {
        context.commit(mutationUserActionCategoryListProcessing, false);
      }
    },
    async listAction(context: { commit: Commit }, params) {
      context.commit(mutationUserActionListProcessing, true);
      try {
        const { data } = await request.get(FLUXES_ACTIONS, {
          params,
        });
        context.commit(mutationUserActionList, data);
      } finally {
        context.commit(mutationUserActionListProcessing, false);
      }
    },
  },
});

// getFluxStore 获取flux store
export function getFluxStore(): Store<FluxState> {
  return fluxStore;
}
