<template lang="pug">
.routerConcurrency
  div(
    v-if="!editMode"
  )
    config-table(
      :category="category"
      name="路由并发配置"
    )
    .add: el-button.addBtn(
      type="primary"
      @click="add"
    ) 添加
  config-editor(
    v-else
    name="添加/更新路由并发配置"
    summary="配置针对各路由并发请求的限制"
    :category="category"
    :defaultValue="defaultValue"
  ): template(
    #data="configProps"
  ): router-concurrency-data(
    :data="configProps.form.data"
    @change.self="configProps.form.data = $event"
  )
</template>

<script lang="ts">
import { defineComponent } from "vue";

import ConfigEditor from "../../components/configs/Editor.vue";
import RouterConcurrencyData from "../../components/configs/RouterConcurrencyData.vue";
import ConfigTable from "../../components/configs/Table.vue";
import { ROUTER_CONCURRENCY, CONFIG_EDIT_MODE } from "../../constants/common";

export default defineComponent({
  name: "RouterConcurrency",
  components: {
    RouterConcurrencyData,
    ConfigTable,
    ConfigEditor,
  },
  data() {
    return {
      defaultValue: {
        category: ROUTER_CONCURRENCY,
      },
      category: ROUTER_CONCURRENCY,
    };
  },
  computed: {
    editMode() {
      return this.$route.query.mode === CONFIG_EDIT_MODE;
    },
  },
  methods: {
    add() {
      this.$router.push({
        query: {
          mode: CONFIG_EDIT_MODE,
        },
      });
    },
  },
});
</script>
<style lang="stylus" scoped>
@import "../../common";

.add
  margin: $mainMargin
.addBtn
  width: 100%
</style>
