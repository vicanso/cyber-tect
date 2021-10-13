import { DeepReadonly, reactive, readonly } from "vue";

import request from "../helpers/request";

import {
  HTTP_DETECTORS,
  HTTP_DETECTORS_ID,
  DETECTOR_LIST_USER,
  DNS_DETECTORS,
  DNS_DETECTORS_ID,
  TCP_DETECTORS,
  PING_DETECTORS,
  TCP_DETECTORS_ID,
  PING_DETECTORS_ID,
  HTTP_DETECTOR_RESULTS,
  TCP_DETECTOR_RESULTS,
  PING_DETECTOR_RESULTS,
  DNS_DETECTOR_RESULTS,
  DATABASE_DETECTORS,
  DATABASE_DETECTORS_ID,
  DATABASE_DETECTOR_RESULTS,
} from "../constants/url";

// http检测配置
interface HTTPDetector {
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

// dns 检测配置
interface DNSDetector {
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
  host: string;
  ips: string[];
  servers: string[];
}

// tcp 检测配置
interface TCPDetector {
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
  addrs: string[];
}

// ping 检测配置
interface PingDetector {
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
}

interface DatabaseDetector {
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
  uris: string[];
}

interface DetectorResult {
  desc: string;
  value: number;
}
interface HTTPDetectorSubResult {
  [key: string]: unknown;
  result: DetectorResult;
  addrs: string[];
  addr: string;
  protocol: string;
  tlsVersion: string;
  tlsCipherSuite: string;
  certificateDNSNames: string[];
  certificateExpirationDates: string[];
  dnsLookup: number;
  tcpConnection: number;
  tlsHandshake: number;
  serverProcessing: number;
  contentTransfer: number;
  duration: number;
  message: string;
  timeline: string[];
}
interface HTTPDetectorResult {
  [key: string]: unknown;
  id: number;
  createdAt: string;
  updatedAt: string;
  task: number;
  result: DetectorResult;
  maxDuration: number;
  messages: string[];
  // 检测url
  url: string;
  results: HTTPDetectorSubResult[];
}

interface TCPDetectorSubResult {
  [key: string]: unknown;
  result: DetectorResult;
  addr: string;
  duration: number;
  message: string;
}
interface TCPDetectorResult {
  [key: string]: unknown;
  id: number;
  createdAt: string;
  updatedAt: string;
  task: number;
  result: DetectorResult;
  maxDuration: number;
  messages: string[];
  // 检测地址
  addrs: string[];
  results: TCPDetectorSubResult[];
}

interface PingDetectorSubResult {
  [key: string]: unknown;
  result: DetectorResult;
  ip: string;
  duration: number;
  message: string;
}
interface PingDetectorResult {
  [key: string]: unknown;
  id: number;
  createdAt: string;
  updatedAt: string;
  task: number;
  result: DetectorResult;
  maxDuration: number;
  messages: string[];
  // 检测IP
  ips: string[];
  results: PingDetectorSubResult[];
}

interface DNSDetectorSubResult {
  [key: string]: unknown;
  result: DetectorResult;
  ips: string[];
  server: string;
  duration: number;
  message: string;
}

interface DNSDetectorResult {
  [key: string]: unknown;
  id: number;
  createdAt: string;
  updatedAt: string;
  task: number;
  result: DetectorResult;
  maxDuration: number;
  messages: string[];
  // 检测IP
  host: string;
  results: DNSDetectorSubResult[];
}

interface DatabaseDetectorSubResult {
  [key: string]: unknown;
  result: DetectorResult;
  uri: string;
  duration: number;
  message: string;
}
interface DatabaseDetectorResult {
  [key: string]: unknown;
  id: number;
  createdAt: string;
  updatedAt: string;
  task: number;
  result: DetectorResult;
  maxDuration: number;
  messages: string[];
  // 连接串
  uris: string[];
  results: DatabaseDetectorSubResult[];
}

interface List<T> {
  processing: boolean;
  items: T[];
  count: number;
}

const httpDetectors: List<HTTPDetector> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const dnsDetectors: List<DNSDetector> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const tcpDetectors: List<TCPDetector> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const pingDetectors: List<PingDetector> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const databaseDetectors: List<DatabaseDetector> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const httpDetectorResults: List<HTTPDetectorResult> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const tcpDetectorResults: List<TCPDetectorResult> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const pingDetectorResults: List<PingDetectorResult> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const dnsDetectorResults: List<DNSDetectorResult> = reactive({
  processing: false,
  items: [],
  count: -1,
});

const databaseDetectorResults: List<DatabaseDetectorResult> = reactive({
  processing: false,
  items: [],
  count: -1,
});

function fillCount(
  params: Record<string, unknown>,
  data: Record<string, unknown>,
  key: string
) {
  const offset = Number(params.offset);
  const limit = Number(params.limit);
  data.count = offset + limit;
  if (!data[key] || !(data[key] as []).length) {
    return;
  }
  // 如果刚好满一页，设置多一条
  if ((data[key] as []).length % limit === 0) {
    data.count = (data.count as number) + 1;
  }
}

export async function detectorListUser(keyword: string): Promise<string[]> {
  const { data } = await request.get<Record<string,unknown>>(DETECTOR_LIST_USER, {
    params: {
      keyword,
      limit: 20,
    },
  });
  return (data.accounts as string[]) || [];
}

const defaultDetectorQueryParams = {
  order: "-updatedAt",
};
const defaultDetectorResultQueryParams = {
  order: "-updatedAt",
  ignoreCount: "true",
};

// 查询http检测配置
export async function httpDetectorList(params: {
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (httpDetectors.processing) {
    return;
  }
  httpDetectors.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      httpDetectors: HTTPDetector[], 
    }>(HTTP_DETECTORS, {
      params: Object.assign(defaultDetectorQueryParams, params),
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

// 查询dns检测配置
export async function dnsDetectorList(params: {
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (dnsDetectors.processing) {
    return;
  }
  dnsDetectors.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      dnsDetectors: DNSDetector[],
    }>(DNS_DETECTORS, {
      params: Object.assign(defaultDetectorQueryParams, params),
    });
    const count = data.count || 0;
    if (count >= 0) {
      dnsDetectors.count = count;
    }
    dnsDetectors.items = data.dnsDetectors || [];
  } finally {
    dnsDetectors.processing = false;
  }
}

// 查询tcp检测配置
export async function tcpDetectorList(params: {
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (tcpDetectors.processing) {
    return;
  }
  tcpDetectors.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      tcpDetectors: TCPDetector[],
    }>(TCP_DETECTORS, {
      params: Object.assign(defaultDetectorQueryParams, params),
    });
    const count = data.count || 0;
    if (count >= 0) {
      tcpDetectors.count = count;
    }
    tcpDetectors.items = data.tcpDetectors || [];
  } finally {
    tcpDetectors.processing = false;
  }
}

// 查询ping检测配置
export async function pingDetectorList(params: {
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (pingDetectors.processing) {
    return;
  }
  pingDetectors.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      pingDetectors: PingDetector[],
    }>(PING_DETECTORS, {
      params: Object.assign(defaultDetectorQueryParams, params),
    });
    const count = data.count || 0;
    if (count >= 0) {
      pingDetectors.count = count;
    }
    pingDetectors.items = data.pingDetectors || [];
  } finally {
    pingDetectors.processing = false;
  }
}

// 查询redis检测配置
export async function databaseDetectorList(params: {
  limit?: number;
  offset?: number;
}): Promise<void> {
  if (databaseDetectors.processing) {
    return;
  }
  databaseDetectors.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      databaseDetectors: DatabaseDetector[],
    }>(DATABASE_DETECTORS, {
      params: Object.assign(defaultDetectorQueryParams, params),
    });
    const count = data.count || 0;
    if (count >= 0) {
      databaseDetectors.count = count;
    }
    databaseDetectors.items = data.databaseDetectors || [];
  } finally {
    databaseDetectors.processing = false;
  }
}

