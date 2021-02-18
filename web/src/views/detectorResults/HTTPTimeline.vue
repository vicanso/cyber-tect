<template lang="pug">
.timeline
  span(
    v-for="block in blocks"
    :style="block"
    :title="block.title"
  )
</template>

<script lang="ts">
import { defineComponent } from "vue";

const dnsColor = "#419488";
const tcpColor = "#f19c38";
const tlsColor = "#8f35aa";
const processingColor = "#5ac461";
const transferColor = "#49a8ee";
const timeoutColor = "#aaa";

export default defineComponent({
  name: "HTTPTimeline",
  props: {
    total: {
      type: Number,
      default: 1,
    },
    dns: {
      type: Number,
      default: 0,
    },
    tcp: {
      type: Number,
      default: 0,
    },
    tls: {
      type: Number,
      default: 0,
    },
    processing: {
      type: Number,
      default: 0,
    },
    transfer: {
      type: Number,
      default: 0,
    },
  },
  data() {
    const { $props } = this;
    const values = [
      $props.dns,
      $props.tcp,
      $props.tls,
      $props.processing,
      $props.transfer,
    ];
    const names = ["dns", "tcp", "tls", "processing", "transfer"];
    const colors = [
      dnsColor,
      tcpColor,
      tlsColor,
      processingColor,
      transferColor,
    ];
    const duration = $props.total;

    const blocks = [];
    let percent = 0;
    values.forEach((value, index) => {
      if (!value) {
        return;
      }
      const widthPercent = (100 * value) / duration;
      percent += widthPercent;
      blocks.push({
        title: `${names[index]}: ${value}ms`,
        width: `${widthPercent}%`,
        backgroundColor: colors[index],
      });
    });
    // 如果汇总时间过少，则添加超时的timeline
    if (percent < 95) {
      blocks.push({
        width: `${100 - percent}%`,
        backgroundColor: timeoutColor,
      });
    }

    return {
      blocks,
    };
  },
});
</script>

<style lang="stylus" scoped>
@import "../../common";

.timeline
  span
    display inline-block
    height 15px
    border-top 1px solid $dark
    border-bottom 1px solid $dark
    &:first-child
      border-left 1px solid $dark
    &:last-child
      border-right 1px solid $dark
</style>
