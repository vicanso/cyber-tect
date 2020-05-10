<template lang="pug">
  .httpDetectorResults(
    v-loading="processing"
  )
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
      pageSizes,
      currentPage: 1,
      query: {
        limit,
        order: '-updatedAt'
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
</style>
