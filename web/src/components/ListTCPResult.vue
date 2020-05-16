<template lang="pug">
  .tcpDetectorResults(
    v-loading="processing"
  )
    FilterResult(
      :onFilter="filter"
      :task="query.task"
      :category="category"
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
        prop="ip"
        label="IP"
        width="150"
      )
      el-table-column(
        prop="port"
        label="端口"
        width="150"
      )
      el-table-column(
        prop="network"
        label="网络类型"
        width="100"
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
        width="100"
      )
      el-table-column(
        prop="message"
        label="出错信息"
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
import { mapState } from 'vuex'

import {
  CAT_TCP
} from '@/constants/category'
import DetectorResult from '@/components/DetectorResult.vue'
import DetectorMessage from '@/components/DetectorMessage.vue'
import FilterResult from '@/components/FilterResult.vue'
import BaseListResult from '@/components/BaseListResult.vue'
import {
  ROUTE_UPDATE_TCP
} from '@/router'

export default {
  name: 'ListTCPResult',
  props: {
    limit: {
      type: Number,
      default: 10
    },
    simplify: Boolean
  },
  extends: BaseListResult,
  components: {
    DetectorResult,
    DetectorMessage,
    FilterResult
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
      category: CAT_TCP,
      updateRoute: ROUTE_UPDATE_TCP,
      pageSizes,
      currentPage: 1,
      query
    }
  },
  computed: mapState({
    processing: state => state.detector.tcpListResult.processing,
    count: state => state.detector.tcpListResult.count,
    results: state => state.detector.tcpListResult.results || []
  }),
  mounted () {
    this.fetch()
  }
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.pagination
  margin-top: 10px
  text-align: right
.op
  margin-right: 20px
  color: $darkGray
  &:last-child
    margin-right: 0
</style>
