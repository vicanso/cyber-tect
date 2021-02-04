<template lang="pug">
//- 筛选条件
base-filter(
  :fields="filterFields"
  @filter="filter"
)
</template>

<script lang="ts">
import { defineComponent } from "vue";

import BaseFilter from "../../components/base/Filter.vue";

const filterFields = [
  {
    label: "状态：",
    key: "result",
    type: "select",
    placeholder: "请选择检测状态",
    span: 6,
    options: [
      {
        name: "所有",
        value: 0,
      },
      {
        name: "成功",
        value: 1,
      },
      {
        name: "失败",
        value: 2,
      },
    ],
  },
  {
    label: "耗时大于：",
    key: "duration",
    type: "number",
    placeholder: "请输入耗时",
    suffix: "秒",
    span: 6,
  },
  {
    label: "任务：",
    key: "task",
    type: "autocomplete",
    placeholder: "请输入关键字",
    clearable: true,
    span: 6,
  },
  {
    label: "",
    type: "filter",
    labelWidth: "0px",
    span: 6,
  },
];

export default defineComponent({
  name: "DetectorResultsFilter",
  components: {
    BaseFilter,
  },
  emits: ["filter"],
  data() {
    return {
      filterFields,
    };
  },
  methods: {
    filter(params) {
      this.$emit("filter", params);
    },
  },
});
</script>
