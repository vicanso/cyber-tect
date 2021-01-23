<template lang="pug">
.signedKey
  div(
    v-if="!editMode"
  )
    config-table(
      :category="category"
      name="SignedKey配置"
    )
    .add: el-button.addBtn(
      type="primary"
      @click="add"
    ) 添加
  config-editor(
    v-else
    name="添加/更新SignedKey配置"
    summary="用于配置session中使用的signed key"
    :category="category"
    :defaultValue="defaultValue"
  ): template(
    #data="configProps"
  ): signed-key-data(
    :data="configProps.form.data"
    @change.self="configProps.form.data = $event"
  )
</template>

<script lang="ts">
import { defineComponent } from "vue";

import ConfigEditor from "../../components/configs/Editor.vue";
import SignedKeyData from "../../components/configs/SignedKeyData.vue";
import ConfigTable from "../../components/configs/Table.vue";
import { SIGNED_KEY, CONFIG_EDIT_MODE } from "../../constants/common";

export default defineComponent({
  name: "BlockIP",
  components: {
    SignedKeyData,
    ConfigTable,
    ConfigEditor,
  },
  data() {
    return {
      defaultValue: {
        category: SIGNED_KEY,
      },
      category: SIGNED_KEY,
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
