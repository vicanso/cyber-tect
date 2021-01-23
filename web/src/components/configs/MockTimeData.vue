<template lang="pug">
el-col(
  :span="8"
): el-form-item.mockTimeData(
  label="时间配置："
): el-date-picker.datePicker.fullFill(
  type="datetime"
  placeholder="选择要设置的时间"
  v-model="time"
) 
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "MockTimeData",
  props: {
    data: {
      type: String,
      default: "",
    },
  },
  emits: ["change"],
  data() {
    return {
      time: this.$props.data || "",
    };
  },
  watch: {
    time: function (value) {
      if (value) {
        value = value.toISOString();
      }
      this.$emit("change", value);
    },
  },
});
</script>
