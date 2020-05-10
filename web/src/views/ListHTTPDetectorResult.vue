<template lang="pug">
  .listHTTPDetectorResult
    el-card
      div(
        slot="header"
      )
        | HTTP
      ListHTTPResult
</template>
<script>
import ListHTTPResult from '@/components/ListHTTPResult.vue'

export default {
  name: 'ListHTTPDetectorResultView',
  components: {
    ListHTTPResult
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.listHTTPDetectorResult
  margin: $mainMargin
</style>
