<template lang="pug">
.listDNS
  ListDetector(
    :category="category"
  )
  .add
    el-button(
      type="primary"
      @click="add"
    ) 添加DNS检测
</template>
<script>
import ListDetector from "@/components/ListDetector.vue";
import { CAT_DNS } from "@/constants/category";
import { ROUTE_ADD_DNS } from "@/router";

export default {
  name: "ListDNS",
  components: {
    ListDetector,
  },
  data() {
    return {
      category: CAT_DNS,
    };
  },
  methods: {
    add() {
      this.$router.push({
        name: ROUTE_ADD_DNS,
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
