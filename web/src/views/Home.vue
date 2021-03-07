<template lang="pug">
.home
  el-row(
    :gutter="20"
  )
    el-col.graph(
      :span="12"
    ): el-card
      template(
        #header
      )
        router-link.more(
          :to="{ name: httpResultRoute}"
        )
          span 查看更多
          i.el-icon-more
        span 最新HTTP检测
      .result(
        v-loading="https.processing"
      )
        result-summary(
          :route="httpResultRoute"
          v-if="!https.processing"
          :results="https.items"
        )

    el-col.graph(
      :span="12"
    ): el-card
      template(
        #header
      )
        router-link.more(
          :to="{ name: dnsResultRoute}"
        )
          span 查看更多
          i.el-icon-more
        span 最新DNS检测
      .result(
        v-loading="dnses.processing"
      )
        result-summary(
          :route="dnsResultRoute"
          v-if="!dnses.processing"
          :results="dnses.items"
        )

    el-col.graph(
      :span="12"
    ): el-card
      template(
        #header
      )
        router-link.more(
          :to="{ name: tcpResultRoute}"
        )
          span 查看更多
          i.el-icon-more
        span 最新TCP检测
      .result(
        v-loading="tcps.processing"
      )
        result-summary(
          :route="tcpResultRoute"
          v-if="!tcps.processing"
          :results="tcps.items"
        )

    el-col.graph(
      :span="12"
    ): el-card
      template(
        #header
      )
        router-link.more(
          :to="{ name: pingResultRoute}"
        )
          span 查看更多
          i.el-icon-more
        span 最新Ping检测
      .result(
        v-loading="pings.processing"
      )
        result-summary(
          :route="pingResultRoute"
          v-if="!pings.processing"
          :results="pings.items"
        )
</template>

<script lang="ts">
import { defineComponent } from "vue";
import {
  getDetectorResultHTTPRouteName,
  getDetectorResultDNSRouteName,
  getDetectorResultTCPRouteName,
  getDetectorResultPingRouteName,
} from "../router";

import { useDetectorResultStore } from "../store";
import ResultSummary from "../components/detectors/ResultSummary.vue";

import { getSetting } from "../services/setting";

export default defineComponent({
  name: "Home",
  components: {
    ResultSummary,
  },
  setup() {
    const detectorResultStore = useDetectorResultStore();
    return {
      httpResultRoute: getDetectorResultHTTPRouteName(),
      dnsResultRoute: getDetectorResultDNSRouteName(),
      tcpResultRoute: getDetectorResultTCPRouteName(),
      pingResultRoute: getDetectorResultPingRouteName(),
      https: detectorResultStore.state.https,
      dnses: detectorResultStore.state.dnses,
      pings: detectorResultStore.state.pings,
      tcps: detectorResultStore.state.tcps,
      listHTTP: (params) => detectorResultStore.dispatch("listHTTP", params),
      listDNS: (params) => detectorResultStore.dispatch("listDNS", params),
      listTCP: (params) => detectorResultStore.dispatch("listTCP", params),
      listPing: (params) => detectorResultStore.dispatch("listPing", params),
    };
  },
  data() {
    return {
      querySize: null,
      query: {
        // 所有结果均返回
        result: "",
        ignoreCount: true,
        order: "-updatedAt",
        offset: 0,
        limit: 30 * 3,
        fields: "updatedAt,id,task,result,maxDuration,messages",
      },
    };
  },
  beforeMount() {
    const setting = getSetting();
    if (setting.mainDetectorResultCount) {
      this.query.limit = setting.mainDetectorResultCount;
      this.querySize = setting.mainDetectorResultCount;
    }
    if (setting.mainDetectorOnlyFailure) {
      // 只返回失败结果
      this.query.result = "2"; 
    }
    this.query.timeRange = setting.mainDetectorTimeRange;
    this.fetch();
    if (setting.mainDetectorRefreshInterval) {
      this.timer = setInterval(
        () => this.fetch(),
        setting.mainDetectorRefreshInterval * 1000
      );
    }
  },
  beforeUnmount() {
    clearInterval(this.timer);
  },
  methods: {
    fetch() {
      this.fetchHTTP();
      this.fetchDNS();
      this.fetchTCP();
      this.fetchPing();
    },
    async fetchHTTP() {
      const { query } = this;
      try {
        await this.listHTTP(query);
      } catch (err) {
        this.$error(err);
      }
    },
    async fetchDNS() {
      const { query } = this;
      try {
        await this.listDNS(query);
      } catch (err) {
        this.$error(err);
      }
    },
    async fetchTCP() {
      const { query } = this;
      try {
        await this.listTCP(query);
      } catch (err) {
        this.$error(err);
      }
    },
    async fetchPing() {
      const { query } = this;
      try {
        await this.listPing(query);
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";

.home
  margin: $mainMargin
.more
  float: right
  font-size: 12px
  color: #1989fa
  i
    margin-left: 2px
    transform: rotate(90deg)
.graph
  margin-bottom: 20px
</style>
