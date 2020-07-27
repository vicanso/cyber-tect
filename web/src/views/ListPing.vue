<template lang="pug">
.listPing
  ListDetector(
    :category="category"
  )
  .add
    el-button(
      type="primary"
      @click="add"
    ) 添加Ping检测
</template>
<script>
import ListDetector from "@/components/ListDetector.vue";
import { CAT_PING } from "@/constants/category";
import { ROUTE_ADD_PING } from "@/router";

export default {
  name: "ListPing",
  components: {
    ListDetector,
  },
  data() {
    return {
      category: CAT_PING,
    };
  },
  methods: {
    add() {
      this.$router.push({
        name: ROUTE_ADD_PING,
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
