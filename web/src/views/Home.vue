<template lang="pug">
  .home
    el-card
      div(
        slot="header"
      )
        | HTTP
        router-link.more(
          :to="moreHTTPResult"
        ) 查看更多
      ListHTTPResult(
        :limit="5"
        :simplify="true"
      )
</template>
<script>
import ListHTTPResult from '@/components/ListHTTPResult.vue'
import {
  LIST_HTTP_RESULT
} from '@/paths'

export default {
  name: 'Home',
  components: {
    ListHTTPResult
  },
  data () {
    return {
      moreHTTPResult: LIST_HTTP_RESULT
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.home
  margin: $mainMargin
.more
  font-size: 12px
  float: right
</style>
