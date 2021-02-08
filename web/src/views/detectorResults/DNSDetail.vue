<template lang="pug">
el-popover(
  placement="left-start"
  width="70%"
  trigger="click"
  @show="onShow"
)
  template(
    #reference
  ): el-button(
    type="text"
    size="small"
  ) 查看更多
  el-card(
    v-if="!processing && detail"
  )
    template(
      #header
    ) DNS检测：{{detail.host}}
    el-table(
      :data="detail.results"
    )
      //- DNS服务器
      el-table-column(
        prop="server"
        key="server"
        label="DNS服务器"
        width="150"
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
        prop="duration"
        key="duration"
        width="100"
        label="耗时(ms)"
      )
      //- IP列表
      el-table-column(
        label="IP列表"
        width="200"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="item in scope.row.ips"
        ) {{item}}
      //- 出错消息
      el-table-column(
        prop="message"
        key="message"
        label="出错消息"
      )
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useDetectorResultStore } from "../../store";
import ResultStatus from "./Status.vue";

export default defineComponent({
  name: "DetectorDNSResultDetail",
  components: {
    ResultStatus,
  },
  props: {
    id: {
      type: Number,
      default: 0,
    },
  },
  setup() {
    const detectorResultStore = useDetectorResultStore();
    return {
      getDNS: (id) => detectorResultStore.dispatch("getDNS", id),
    };
  },
  data() {
    return {
      processing: false,
      detail: null,
    };
  },
  methods: {
    async onShow(): Promise<void> {
      const { processing, detail } = this;
      if (processing || detail) {
        return;
      }
      try {
        this.processing = true;
        const data = await this.getDNS(this.$props.id);
        this.detail = data;
      } catch (err) {
        this.$error(err);
      } finally {
        this.processing = false;
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
ul
  li
    list-style inside
    line-height 18px
</style>
