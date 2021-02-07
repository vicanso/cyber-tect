<template lang="pug">
.timeline
  span(
    v-for="block in blocks"
    :style="block"
  )
</template>

<script lang="ts">
import { defineComponent } from "vue";

const dnsColor = "#419488";
const tcpColor = "#f19c38";
const tlsColor = "#8f35aa";
const processingColor = "#5ac461";
const transferColor = "#49a8ee";

export default defineComponent({
  name: "HTTPTimeline",
  props: {
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
    const colors = [
      dnsColor,
      tcpColor,
      tlsColor,
      processingColor,
      transferColor,
    ];
    let duration = 0;
    values.forEach((value) => {
      duration += value;
    });

    const blocks = [];
    values.forEach((value, index) => {
      if (!value) {
        return;
      }
      blocks.push({
        width: `${(100 * value) / duration}%`,
        backgroundColor: colors[index],
      });
    });

    return {
      blocks,
    };
  },
});
</script>

<style lang="stylus" scoped>
.timeline span
  display inline-block
  height 15px
</style>
