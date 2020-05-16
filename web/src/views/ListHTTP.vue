<template lang="pug">
ListDetector(
  :category="category"
)
</template>
<script>
import ListDetector from '@/components/ListDetector.vue'
import {
  CAT_HTTP
} from '@/constants/category'

export default {
  name: 'ListHTTP',
  components: {
    ListDetector
  },
  data () {
    return {
      category: CAT_HTTP
    }
  }
}
</script>