// 通过id查询http检测配置
export async function httpDetectorFindByID(id: number): Promise<HTTPDetector> {
  const { data } = await request.get(
    HTTP_DETECTORS_ID.replace(":id", id.toString())
  );
  return data as HTTPDetector;
}

// 通过id查询dns检测配置
export async function dnsDetectorFindByID(id: number): Promise<DNSDetector> {
  const { data } = await request.get(
    DNS_DETECTORS_ID.replace(":id", id.toString())
  );
  return data as DNSDetector;
}

// 通过id查询tcp检测配置
export async function tcpDetectorFindByID(id: number): Promise<TCPDetector> {
  const { data } = await request.get(
    TCP_DETECTORS_ID.replace(":id", id.toString())
  );
  return data as TCPDetector;
}

// 通过id查询ping检测配置
export async function pingDetectorFindByID(id: number): Promise<PingDetector> {
  const { data } = await request.get(
    PING_DETECTORS_ID.replace(":id", id.toString())
  );
  return data as PingDetector;
}

// 通过id查询database检测配置
export async function databaseDetectorFindByID(
  id: number
): Promise<DatabaseDetector> {
  const { data } = await request.get<DatabaseDetector>(
    DATABASE_DETECTORS_ID.replace(":id", id.toString())
  );
  return data;
}

