<template lang="pug">
  .httpDetectorResults(
    v-loading="processing"
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
          span(
            v-if="scope.row.result === 1"
          ).success
            i.el-icon-circle-check
          span(
            v-else
          ).fail
            i.el-icon-circle-close
      el-table-column(
        prop="durationDesc"
        label="耗时"
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
          el-tooltip(
            v-if="scope.row.message"
          )
            div(
              slot="content"
            ) {{scope.row.message}}
            i.el-icon-warning-outline
          span(
            v-else
          ) --
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
          el-link(
            @click="showDetail(scope.row)"
          )
            i.el-icon-more
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
import { mapActions, mapState } from 'vuex'

import {
  CAT_HTTP
} from '@/constants/category'
import HTTPTimeline from '@/components/HTTPTimeline.vue'
import {
  formatDate
} from '@/helpers/util'

export default {
  name: 'ListHTTPResult',
  props: {
    limit: {
      type: Number,
      default: 10
    },
    simplify: Boolean
  },
  components: {
    HTTPTimeline
  },
  data () {
    const {
      limit
    } = this.$props
    const pageSizes = [
      10,
      20,
      30,
      50,
      100
    ]
    if (!pageSizes.includes(limit)) {
      pageSizes.unshift(limit)
    }
    return {
      showingDetail: false,
      currentResult: null,
      pageSizes,
      currentPage: 1,
      query: {
        limit,
        order: '-id'
      }
    }
  },
  computed: mapState({
    processing: state => state.detector.httpListResult.processing,
    count: state => state.detector.httpListResult.count,
    results: state => state.detector.httpListResult.results || []
  }),
  methods: {
    ...mapActions([
      'listDetectorResult'
    ]),
    handleSizeChange (pageSize) {
      this.query.limit = pageSize
      this.currentPage = 1
      this.fetch()
    },
    handleCurrentChange (page) {
      this.currentPage = page
      this.fetch()
    },
    async fetch () {
      const {
        query,
        currentPage
      } = this
      try {
        await this.listDetectorResult({
          category: CAT_HTTP,
          params: Object.assign({
            offset: (currentPage - 1) * query.limit
          }, query)
        })
      } catch (err) {
        this.$message.error(err.message)
      }
    },
    showDetail (data) {
      if (data.certificateExpirationDates && data.certificateExpirationDates.length === 2) {
        const start = formatDate(data.certificateExpirationDates[0])
        const end = formatDate(data.certificateExpirationDates[1])
        data.expirationDate = `${start} 至 ${end}`
      }
      this.currentResult = data
      this.showingDetail = true
    }
  },
  mounted () {
    this.fetch()
  }
}
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
</style>
