import { createApp } from "vue";
import ElementPlus from "element-plus";
import { ElMessage } from "element-plus";
import "element-plus/lib/theme-chalk/index.css";
import App from "./App.vue";
import ExButton from "./components/ExButton.vue";
import Router from "./router";
import stores from "./store";
import "./main.styl";

const app = createApp(App);

// 全局注册组件
app.component("ExButton", ExButton);

stores.forEach(app.use);

app.use(Router).use(ElementPlus).mount("#app");
// 自定义全局出错提示
app.config.globalProperties.$error = function (err: any) {
  if (!(err instanceof Error)) {
    ElMessage.error(err);
    return;
  }
  let message = err.message;
  if (err.category) {
    message += ` [${err.category}]`;
  }
  if (err.code) {
    message += ` [${err.code}]`;
  }
  ElMessage.error(message);
  return;
};
