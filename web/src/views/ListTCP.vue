<template lang="pug">
.listTCP
  ListDetector(
    :category="category"
  )
  .add
    el-button(
      type="primary"
      @click="add"
    ) 添加TCP检测
</template>
<script>
import ListDetector from "@/components/ListDetector.vue";
import { CAT_TCP } from "@/constants/category";
import { ROUTE_ADD_TCP } from "@/router";

export default {
  name: "ListTCP",
  components: {
    ListDetector,
  },
  data() {
    return {
      category: CAT_TCP,
    };
  },
  methods: {
    add() {
      this.$router.push({
        name: ROUTE_ADD_TCP,
      });
    },
  },
};
</script>
<style lang="sass" scoped>
@import '@/common.sass'
.add
  margin: 0 $mainMargin $mainMargin $mainMargin
  .el-button
    width: 100%
</style>
