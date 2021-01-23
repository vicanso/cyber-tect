<template lang="pug">
el-form-item(
  label="接收者："
): el-select.select(
  placeholder="请选择接收者"
  v-model="values"
  multiple
  :loading="detectorReceivers.processing"
  @change="change"
): el-option(
  v-for="item in detectorReceivers.items"
  :key="item"
  :label="item"
  :value="item"
) 
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useDetectorStore } from "../../store";

export default defineComponent({
  name: "DetectorReceiverSelector",
  props: {
    receivers: {
      type: Array,
      default: () => [],
    },
  },
  emits: ["change"],
  setup() {
    const detectorStore = useDetectorStore();
    return {
      detectorReceivers: detectorStore.state.receivers,
      listReceiver: () => detectorStore.dispatch("listReceiver"),
    };
  },
  data() {
    return {
      values: this.$props.receivers,
    };
  },
  async mounted() {
    try {
      await this.listReceiver();
    } catch (err) {
      this.$error(err);
    }
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