// 通过id更新http检测配置
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

// 通过id更新dns检测配置
export async function dnsDetectorUpdateByID(
  id: number,
  updateData: Record<string, unknown>
): Promise<DNSDetector> {
  const { data } = await request.patch(
    DNS_DETECTORS_ID.replace(":id", id.toString()),
    updateData
  );
  return data as DNSDetector;
}

// 通过id更新tcp检测配置
export async function tcpDetectorUpdateByID(
  id: number,
  updateData: Record<string, unknown>
): Promise<TCPDetector> {
  const { data } = await request.patch(
    TCP_DETECTORS_ID.replace(":id", id.toString()),
    updateData
  );
  return data as TCPDetector;
}

// 通过id更新ping检测配置
export async function pingDetectorUpdateByID(
  id: number,
  updateData: Record<string, unknown>
): Promise<PingDetector> {
  const { data } = await request.patch(
    PING_DETECTORS_ID.replace(":id", id.toString()),
    updateData
  );
  return data as PingDetector;
}

// 通过id更新database检测配置
export async function databaseDetectorUpdateByID(
  id: number,
  updateData: Record<string, unknown>
): Promise<DatabaseDetector> {
  const { data } = await request.patch(
    DATABASE_DETECTORS_ID.replace(":id", id.toString()),
    updateData
  );
  return data as DatabaseDetector;
}

// 创建http检测配置
export async function httpDetectorCreate(
  createData: Record<string, unknown>
): Promise<HTTPDetector> {
  const { data } = await request.post(HTTP_DETECTORS, createData);
  return data as HTTPDetector;
}

// 创建dns检测配置
export async function dnsDetectorCreate(
  createData: Record<string, unknown>
): Promise<DNSDetector> {
  const { data } = await request.post(DNS_DETECTORS, createData);
  return data as DNSDetector;
}

// 创建tcp检测配置
export async function tcpDetectorCreate(
  createData: Record<string, unknown>
): Promise<TCPDetector> {
  const { data } = await request.post(TCP_DETECTORS, createData);
  return data as TCPDetector;
}

// 创建ping检测配置
export async function pingDetectorCreate(
  createData: Record<string, unknown>
): Promise<PingDetector> {
  const { data } = await request.post(PING_DETECTORS, createData);
  return data as PingDetector;
}

// 创建database检测配置
export async function databaseDetectorCreate(
  createData: Record<string, unknown>
): Promise<DatabaseDetector> {
  const { data } = await request.post(DATABASE_DETECTORS, createData);
  return data as DatabaseDetector;
}

export async function httpDetectorResultList(
  params: Record<string, unknown>
): Promise<void> {
  if (httpDetectorResults.processing) {
    return;
  }
  httpDetectorResults.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      httpDetectorResults: HTTPDetectorResult[],
    }>(HTTP_DETECTOR_RESULTS, {
      params: Object.assign(defaultDetectorResultQueryParams, params),
    });
    fillCount(params, data, "httpDetectorResults");
    httpDetectorResults.items = data.httpDetectorResults || [];
    httpDetectorResults.items.forEach((item) => {
      if (!item.results) {
        return;
      }
      item.results.forEach((subItem) => {
        const dates = subItem.certificateExpirationDates;
        if (dates && dates.length === 2) {
          const size = 10;
          subItem.certificateExpirationDates = [
            dates[0].substring(0, size),
            dates[1].substring(0, size),
          ];
        }
        const values: Record<string, number> = {
          TOTAL: subItem.duration,
          DNS: subItem.dnsLookup,
          TCP: subItem.tcpConnection,
          TLS: subItem.tlsHandshake,
          PROCESSING: subItem.serverProcessing,
          TRANSFER: subItem.contentTransfer,
        };
        const timeline: string[] = [];
        Object.keys(values).forEach((key) => {
          const v = values[key];
          if (!v) {
            return;
          }
          timeline.push(`${key}: ${v}`);
        });
        subItem.timeline = timeline;
      });
    });
    httpDetectorResults.count = data.count;
  } finally {
    httpDetectorResults.processing = false;
  }
}

