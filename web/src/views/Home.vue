<template lang="pug">
.home
  el-row(
    v-if="userAccount"
    :gutter="20"
  )
    el-col(
      :span="12"
    ): el-card.detectResult
      div(
        slot="header"
      )
        | 我的最新HTTP检测
        router-link.more(
          :to="moreHTTPResult"
        )
          | 查看更多
          i.el-icon-more
      DetectorResultSummary(
        :category="http"
        :mime="true"
        :route="httpRoute"
      )
    el-col(
      :span="12"
    ): el-card.detectResult
      div(
        slot="header"
      )
        | 我的最新DNS检测
        router-link.more(
          :to="moreDNSResult"
        )
          | 查看更多
          i.el-icon-more
      DetectorResultSummary(
        :category="dns"
        :mime="true"
        :route="dnsRoute"
      )
    el-col(
      :span="12"
    ): el-card.detectResult
      div(
        slot="header"
      )
        | 我的最新TCP检测
        router-link.more(
          :to="moreTCPResult"
        )
          | 查看更多
          i.el-icon-more
      DetectorResultSummary(
        :category="tcp"
        :mime="true"
        :route="tcpRoute"
      )
    el-col(
      :span="12"
    ): el-card.detectResult
      div(
        slot="header"
      )
        | 我的最新Ping检测
        router-link.more(
          :to="morePingResult"
        )
          | 查看更多
          i.el-icon-more
      DetectorResultSummary(
        :category="ping"
        :mime="true"
        :route="pingRoute"
      )
  .tips(
    v-else
  ) 请先登录系统
</template>
<script>
import { mapState } from "vuex";

import DetectorResultSummary from "@/components/DetectorResultSummary.vue";
import {
  LIST_HTTP_RESULT,
  LIST_DNS_RESULT,
  LIST_TCP_RESULT,
  LIST_PING_RESULT,
} from "@/paths";
import {
  ROUTE_LIST_HTTP_DETECTOR_RESULT,
  ROUTE_LIST_DNS_DETECTOR_RESULT,
  ROUTE_LIST_PING_DETECTOR_RESULT,
  ROUTE_LIST_TCP_DETECTOR_RESULT,
} from "@/router";

import { CAT_HTTP, CAT_PING, CAT_DNS, CAT_TCP } from "@/constants/category";

export default {
  name: "Home",
  components: {
    DetectorResultSummary,
  },
  data() {
    return {
      http: CAT_HTTP,
      ping: CAT_PING,
      dns: CAT_DNS,
      tcp: CAT_TCP,
      httpRoute: ROUTE_LIST_HTTP_DETECTOR_RESULT,
      pingRoute: ROUTE_LIST_PING_DETECTOR_RESULT,
      dnsRoute: ROUTE_LIST_DNS_DETECTOR_RESULT,
      tcpRoute: ROUTE_LIST_TCP_DETECTOR_RESULT,
      moreHTTPResult: LIST_HTTP_RESULT,
      moreDNSResult: LIST_DNS_RESULT,
      moreTCPResult: LIST_TCP_RESULT,
      morePingResult: LIST_PING_RESULT,
    };
  },
  computed: mapState({
    userAccount: (state) => state.user.account,
  }),
};
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
.tips
  text-align: center
  margin-top: 100px
</style>
