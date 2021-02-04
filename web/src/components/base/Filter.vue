<template lang="pug">
el-form.baseFilter(
  :label-width="$props.labelWidth"
): el-row(
  :gutter="15"
)
  el-col(
    v-for="field in $props.fields"
    :span="field.span || 8"
    :key="field.key"
  )
    el-form-item(
      :label="field.label"
      :label-width="field.labelWidth"
      :class="field.itemClass"
    )
      //- 列表选择
      el-select.select(
        v-if="field.type === 'select'"
        :placeholder="field.placeholder"
        v-model="current[field.key]"
        :multiple="field.multiple || false"
      )
        el-option(
          v-for="item in field.options"
          :key="item.key || item.value"
          :label="item.label || item.name"
          :value="item.value"
        )
      //- 点击筛选
      ex-button(
        v-else-if="field.type === 'filter'"
        :onClick="filter"
        icon="el-icon-search"
      ) 筛选
      //- 日期时间筛选
      el-date-picker.dateRange.fullFill(
        v-else-if="field.type === 'dateTimeRange'"
        v-model="current[field.key]"
        type="datetimerange"
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        :shortcuts="field.shortcuts"
      )
      //- 日期筛选
      el-date-picker.dateRange.fullFill(
        v-else-if="field.type === 'dateRange'"
        v-model="current[field.key]"
        type="daterange"
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        :shortcuts="field.shortcuts"
      )
      el-input(
        v-else-if="field.type === 'number'"
        v-model="current[field.key]"
        type="number"
        :placeholder="field.placeholder"
        :default="field.defaultValue"
      )
        template(
          v-if="field.suffix"
          #suffix
        ) {{field.suffix}}
      //- 关键字搜索
      el-input(
        v-else
        @keyup.enter.native="filter"
        :clearable="field.clearable"
        v-model="current[field.key]"
        :disabled="field.disabled || false"
        :placeholder="field.placeholder"
      )
</template>

<script lang="ts">
import { defineComponent } from "vue";

import ExButton from "../ExButton.vue";

export default defineComponent({
  name: "BaseFilter",
  components: {
    ExButton,
  },
  props: {
    labelWidth: {
      type: String,
      default: "90px",
    },
    fields: {
      type: Array,
      required: true,
    },
  },
  emits: ["filter"],
  data() {
    const current = {};
    const { fields } = this.$props;
    fields.forEach((item) => {
      const { type, key, defaultValue } = item;
      if (type === "filter") {
        return;
      }
      current[key] = defaultValue || "";
    });
    return {
      processing: false,
      current,
    };
  },
  methods: {
    filter() {
      this.$emit("filter", this.current);
    },
  },
});
</script>
<style lang="stylus" scoped>
.baseFilter
  .select, .btn, .dateRange
    width: 100%
</style>
