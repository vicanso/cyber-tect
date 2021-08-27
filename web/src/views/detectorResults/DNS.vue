<template lang="pug">
.dnsResults: el-card
  template(
    #header
  )
    span DNS检测结果列表
  result-filter(
    @filter="filter"
    category="dnses"
  )
  div(
    v-loading="dnses.processing"
  ): el-table(
    :data="dnses.items"
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
    //- host
    el-table-column(
      label="HOST"
      width="150"
      prop="host"
      key="host"
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
    //- 操作
    el-table-column(
      label="操作"
      width="120"
    ): template(
      #default="scope"
    ): DNSDetail(
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
    layout="prev, next, sizes"
    :current-page="currentPage"
    :page-size="query.limit"
    :page-sizes="pageSizes"
    :total="dnses.count"
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
import DNSDetail from "./DNSDetail.vue";

export default defineComponent({
  name: "DetectorResultsDNS",
  components: {
    ResultFilter,
    TimeFormater,
    BaseTooltip,
    ResultStatus,
    DNSDetail,
  },
  mixins: [FilterTable],
  setup() {
    const detectorResultStore = useDetectorResultStore();
    return {
      dnses: detectorResultStore.state.dnses,
      list: (params) => detectorResultStore.dispatch("listDNS", params),
    };
  },
  data() {
    return {
      pageSizes: PAGE_SIZES,
      query: {
        offset: 0,
        limit: PAGE_SIZES[0],
        order: "-updatedAt",
        fields: "updatedAt,task,result,maxDuration,host,messages",
      },
    };
  },
  beforeMount() {
    this.fetch();
  },
  methods: {
    async fetch() {
      const { dnses, query } = this;
      if (dnses.processing) {
        return;
      }
      const params = Object.assign({}, query);
      if (!params.result) {
        delete params.result;
      }
      // 如果未有指定筛选的task，则判断query中是否有指定
      if (!params.task && this.$route.query.task) {
        params.task = Number(this.$route.query.task);
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
.dnsResults
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
