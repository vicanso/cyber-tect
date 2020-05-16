<template lang="pug">
el-card.addPing
  div(
    slot="header"
  )
    span 添加Ping检测
  PingDetector
</template>

<script>
import PingDetector from '@/components/PingDetector.vue'

export default {
  name: 'AddPing',
  components: {
    PingDetector
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.addPing
  margin: $mainMargin
</style>
