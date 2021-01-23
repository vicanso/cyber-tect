<template lang="pug">
.router
  div(
    v-if="!editMode"
  )
    config-table(
      :category="category"
      name="路由配置"
    )
    .add: el-button.addBtn(
      type="primary"
      @click="add"
    ) 添加
  config-editor(
    v-else
    name="添加/更新路由配置"
    summary="配置针对各路由响应的Mock"
    :category="category"
    :defaultValue="defaultValue"
  ): template(
    #data="configProps"
  ): router-data(
    :data="configProps.form.data"
    @change.self="configProps.form.data = $event"
  )
</template>

<script lang="ts">
import { defineComponent } from "vue";

import ConfigEditor from "../../components/configs/Editor.vue";
import RouterData from "../../components/configs/RouterData.vue";
import ConfigTable from "../../components/configs/Table.vue";
import { ROUTER, CONFIG_EDIT_MODE } from "../../constants/common";

export default defineComponent({
  name: "Router",
  components: {
    RouterData,
    ConfigTable,
    ConfigEditor,
  },
  data() {
    return {
      defaultValue: {
        category: ROUTER,
      },
      category: ROUTER,
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
