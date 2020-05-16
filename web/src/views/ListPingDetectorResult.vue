<template lang="pug">
.listPingDetectorResult
  el-card
    div(
      slot="header"
    ) Ping
    ListPingResult
</template>
<script>
import ListPingResult from '@/components/ListPingResult.vue'
export default {
  name: 'ListPingDetectorResult',
  components: {
    ListPingResult
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.listPingDetectorResult
  margin: $mainMargin
</style>
