<template lang="pug">
.detectorResult
  span(
    v-if="result === 1"
  ).success
    i.el-icon-circle-check
  span(
    v-else
  ).fail
    i.el-icon-circle-close
</template>
<script>
export default {
  name: 'DetectorResult',
  props: {
    result: {
      type: Number,
      required: true
    }
  }
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.success
  color: $green
.fail
  color: $brown
</style>
