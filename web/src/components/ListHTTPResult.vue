<template lang="pug">
.httpDetectorResults(
  v-loading="processing"
)
  FilterResult(
    :onFilter="filter"
    :task="query.task"
    :category="category"
  )
  el-dialog(
    title="更多信息"
    :visible.sync="showingDetail"
  )
    ul.detail(
      v-if="currentResult"
    )
      li
        span 状态码：
        | {{currentResult.statusCode}}
      li
        span IP地址：
        | {{currentResult.addrs.join(',')}}
      li(
        v-if="currentResult.message"
      )
        span 出错信息：
        | {{currentResult.message}}
      li
        span 协议：
        | {{currentResult.protocol}}
      li(
        v-if="currentResult.tlsVersion"
      )
        span TLS版本：
        | {{currentResult.tlsVersion}}
      li(
        v-if="currentResult.tlsCipherSuite"
      )
        span TLS加密套件：
        | {{currentResult.tlsCipherSuite}}
      li(
        v-if="currentResult.expirationDate"
      )
        span 证件有效期：
        | {{currentResult.expirationDate}}
      li.dnsNames(
        v-if="currentResult.certificateDNSNames && currentResult.certificateDNSNames.length"
      )
        span 证书域名：
        ul
          li(
            v-for="name in currentResult.certificateDNSNames"
          ) {{name}}
  el-table(
    :data="results"
    row-key="id"
    stripe
  )
    el-table-column(
      prop="id"
      label="ID"
      width="100"
    )
    el-table-column(
      prop="url"
      label="URL"
    )
    el-table-column(
      label="结果"
      width="60"
    )
      template(
        slot-scope="scope"
      )
        DetectorResult(
          :result="scope.row.result"
        )
    el-table-column(
      prop="durationDesc"
      label="耗时"
      width="350"
    )
      template(
        slot-scope="scope"
      )
        HTTPTimeline(
          :dnsLookup="scope.row.dnsLookup || 0"
          :tcpConnection="scope.row.tcpConnection || 0"
          :tlsHandshake="scope.row.tlsHandshake || 0"
          :serverProcessing="scope.row.serverProcessing || 0"
          :contentTransfer="scope.row.contentTransfer || 0"
          :duration="scope.row.duration || 0"
        )
    el-table-column(
      prop="message"
      label="出错信息"
      width="100"
    )
      template(
        slot-scope="scope"
      )
        DetectorMessage(
          :message="scope.row.message"
        )
    el-table-column(
      width="180"
      prop="updatedAtDesc"
      label="更新于"
    )
    el-table-column(
      label="操作"
      width="100"
    )
      template(
        slot-scope="scope"
      )
        el-link.op(
          @click="showDetail(scope.row)"
          title="查看详情"
        )
          i.el-icon-more
        router-link.op(
          title="更新任务"
          :to="{name: updateRoute, params: { id: scope.row.task }}"
        )
          i.el-icon-edit-outline
  .pagination(
    v-if="!simplify"
  ): el-pagination(
    layout="prev, pager, next, sizes"
    :page-size="query.limit"
    :total="count"
    :page-sizes="pageSizes"
    :current-page="currentPage"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  )
</template>
<script>
import { mapState } from "vuex";

import { CAT_HTTP } from "@/constants/category";
import HTTPTimeline from "@/components/HTTPTimeline.vue";
import DetectorResult from "@/components/DetectorResult.vue";
import DetectorMessage from "@/components/DetectorMessage.vue";
import FilterResult from "@/components/FilterResult.vue";
import BaseListResult from "@/components/BaseListResult.vue";
import { formatDate } from "@/helpers/util";
import { ROUTE_UPDATE_HTTP } from "@/router";

export default {
  name: "ListHTTPResult",
  props: {
    limit: {
      type: Number,
      default: 10,
    },
    simplify: Boolean,
  },
  extends: BaseListResult,
  components: {
    HTTPTimeline,
    DetectorMessage,
    DetectorResult,
    FilterResult,
  },
  data() {
    const { limit } = this.$props;
    const pageSizes = [10, 20, 30, 50, 100];
    if (!pageSizes.includes(limit)) {
      pageSizes.unshift(limit);
    }
    const query = Object.assign(
      {
        limit,
        order: "-id",
      },
      this.$route.query
    );
    return {
      category: CAT_HTTP,
      updateRoute: ROUTE_UPDATE_HTTP,
      showingDetail: false,
      currentResult: null,
      pageSizes,
      currentPage: 1,
      query,
    };
  },
  computed: mapState({
    processing: (state) => state.detector.httpListResult.processing,
    count: (state) => state.detector.httpListResult.count,
    results: (state) => state.detector.httpListResult.results || [],
  }),
  methods: {
    showDetail(data) {
      if (
        data.certificateExpirationDates &&
        data.certificateExpirationDates.length === 2
      ) {
        const start = formatDate(data.certificateExpirationDates[0]);
        const end = formatDate(data.certificateExpirationDates[1]);
        data.expirationDate = `${start} 至 ${end}`;
      }
      this.currentResult = data;
      this.showingDetail = true;
    },
  },
  mounted() {
    this.fetch();
  },
};
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.success
  color: $green
.fail
  color: $brown
.pagination
  margin-top: 10px
  text-align: right
.detail
  margin: 0
  padding: 0
  li
    list-style: none
    padding: 5px 0
    span
      display: inline-block
      width: 100px
      text-align: right
.dnsNames
  span
    float: left
  ul
    margin-left: 100px
    margin-top: -5px
.op
  margin-right: 20px
  color: $darkGray
  &:last-child
    margin-right: 0
</style>
