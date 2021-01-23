<template lang="pug">
.blockIP
  div(
    v-if="!editMode"
  )
    config-table(
      :category="category"
      name="黑名单IP配置"
    )
    .add
      el-button.addBtn(
        type="primary"
        @click="add"
      ) 添加
  config-editor(
    name="添加/更新IP黑名单配置"
    summary="用于拦截访问IP"
    :category="category"
    :defaultValue="defaultValue"
    v-else
  ): template(
    #data="configProps"
  ): block-iP-data(
    :data="configProps.form.data"
    @change.self="configProps.form.data = $event"
  )
</template>

<script lang="ts">
import { defineComponent } from "vue";

import ConfigEditor from "../../components/configs/Editor.vue";
import BlockIPData from "../../components/configs/BlockIPData.vue";
import ConfigTable from "../../components/configs/Table.vue";
import { BLOCK_IP, CONFIG_EDIT_MODE } from "../../constants/common";

export default defineComponent({
  name: "BlockIP",
  components: {
    BlockIPData,
    ConfigTable,
    ConfigEditor,
  },
  data() {
    return {
      defaultValue: {
        category: BLOCK_IP,
      },
      category: BLOCK_IP,
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