export async function tcpDetectorResultList(
  params: Record<string, unknown>
): Promise<void> {
  if (tcpDetectorResults.processing) {
    return;
  }
  tcpDetectorResults.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      tcpDetectorResults: TCPDetectorResult[],
    }>(TCP_DETECTOR_RESULTS, {
      params: Object.assign(defaultDetectorResultQueryParams, params),
    });
    fillCount(params, data, "tcpDetectorResults");
    tcpDetectorResults.items = data.tcpDetectorResults || [];
    tcpDetectorResults.count = data.count;
  } finally {
    tcpDetectorResults.processing = false;
  }
}

export async function pingDetectorResultList(
  params: Record<string, unknown>
): Promise<void> {
  if (pingDetectorResults.processing) {
    return;
  }
  pingDetectorResults.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      pingDetectorResults: PingDetectorResult[],
    }>(PING_DETECTOR_RESULTS, {
      params: Object.assign(defaultDetectorResultQueryParams, params),
    });
    fillCount(params, data, "pingDetectorResults");
    pingDetectorResults.items = data.pingDetectorResults || [];
    pingDetectorResults.count = data.count;
  } finally {
    pingDetectorResults.processing = false;
  }
}

export async function dnsDetectorResultList(
  params: Record<string, unknown>
): Promise<void> {
  if (dnsDetectorResults.processing) {
    return;
  }
  dnsDetectorResults.processing = true;
  try {
    const { data } = await request.get<{
      count: number,
      dnsDetectorResults: DNSDetectorResult[],
    }>(DNS_DETECTOR_RESULTS, {
      params: Object.assign(defaultDetectorResultQueryParams, params),
    });
    fillCount(params, data, "dnsDetectorResults");
    dnsDetectorResults.items = data.dnsDetectorResults || [];
    dnsDetectorResults.count = data.count;
  } finally {
    dnsDetectorResults.processing = false;
  }
}

export async function databaseDetectorResultList(
  params: Record<string, unknown>
): Promise<void> {
  if (databaseDetectorResults.processing) {
    return;
  }
  databaseDetectorResults.processing = true;
  try {
    const { data } = await request.get<
    {
      count: number,
      databaseDetectorResults: DatabaseDetectorResult[],
    }>(DATABASE_DETECTOR_RESULTS, {
      params: Object.assign(defaultDetectorResultQueryParams, params),
    });
    fillCount(params, data, "databaseDetectorResults");
    databaseDetectorResults.items = data.databaseDetectorResults || [];
    databaseDetectorResults.count = data.count;
  } finally {
    databaseDetectorResults.processing = false;
  }
}

interface ReadonlyDetectorState {
  httpDetectors: DeepReadonly<List<HTTPDetector>>;
  dnsDetectors: DeepReadonly<List<DNSDetector>>;
  tcpDetectors: DeepReadonly<List<TCPDetector>>;
  pingDetectors: DeepReadonly<List<PingDetector>>;
  databaseDetectors: DeepReadonly<List<DatabaseDetector>>;
  httpDetectorResults: DeepReadonly<List<HTTPDetectorResult>>;
  tcpDetectorResults: DeepReadonly<List<TCPDetectorResult>>;
  pingDetectorResults: DeepReadonly<List<PingDetectorResult>>;
  dnsDetectorResults: DeepReadonly<List<DNSDetectorResult>>;
  databaseDetectorResults: DeepReadonly<List<DatabaseDetectorResult>>;
}

const state = {
  httpDetectors: readonly(httpDetectors),
  dnsDetectors: readonly(dnsDetectors),
  tcpDetectors: readonly(tcpDetectors),
  pingDetectors: readonly(pingDetectors),
  databaseDetectors: readonly(databaseDetectors),
  httpDetectorResults: readonly(httpDetectorResults),
  tcpDetectorResults: readonly(tcpDetectorResults),
  pingDetectorResults: readonly(pingDetectorResults),
  dnsDetectorResults: readonly(dnsDetectorResults),
  databaseDetectorResults: readonly(databaseDetectorResults),
};

export default function useDetectorState(): ReadonlyDetectorState {
  return state;
}
