<template lang="pug">
el-select(
  v-model="currentRouter"
  placeholder="请选择路由"
  v-loading="routers.processing"
  @change="handleChange"
): el-option(
  v-for="item in routers.items"
  :key="`${item.method} ${item.route}`"
  :label="`${item.method} ${item.route}`"
  :value="`${item.method} ${item.route}`"
)
</template>
<script lang="ts">
import { defineComponent } from "vue";

import { useCommonStore } from "../../store";

export default defineComponent({
  name: "RouterSelector",
  props: {
    router: {
      type: String,
      default: "",
    },
  },
  emits: ["change"],
  setup() {
    const commonStore = useCommonStore();
    return {
      routers: commonStore.state.routers,
      listRouter: () => commonStore.dispatch("listRouter"),
    };
  },
  data() {
    return {
      currentRouter: this.$props.router || "",
    };
  },
  beforeMount() {
    this.fetch();
  },
  methods: {
    handleChange(value) {
      this.$emit("change", value);
    },
    async fetch() {
      try {
        await this.listRouter();
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>
