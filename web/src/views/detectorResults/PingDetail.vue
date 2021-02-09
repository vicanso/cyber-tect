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
    ) Ping检测：{{detail.ips}}
    el-table(
      :data="detail.results"
    )
      //- IP
      el-table-column(
        prop="ip"
        key="ip"
        label="IP"
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
  name: "DetectorPingResultDetail",
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
      getPing: (id) => detectorResultStore.dispatch("getPing", id),
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
        const data = await this.getPing(this.$props.id);
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
