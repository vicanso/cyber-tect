<template lang="pug">
el-card.addHTTP
  div(
    slot="header"
  )
    span 添加HTTP检测
  HTTPDetector
</template>

<script>
import HTTPDetector from "@/components/HTTPDetector.vue";

export default {
  name: "AddHTTP",
  components: {
    HTTPDetector,
  },
};
</script>
<style lang="sass" scoped>
@import '@/common'
.addHTTP
  margin: $mainMargin
</style>
