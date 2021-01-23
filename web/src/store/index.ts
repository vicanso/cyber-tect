import { Store } from "vuex";

import { getUserStore, userStore } from "./modules/user";
import { getCommonStore, commonStore } from "./modules/common";
import { getFluxStore, fluxStore } from "./modules/flux";
import { getConfigStore, configStore } from "./modules/config";
import { getDetectorStore, detectorStore } from "./modules/detector";

const stores: Store<any>[] = [
  userStore,
  commonStore,
  fluxStore,
  configStore,
  detectorStore,
];

export const useUserStore = getUserStore;
export const useCommonStore = getCommonStore;
export const useFluxStore = getFluxStore;
export const useConfigStore = getConfigStore;
export const useDetectorStore = getDetectorStore;

export default stores;
