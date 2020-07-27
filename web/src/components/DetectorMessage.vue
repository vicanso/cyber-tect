<template lang="pug">
.detectorMessage
  el-tooltip(
    v-if="message"
  )
    div(
      slot="content"
    ) {{message}}
    i.el-icon-warning-outline
  span(
    v-else
  ) --
</template>
<script>
export default {
  name: "DetectorMessage",
  props: {
    message: String,
  },
};
</script>
