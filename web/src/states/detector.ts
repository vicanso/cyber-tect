import { DeepReadonly, reactive, readonly } from "vue";

import request from "../helpers/request";

import {
  HTTP_DETECTORS,
  HTTP_DETECTORS_ID,
  DETECTOR_LIST_USER,
} from "../constants/url";

// http检测配置
export interface HTTPDetector {
  [key: string]: unknown;
  id: number;
  createdAt: string;
  updatedAt: string;
  status: string;
  statusDesc: string;
  name: string;
  owners: string[];
  receivers: string[];
  timeout: string;
  description: string;
  ips: string[];
  url: string;
}

interface HTTPDetectors {
  processing: boolean;
  items: HTTPDetector[];
  count: number;
}

const httpDetectors: HTTPDetectors = reactive({
  processing: false,
  items: [],
  count: -1,
});

// 查询http检测配置
export async function httpDetectorList(params: {
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (httpDetectors.processing) {
    return;
  }
  try {
    const { data } = await request.get(HTTP_DETECTORS, {
      params,
    });
    const count = data.count || 0;
    if (count >= 0) {
      httpDetectors.count = count;
    }
    httpDetectors.items = data.httpDetectors || [];
  } finally {
    httpDetectors.processing = false;
  }
}

export async function detectorListUser(keyword: string): Promise<string[]> {
  const { data } = await request.get(DETECTOR_LIST_USER, {
    params: {
      keyword,
      limit: 20,
    },
  });
  return data.accounts || [];
}

export async function httpDetectorFindByID(id: number): Promise<HTTPDetector> {
  const { data } = await request.get(
    HTTP_DETECTORS_ID.replace(":id", id.toString())
  );
  return data as HTTPDetector;
}

export async function httpDetectorUpdateByID(
  id: number,
  updateData: Record<string, unknown>
): Promise<HTTPDetector> {
  const { data } = await request.patch(
    HTTP_DETECTORS_ID.replace(":id", id.toString()),
    updateData
  );
  return data as HTTPDetector;
}

export async function httpDetectorCreate(
  createData: Record<string, unknown>
): Promise<HTTPDetector> {
  const { data } = await request.post(HTTP_DETECTORS, createData);
  return data as HTTPDetector;
}

interface ReadonlyDetectorState {
  httpDetectors: DeepReadonly<HTTPDetectors>;
}

const state = {
  httpDetectors: readonly(httpDetectors),
};

export default function useDetectorState(): ReadonlyDetectorState {
  return state;
}
