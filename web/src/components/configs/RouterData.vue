<template lang="pug">
//- 占位
el-col(
  :span="8"
)
//- 路由配置选择
el-col(
  :span="8"
): el-form-item(
  label="路由选择："
): router-selector.selector(
  :router="form.router"
  @change.self="handleChangeRouter"
)
//- 响应状态码
el-col(
  :span="8"
): el-form-item(
  label="状态码："
): el-input(
  v-model="form.status"
  type="number"
  placeholder="请输入响应状态码"
)
//- 响应类型
el-col(
  :span="8"
): el-form-item(
  label="响应类型："
): el-select.selector(
  v-model="form.contentType"
  placeholder="请选择响应类型"
): el-option(
  v-for="item in contentTypeList"
  :key="item"
  :label="item"
  :value="item"
)
//- 延时响应
el-col(
  :span="8"
): el-form-item(
  label="延时响应："
): el-input(
  type="number"
  v-model="form.delaySeconds",
  placeholder="请输入延时响应时长，可选"
): template(
  #append
) 秒
//- 完整URL
el-col(
  :span="16"
): el-form-item(
  label="完整URL："
): el-input(
  type="text"
  v-model="form.url"
  placeholder="请输入完整的请求URL(包含参数部分），可选"
  clearable
)
//- 响应数据
el-col(
  :span="24"
): el-form-item(
  label="响应数据："
): el-input(
  v-model="form.response"
  type="textarea"
  :autosize="{ minRows: 5, maxRows: 10 }"
  placeholder="请按选择的响应类型输入对应的响应数据"
)
</template>
<script lang="ts">
import { defineComponent } from "vue";

import RouterSelector from "../../components/configs/RouterSelector.vue";

export default defineComponent({
  name: "RouterData",
  components: {
    RouterSelector,
  },
  props: {
    data: {
      type: String,
      default: "",
    },
  },
  emits: ["change"],
  data() {
    const form = {
      router: "",
      status: null,
      contentType: "",
      response: "",
      delaySeconds: null,
      path: "",
    };
    if (this.$props.data) {
      const data = JSON.parse(this.$props.data);
      data.router = `${data.method} ${data.route}`;
      if (data.response && data.response[0] === "{") {
        data.response = JSON.stringify(JSON.parse(data.response), null, 2);
      }
      Object.assign(form, data);
    }
    return {
      contentTypeList: [
        "application/json; charset=UTF-8",
        "text/plain; charset=UTF-8",
        "text/html; charset=UTF-8",
      ],
      form,
    };
  },
  watch: {
    "form.status": function () {
      this.handleChange();
    },
    "form.contentType": function () {
      this.handleChange();
    },
    "form.response": function () {
      this.handleChange();
    },
    "form.delaySeconds": function () {
      this.handleChange();
    },
    "form.url": function () {
      this.handleChange();
    },
  },
  methods: {
    handleChangeRouter(value) {
      this.form.router = value;
      this.handleChange();
    },
    handleChange() {
      const {
        router,
        status,
        contentType,
        response,
        delaySeconds,
        url,
      } = this.form;
      let value = "";
      if (router && status && contentType && response) {
        const [method, route] = router.split(" ");
        const data = {
          route,
          method,
          status: Number(status),
          contentType,
          response: response.trim(),
          delaySeconds: 0,
          url: "",
        };
        if (delaySeconds) {
          data.delaySeconds = Number(delaySeconds);
          if (data.delaySeconds < 0) {
            this.$error("延时时长不能小于0");
            return;
          }
        }
        if (url) {
          data.url = url;
        }
        value = JSON.stringify(data);
      }
      this.$emit("change", value);
    },
  },
});
</script>
<style lang="stylus" scoped>
.selector
  width: 100%
</style>
