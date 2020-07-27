import Vue from "vue";
import Element from "element-ui";
import "element-ui/lib/theme-chalk/index.css";

import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import "@/main.sass";
import { isDevelopment } from "@/constants/env";

Vue.config.productionTip = false;
Vue.config.errorHandler = function (err, vm, info) {
  // TODO 错误异常上报
  const { tag } = vm.$vnode.componentOptions;
  console.error(err);
  console.info(vm);
  console.dir(tag);
  console.info(info);
  // 如果开发环境，弹alert
  if (isDevelopment()) {
    alert(`${tag} ${info}`);
  }
};
window.onerror = function (msg, url, lineNo, columnNo, err) {
  // TODO 错误异常上报
  console.error(err);
};
Vue.config.keyCodes = {
  // 回车
  enter: 0x0d,
};
Vue.use(Element);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
