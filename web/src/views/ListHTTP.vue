<template lang="pug">
.listHTTP
  ListDetector(
    :category="category"
  )
  .add
    el-button(
      type="primary"
      @click="add"
    ) 添加HTTP检测
</template>
<script>
import ListDetector from '@/components/ListDetector.vue'
import {
  CAT_HTTP
} from '@/constants/category'
import {
  ROUTE_ADD_HTTP
} from '@/router'

export default {
  name: 'ListHTTP',
  components: {
    ListDetector
  },
  data () {
    return {
      category: CAT_HTTP
    }
  },
  methods: {
    add () {
      this.$router.push({
        name: ROUTE_ADD_HTTP
      })
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common.sass'
.add
  margin: 0 $mainMargin $mainMargin $mainMargin
  .el-button
    width: 100%
</style>
