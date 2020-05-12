import Vue from 'vue'
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import App from '@/App.vue'
import router from '@/router'
import store from '@/store'
import '@/main.sass'

Vue.use(Element)
Vue.config.productionTip = false
Vue.config.errorHandler = function (err, vm, info) {
  // TODO 错误异常上报
  console.error(err)
}
window.onerror = function (msg, url, lineNo, columnNo, err) {
  // TODO 错误异常上报
  this.console.error(err)
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
