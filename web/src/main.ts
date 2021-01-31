import { createApp } from "vue";
import ElementPlus from "element-plus";
import { ElMessage } from "element-plus";
import "element-plus/lib/theme-chalk/index.css";
import App from "./App.vue";
import ExButton from "./components/ExButton.vue";
import Router, { getCurrentLocation } from "./router";
import stores from "./store";
import "./main.styl";
import { addUserAction, ERROR, FAIL } from "./services/action";
import { isDevelopment } from "./constants/env";

const app = createApp(App);
// 全局出错处理
app.config.errorHandler = (err: any, vm, info) => {
  // 处理错误
  let message = "";
  if (err && err.message) {
    message = err.message;
  }
  if (info) {
    message += ` [${info}]`;
  }
  const currentLocation = getCurrentLocation();
  addUserAction({
    category: ERROR,
    route: currentLocation.name,
    path: currentLocation.path,
    result: FAIL,
    message,
  });
  throw err;
};
// 自定义全局出错提示
app.config.globalProperties.$error = function (err: any) {
  if (!(err instanceof Error)) {
    ElMessage.error(err);
    return;
  }
  let message = err.message;
  if (err.category) {
    message += ` [${err.category.toUpperCase()}]`;
  }
  if (err.code) {
    message += ` [${err.code}]`;
  }
  ElMessage.error(message);
  // 如果是异常（客户端异常，如请求超时，中断等），则上报user action
  if (err.exception) {
    const currentLocation = getCurrentLocation();
    addUserAction({
      category: ERROR,
      route: currentLocation.name,
      path: currentLocation.path,
      result: FAIL,
      message,
    });
  }
  if (isDevelopment()) {
    console.error(err);
  }
  return;
};

// 全局注册组件
app.component("ExButton", ExButton);

stores.forEach(app.use);

app.use(Router).use(ElementPlus).mount("#app");
