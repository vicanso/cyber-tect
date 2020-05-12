<template lang="pug">
  .dnsDetectorResults(
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
        prop="hostname"
        label="域名"
        width="200"
      )
      el-table-column(
        prop="server"
        label="DNS服务器"
        width="150"
      )
      el-table-column(
        label="解析结果"
      )
        template(
          slot-scope="scope"
        )
          ul(
            v-if="scope.row.ipAddrs"
          )
            li(
              v-for="ip in scope.row.ipAddrs"
            ) {{ip}}
          span(
            v-else
          ) --
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
        width="60"
      )
        template(
          slot-scope="scope"
        )
          router-link.op(
            title="查看任务"
            :to="{name: updateRoute, params: { id: scope.row.task }}"
          )
            i.el-icon-view
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
  CAT_DNS
} from '@/constants/category'
import DetectorResult from '@/components/DetectorResult.vue'
import DetectorMessage from '@/components/DetectorMessage.vue'
import {
  ROUTE_UPDATE_DNS
} from '@/router'

export default {
  name: 'ListDNSResult',
  props: {
    limit: {
      type: Number,
      default: 10
    },
    simplify: Boolean
  },
  components: {
    DetectorResult,
    DetectorMessage
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
    const query = Object.assign({
      limit,
      order: '-id'
    }, this.$route.query)
    return {
      updateRoute: ROUTE_UPDATE_DNS,
      pageSizes,
      currentPage: 1,
      query
    }
  },
  computed: mapState({
    processing: state => state.detector.dnsListResult.processing,
    count: state => state.detector.dnsListResult.count,
    results: state => state.detector.dnsListResult.results || []
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
          category: CAT_DNS,
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
.op
  margin-right: 20px
  color: $darkGray
  &:last-child
    margin-right: 0
</style>
