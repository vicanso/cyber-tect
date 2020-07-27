<template lang="pug">
el-card.addDNS
  div(
    slot="header"
  )
    span 添加DNS检测
  DNSDetector
</template>

<script>
import DNSDetector from "@/components/DNSDetector.vue";

export default {
  name: "AddDNS",
  components: {
    DNSDetector,
  },
};
</script>
<style lang="sass" scoped>
@import '@/common'
.addDNS
  margin: $mainMargin
</style>
