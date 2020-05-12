<template lang="pug">
  .home
    template(
      v-if="userAccount"
    )
      el-card.detectResult
        div(
          slot="header"
        )
          | 最新HTTP检测
          router-link.more(
            :to="moreHTTPResult"
          )
            | 查看更多
            i.el-icon-more
        ListHTTPResult(
          :limit="limit"
          :simplify="simplify"
        )
      el-card.detectResult
        div(
          slot="header"
        )
          | 最新DNS检测
          router-link.more(
            :to="moreDNSResult"
          )
            | 查看更多
            i.el-icon-more
        ListDNSResult(
          :limit="limit"
          :simplify="simplify"
        )
      el-card.detectResult
        div(
          slot="header"
        )
          | 最新TCP检测
          router-link.more(
            :to="moreTCPResult"
          )
            | 查看更多
            i.el-icon-more
        ListTCPResult(
          :limit="limit"
          :simplify="simplify"
        )
      el-card.detectResult
        div(
          slot="header"
        )
          | 最新Ping检测
          router-link.more(
            :to="morePingResult"
          )
            | 查看更多
            i.el-icon-more
        ListPingResult(
          :limit="limit"
          :simplify="simplify"
        )
</template>
<script>
import { mapState } from 'vuex'

import ListHTTPResult from '@/components/ListHTTPResult.vue'
import ListDNSResult from '@/components/ListDNSResult.vue'
import ListTCPResult from '@/components/ListTCPResult.vue'
import ListPingResult from '@/components/ListPingResult.vue'
import {
  LIST_HTTP_RESULT,
  LIST_DNS_RESULT,
  LIST_TCP_RESULT,
  LIST_PING_RESULT
} from '@/paths'

export default {
  name: 'Home',
  components: {
    ListHTTPResult,
    ListDNSResult,
    ListTCPResult,
    ListPingResult
  },
  data () {
    return {
      limit: 5,
      simplify: true,
      moreHTTPResult: LIST_HTTP_RESULT,
      moreDNSResult: LIST_DNS_RESULT,
      moreTCPResult: LIST_TCP_RESULT,
      morePingResult: LIST_PING_RESULT
    }
  },
  computed: mapState({
    userAccount: state => state.user.account
  })
}
</script>
<style lang="sass" scoped>
@import '@/common'
.home
  margin: $mainMargin
.more
  font-size: 12px
  float: right
  color: $darkBlue
  i
    margin-left: 2px
    transform: rotate(90deg)
.detectResult
  margin-bottom: 20px
</style>
