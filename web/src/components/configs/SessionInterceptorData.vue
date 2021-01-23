<template lang="pug">
el-row.sessionInterceptorData
  el-col(
    :span="8"
  ): el-form-item.sessionInterceptorData(
    label="提示信息："
  ): el-input(
    placeholder="请输入提示信息"
    v-model="form.message"
    clearable
  )
  el-col(
    :span="16"
  ): el-form-item(
    label="允许账号："
  ): el-input(
    placeholder="请输入允许的账号列表，多个账号以,分割"
    v-model="form.allowAccounts"
    clearable
  )
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "SessionInterceptorData",
  props: {
    data: {
      type: String,
      default: "",
    },
  },
  emits: ["change"],
  data() {
    const form = {
      message: "",
      allowAccounts: "",
    };
    if (this.$props.data) {
      const data = JSON.parse(this.$props.data);
      form.message = data.message;
      form.allowAccounts = (data.allowAccounts || []).join(",");
    }
    return {
      form,
    };
  },
  watch: {
    "form.message": function () {
      this.handleChange();
    },
    "form.allowAccounts": function () {
      this.handleChange();
    },
  },
  methods: {
    handleChange() {
      this.$emit(
        "change",
        JSON.stringify({
          message: this.form.message,
          allowAccounts: this.form.allowAccounts.split(","),
        })
      );
    },
  },
});
</script>
<style lang="stylus" scoped>
.sessionInterceptorData
  width: 100%
</style>
