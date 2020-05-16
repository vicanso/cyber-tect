<template lang="pug">
.listDNSDetectorResult
  el-card
    div(
      slot="header"
    ) DNS
    ListDNSResult
</template>
<script>
import ListDNSResult from '@/components/ListDNSResult.vue'
export default {
  name: 'ListDNSDetectorResult',
  components: {
    ListDNSResult
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.listDNSDetectorResult
  margin: $mainMargin
</style>
