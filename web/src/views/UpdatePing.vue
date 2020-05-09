<template lang="pug">
  el-card.updatePing
    div(
      slot="header"
    )
      span 修改Ping检测
    UpdateDetector(
      :category="category"
    )
</template>
<script>
import UpdateDetector from '@/components/UpdateDetector.vue'
import {
  CAT_PING
} from '@/constants/category'

export default {
  name: 'UpdatePing',
  components: {
    UpdateDetector
  },
  data () {
    return {
      category: CAT_PING
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.updatePing
  margin: $mainMargin
</style>
