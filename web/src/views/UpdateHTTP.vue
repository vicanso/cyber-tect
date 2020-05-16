<template lang="pug">
el-card.updateHTTP
  div(
    slot="header"
  )
    span 修改HTTP检测
  UpdateDetector(
    :category="category"
  )
</template>

<script>
import UpdateDetector from '@/components/UpdateDetector.vue'
import {
  CAT_HTTP
} from '@/constants/category'

export default {
  name: 'UpdateHTTP',
  components: {
    UpdateDetector
  },
  data () {
    return {
      category: CAT_HTTP
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
.updateHTTP
  margin: $mainMargin
</style>
