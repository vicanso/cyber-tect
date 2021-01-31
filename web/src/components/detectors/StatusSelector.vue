<template lang="pug">
el-form-item(
  label="状态："
): el-select.select(
  placeholder="请选择配置状态"
  v-model="value"
  @change="change"
): el-option(
  v-for="item in statusList"
  :key="item.name"
  :label="item.name"
  :value="item.value"
)
</template>
<script lang="ts">
import { defineComponent } from "vue";

import { getDetectorStatusList } from "../../helpers/util";

export default defineComponent({
  name: "DetectorStatusSelector",
  props: {
    status: {
      type: Number,
      default: null,
    },
  },
  emits: ["change"],
  data() {
    return {
      value: this.$props.status,
      statusList: getDetectorStatusList(),
    };
  },
  methods: {
    change(value) {
      this.$emit("change", value);
    },
  },
});
</script>
<style lang="stylus" scoped>
.select
  width: 100%
</style>
