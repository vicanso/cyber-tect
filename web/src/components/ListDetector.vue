<template lang="pug">
el-card.listDetector(
  v-loading='processing'
)
  div(
    slot="header"
  )
    | 服务检测
    span.category {{$props.category}}
  el-form(
    label-width="80px"
  )
    el-row(
      :gutter="20"
    )
      el-col(
        :span="6"
      )
        el-form-item(
          label="状态："
        )
          el-select(
            v-model="query.status"
          )
            el-option(
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            )
      el-col(
        :span="6"
      )
        el-form-item(
          label="拥有者："
        )
          el-checkbox(
            v-model="query.mine"
          ) 仅展示我的
      el-col(
        :span="12"
      )
        el-form-item(
          label-width="0"
        )
          el-button.submit(
            @click="search"
          ) 查询
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
        div(
          v-else-if="field.key === 'description'"
        )
          el-tooltip.tac(
            v-if="scope.row.description"
          )
            div(
              slot="content"
            ) {{scope.row.description}}
            i.el-icon-warning-outline
          span(
            v-else
          ) --
        span(v-else) {{scope.row[field.key]}}
    el-table-column(
      label="操作"
      width="160"
      fixed="right"
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
          a.op(
            href="#"
            @click.prevent="handlerView(scope.row.id)"
          )
            i.el-icon-view
            | 查看结果
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
  ROUTE_LIST_HTTP_DETECTOR_RESULT,
  ROUTE_UPDATE_DNS,
  ROUTE_LIST_DNS_DETECTOR_RESULT,
  ROUTE_UPDATE_TCP,
  ROUTE_LIST_TCP_DETECTOR_RESULT,
  ROUTE_UPDATE_PING,
  ROUTE_LIST_PING_DETECTOR_RESULT
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
  getPingListFields,
  getStatusOptions
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

function getListRoute (category) {
  switch (category) {
    case CAT_DNS:
      return ROUTE_LIST_DNS_DETECTOR_RESULT
    case CAT_HTTP:
      return ROUTE_LIST_HTTP_DETECTOR_RESULT
    case CAT_TCP:
      return ROUTE_LIST_TCP_DETECTOR_RESULT
    default:
      return ROUTE_LIST_PING_DETECTOR_RESULT
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
    const statusOptions = getStatusOptions()
    statusOptions.unshift({
      value: 0,
      label: '所有'
    })
    return {
      fields: getFields(category),
      updateRoute: getUpdateRoute(category),
      listRoute: getListRoute(category),
      pageSizes,
      query: {
        limit: pageSizes[0],
        offset: 0,
        order: '-updatedAt',
        status: 0,
        mine: false
      },
      statusOptions
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
    handlerView (id) {
      this.$router.push({
        name: this.listRoute,
        query: {
          task: `${id}`
        }
      })
    },
    search () {
      this.query.offset = 0
      this.fetch()
    },
    async fetch () {
      if (this.processing) {
        return
      }
      const {
        category
      } = this.$props
      try {
        const params = Object.assign({}, this.query)
        if (params.mine) {
          params.owner = this.userAccount
        }
        if (params.status === 0) {
          delete params.status
        }
        delete params.mine
        await this.listDetector({
          category,
          params
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
  margin-right: 10px
  &:last-child
    margin-right: 0
  i
    margin-right: 3px
.submit
  width: 100%
</style>
