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
    ) HTTP检测：{{detail.url}}
    el-table(
      :data="detail.results"
      stripe
    )
      el-table-column(
        prop="addr"
        key="addr"
        label="地址"
        fixed="left"
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
      //- 出错信息
      el-table-column(
        label="出错信息"  
        width="80"
      ): template(
        #default="scope"
      ): base-tooltip(
        icon="el-icon-info"
        :content="scope.row.message"
      )
      //- 耗时
      el-table-column(
        width="150"
        label="耗时(ms)"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="item in getDurationDesc(scope)"
        ) {{`${item.name}: ${item.duration}`}}
      //- timeline
      el-table-column(
        label="时间线"
        width="150"
      ): template(
        #default="scope"
      ): HTTPTimeline(
        :dns="scope.row.dnsLookup"
        :tcp="scope.row.tcpConnection"
        :tls="scope.row.tlsHandshake"
        :processing="scope.row.serverProcessing"
        :transfer="scope.row.contentTransfer"
      )
      //- HTTP协议
      el-table-column(
        label="HTTP协议"
        prop="protocol"
        key="protocol"
        width="90"
      )
      //- TLS版本
      el-table-column(
        label="TLS"
        prop="tlsVersion"
        key="tlsVersion"
        width="80"
      )
      //- TLS加密套件
      el-table-column(
        label="TLS加密"
        prop="tlsCipherSuite"
        key="tlsCipherSuite"
        width="250"
      )
      //- 证书域名
      el-table-column(
        label="证书域名"
        width="150"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="item in scope.row.certificateDNSNames"
        ) {{item}}
      //- 证书有效期
      el-table-column(
        label="证书有效期"
        width="180"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="item in scope.row.certificateExpirationDates"
        ): time-formater(
          :time="item"
        )
  p(
    v-else
  ) 正在加载中，请稍候...
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useDetectorResultStore } from "../../store";
import ResultStatus from "./Status.vue";
import HTTPTimeline from "./HTTPTimeline.vue";
import TimeFormater from "../../components/TimeFormater.vue";
import BaseTooltip from "../../components/Tooltip.vue";

interface DurationDesc {
  name: string;
  duration: number;
}

export default defineComponent({
  name: "DetectorHTTPResultDetail",
  components: {
    ResultStatus,
    TimeFormater,
    HTTPTimeline,
    BaseTooltip,
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
      getHTTP: (id) => detectorResultStore.dispatch("getHTTP", id),
    };
  },
  data() {
    return {
      processing: false,
      detail: null,
    };
  },
  methods: {
    getDurationDesc(result: Record<string, unknown>): DurationDesc[] {
      const total: number = result.row.duration || 0;
      const arr: DurationDesc[] = [];
      if (total) {
        arr.push({
          name: "total",
          duration: total,
        });
      }
      const names = {
        dnsLookup: "dns",
        tcpConnection: "tcp",
        tlsHandshake: "tls",
        serverProcessing: "processing",
        contentTransfer: "transfer",
      };
      [
        "dnsLookup",
        "tcpConnection",
        "tlsHandshake",
        "serverProcessing",
        "contentTransfer",
      ].forEach((key) => {
        const value = result.row[key];
        if (value) {
          const duration: number = value;
          const desc: DurationDesc = {
            name: names[key],
            duration,
          };
          arr.push(desc);
        }
      });
      return arr;
    },
    async onShow(): Promise<void> {
      const { processing, detail } = this;
      if (processing || detail) {
        return;
      }
      try {
        this.processing = true;
        const data = await this.getHTTP(this.$props.id);
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
