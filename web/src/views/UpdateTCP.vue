<template lang="pug">
  el-card.updateTCP
    div(
      slot="header"
    )
      span 修改TCP检测
    UpdateDetector(
      :category="category"
    )
</template>
<script>
import UpdateDetector from '@/components/UpdateDetector.vue'
import {
  CAT_TCP
} from '@/constants/category'

export default {
  name: 'UpdateTCP',
  components: {
    UpdateDetector
  },
  data () {
    return {
      category: CAT_TCP
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.updateTCP
  margin: $mainMargin
</style>
