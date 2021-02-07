<template lang="pug">
.httpResults: el-card
  template(
    #header
  )
    span HTTP检测结果列表
  result-filter(
    @filter="filter"
  )
  div(
    v-loading="https.processing"
  ): el-table(
    :data="https.items"
    row-key="id"
    stripe
    @sort-change="handleSortChange"
  )
    //- name
    el-table-column(
      label="名称"
      width="150"
      prop="name"
      key="name"
    )
    //- url
    el-table-column(
      label="URL"
      width="90"
    ): template(
      #default="scope"
    ): base-tooltip(
      icon="el-icon-info"
      :content="scope.row.url"
    )
    //- 状态
    el-table-column(
      label="状态"
      width="80"
    ): template(
      #default="scope"
    ): result-status(
      :status="scope.row.result"
    )
    //- 耗时
    el-table-column(
      label="耗时(ms)"
      key="maxDuration"
      prop="maxDuration"
      width="100"
    ) 
    //- 出错消息
    el-table-column(
      label="失败信息"
    ): template(
      #default="scope"
    ): ul
      li(
        v-for="item in scope.row.messages"
      ) {{item}}
    el-table-column(
      label="操作"
      width="120"
    ): template(
      #default="scope"
    ): HTTPDetail(
      :id="scope.row.id"
    )
    //- 时间
    el-table-column(
      label="时间"
      width="160"
    ): template(
      #default="scope"
    ): time-formater(
      :time="scope.row.updatedAt"
    )
  //- 分页
  el-pagination.pagination(
    layout="prev, pager, next, sizes"
    :current-page="currentPage"
    :page-size="query.limit"
    :page-sizes="pageSizes"
    :total="https.count"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  )
    
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useDetectorResultStore } from "../../store";
import BaseTooltip from "../../components/Tooltip.vue";
import TimeFormater from "../../components/TimeFormater.vue";
import ResultStatus from "./Status.vue";
import ResultFilter from "./Filter.vue";
import { PAGE_SIZES } from "../../constants/common";
import FilterTable from "../../mixins/FilterTable";
import HTTPDetail from "./HTTPDetail.vue";

export default defineComponent({
  name: "DetectorResultsHTTP",
  components: {
    ResultFilter,
    TimeFormater,
    BaseTooltip,
    ResultStatus,
    HTTPDetail,
  },
  mixins: [FilterTable],
  setup() {
    const detectorResultStore = useDetectorResultStore();
    return {
      https: detectorResultStore.state.https,
      list: (params) => detectorResultStore.dispatch("listHTTP", params),
    };
  },
  data() {
    return {
      pageSizes: PAGE_SIZES,
      query: {
        offset: 0,
        limit: PAGE_SIZES[0],
        order: "-updatedAt",
        fields: "updatedAt,task,result,maxDuration,url,messages",
      },
    };
  },
  beforeMount() {
    this.fetch();
  },
  methods: {
    async fetch() {
      const { https, query } = this;
      if (https.processing) {
        return;
      }
      const params = Object.assign({}, query);
      if (!params.result) {
        delete params.result;
      }
      if (params.duration) {
        params.duration = `${params.duration}s`;
      } else {
        delete params.duration;
      }
      try {
        await this.list(params);
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../../common";
.httpResults
  margin: $mainMargin
  i
    margin-right: 5px
  ul
    li
      list-style: inside
.selector, .submit
  width: 100%
.pagination
  text-align: right
  margin-top: 15px
</style>
