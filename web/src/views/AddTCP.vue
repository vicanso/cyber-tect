<template lang="pug">
el-card.addTCP
  div(
    slot="header"
  )
    span 添加TCP检测
  TCPDetector
</template>

<script>
import TCPDetector from "@/components/TCPDetector.vue";

export default {
  name: "AddTCP",
  components: {
    TCPDetector,
  },
};
</script>
<style lang="sass" scoped>
@import '@/common'
.addTCP
  margin: $mainMargin
</style>
