import axios from "axios";
import { gzip } from "pako";

import { isDevelopment } from "../constants/env";

const requestedAt = "X-Requested-At";
// 最小压缩长度
const compressMinLength = 10 * 1024;
const supportGzip = typeof TextEncoder !== "undefined";
const request = axios.create({
  // 默认超时为10秒
  timeout: 10 * 1000,
  transformRequest: [
    (data, header) => {
      if (!data) {
        return;
      }
      header["Content-Type"] = "application/json;charset=UTF-8";
      const postData = JSON.stringify(data);
      // 如果数据较小或者不支持压缩
      if (postData.length < compressMinLength || !supportGzip) {
        return postData;
      }
      header["Content-Encoding"] = "gzip";
      return gzip(new TextEncoder().encode(postData));
    },
  ],
});

request.interceptors.request.use(
  (config) => {
    // 对请求的query部分清空值
    if (config.params) {
      Object.keys(config.params).forEach((element) => {
        // 空字符
        if (config.params[element] === "") {
          delete config.params[element];
        }
      });
    }
    if (isDevelopment()) {
      config.url = `/api${config.url}`;
    }
    config.headers[requestedAt] = `${Date.now()}`;
    return config;
  },
  (err) => {
    return Promise.reject(err);
  }
);

// 设置接口最少要xms才完成，能让客户看到loading
const minUse = 150;
const timeoutErrorCodes = ["ECONNABORTED", "ECONNREFUSED", "ECONNRESET"];
request.interceptors.response.use(
  async (res) => {
    const value = res.config.headers[requestedAt];
    if (value) {
      const use = Date.now() - Number(value);
      if (use >= 0 && use < minUse) {
        await new Promise((resolve) => setTimeout(resolve, minUse - use));
      }
    }
    return res;
  },
  (err) => {
    const { response } = err;
    if (timeoutErrorCodes.includes(err.code)) {
      err.exception = true;
      err.category = "timeout";
      err.message = "请求超时，请稍候再试";
    } else if (response) {
      if (response.data && response.data.message) {
        err.message = response.data.message;
        err.code = response.data.code;
        err.category = response.data.category;
      } else {
        err.exception = true;
        err.category = "exception";
        err.message = `未知错误[${response.statusCode || -1}]`;
      }
    }
    return Promise.reject(err);
  }
);

export default request;
