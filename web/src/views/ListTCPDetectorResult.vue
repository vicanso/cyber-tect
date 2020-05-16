<template lang="pug">
.listTCPDetectorResult
  el-card
    div(
      slot="header"
    ) TCP
    ListTCPResult
</template>
<script>
import ListTCPResult from '@/components/ListTCPResult.vue'
export default {
  name: 'ListTCPDetectorResult',
  components: {
    ListTCPResult
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.listTCPDetectorResult
  margin: $mainMargin
</style>
