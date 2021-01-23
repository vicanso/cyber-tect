<template lang="pug">
el-col(
  :span="8"
): el-form-item.signedKeyData(
  label="SignedKey："
): el-input(
  placeholder="请输入签名使用的密钥，多个密钥以,分隔"
  v-model="key"
  clearable
)
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "SignedKeyData",
  props: {
    data: {
      type: String,
      default: "",
    },
  },
  emits: ["change"],
  data() {
    return {
      key: this.$props.data || "",
    };
  },
  watch: {
    key: function (value) {
      this.$emit("change", value);
    },
  },
});
</script>
