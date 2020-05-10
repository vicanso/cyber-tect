<template lang="pug">
  el-card.listDetector(
    v-loading='processing'
  )
    div(
      slot="header"
    )
      | 服务检测
      span.category {{$props.category}}
    el-table(
      :data="detectors"
      row-key="id"
      stripe
    )
      el-table-column(
        v-for="field in fields"
        :key="field.key"
        :prop="field.key"
        :label="field.name"
        :width="field.width"
      )
        template(
          slot-scope="scope"
        )
          ul.receivers(
            v-if="field.key === 'receivers'"
          )
            li(
              v-for="account in scope.row[field.key]"
              :key="account"
            ) {{account}}
          span(v-else) {{scope.row[field.key]}}
      el-table-column(
        label="操作"
      )
        template(
          slot-scope="scope"
        )
          div(
            v-if="scope.row.owner == userAccount"
          )
            a.op(
              href="#"
              @click.prevent="handleUpdate(scope.row.id)"
            )
              i.el-icon-edit-outline
              | 更新
          div(
            v-else
          ) --
    .pagination: el-pagination(
      layout="prev, pager, next, sizes"
      :page-size="query.limit"
      :total="detectorCount"
      :page-sizes="pageSizes"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    )
</template>

<script>
import { mapState, mapActions } from 'vuex'

import {
  ROUTE_UPDATE_HTTP,
  ROUTE_UPDATE_DNS,
  ROUTE_UPDATE_TCP,
  ROUTE_UPDATE_PING
} from '@/router'
import {
  CAT_DNS,
  CAT_HTTP,
  CAT_TCP
} from '@/constants/category'
import {
  getDNSListFields,
  getHTTPListFields,
  getTCPListFields,
  getPingListFields
} from '@/helpers/field'

function getFields (category) {
  switch (category) {
    case CAT_DNS:
      return getDNSListFields()
    case CAT_HTTP:
      return getHTTPListFields()
    case CAT_TCP:
      return getTCPListFields()
    default:
      return getPingListFields()
  }
}

function getUpdateRoute (category) {
  switch (category) {
    case CAT_DNS:
      return ROUTE_UPDATE_DNS
    case CAT_HTTP:
      return ROUTE_UPDATE_HTTP
    case CAT_TCP:
      return ROUTE_UPDATE_TCP
    default:
      return ROUTE_UPDATE_PING
  }
}

export default {
  name: 'ListDetector',
  props: {
    category: {
      type: String,
      required: true
    }
  },
  data () {
    const {
      category
    } = this.$props
    const pageSizes = [
      10,
      20,
      30,
      50,
      100
    ]
    return {
      fields: getFields(category),
      updateRoute: getUpdateRoute(category),
      pageSizes,
      query: {
        limit: pageSizes[0],
        offset: 0,
        order: '-updatedAt'
      }
    }
  },
  computed: mapState({
    userAccount: state => state.user.account,
    processing: state => state.detector.processing,
    detectorCount: state => {
      if (!state.detector.currentCategory) {
        return
      }
      return state.detector[state.detector.currentCategory].count
    },
    detectors: state => {
      if (!state.detector.currentCategory) {
        return []
      }
      return state.detector.currentDetectors || []
    }
  }),
  methods: {
    ...mapActions([
      'listDetector'
    ]),
    handleCurrentChange (page) {
      this.query.offset = (page - 1) * this.query.limit
      this.fetch()
    },
    handleSizeChange (pageSize) {
      this.query.limit = pageSize
      this.query.offset = 0
      this.fetch()
    },
    handleUpdate (id) {
      this.$router.push({
        name: this.updateRoute,
        params: {
          id
        }
      })
    },
    async fetch () {
      if (this.processing) {
        return
      }
      const {
        category
      } = this.$props
      try {
        await this.listDetector({
          category,
          params: this.query
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
@import '@/common'
.listDetector
  margin: $mainMargin
.category
  margin-left: 3px
.pagination
  margin-top: 10px
  text-align: right
.receivers
  list-style: inside
.op
  color: $darkBlue
  i
    margin-right: 3px
</style>
