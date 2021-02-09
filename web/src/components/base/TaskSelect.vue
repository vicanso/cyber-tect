<template lang="pug">
el-select.select(
  @change="handleChange"
  filterable
  remote
  reserve-keyword
  v-model="task"
  placeholder="请选择任务"
  :remote-method="fetch"
  :loading="processing"
): el-option(
  v-for="item in tasks"
  :key="item.id"
  :value="item.id"
  :label="item.name"
)
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useDetectorStore } from "../../store";

export default defineComponent({
  name: "BaseTaskSelect",
  props: {
    category: {
      type: String,
      required: true,
    },
  },
  emits: ["input"],
  setup() {
    const detectorStore = useDetectorStore();
    return {
      filter: (params) => detectorStore.dispatch("filter", params),
    };
  },
  data() {
    return {
      task: null,
      processing: false,
      tasks: [],
    };
  },
  methods: {
    handleChange(value) {
      this.$emit("input", value);
    },
    async fetch(keyword) {
      const { processing } = this;
      if (processing) {
        return;
      }
      this.tasks = [];
      if (!keyword) {
        return;
      }
      try {
        this.processing = true;
        const data = await this.filter({
          keyword,
          category: this.$props.category,
        });
        Object.keys(data).forEach((key) => {
          this.tasks = data[key] || [];
        });
      } catch (err) {
        this.$error(err);
      } finally {
        this.processing = false;
      }
    },
  },
});
</script>
<style lang="stylus" scoped>
.select
  width 100%
</style>
