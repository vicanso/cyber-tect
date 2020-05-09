<template lang="pug">
  ListDetector(
    :category="category"
  )
</template>
<script>
import ListDetector from '@/components/ListDetector.vue'
import {
  CAT_PING
} from '@/constants/category'

export default {
  name: 'ListPing',
  components: {
    ListDetector
  },
  data () {
    return {
      category: CAT_PING
    }
  }
}
</script>
