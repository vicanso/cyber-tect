<template lang="pug">
el-card.updateDNS
  div(
    slot="header"
  )
    span 修改DNS检测
  UpdateDetector(
    :category="category"
  )
</template>
<script>
import UpdateDetector from '@/components/UpdateDetector.vue'
import {
  CAT_DNS
} from '@/constants/category'

export default {
  name: 'UpdateDNS',
  components: {
    UpdateDetector
  },
  data () {
    return {
      category: CAT_DNS
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.updateDNS
  margin: $mainMargin
</style>
