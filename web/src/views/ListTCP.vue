<template lang="pug">
ListDetector(
  :category="category"
)
</template>
<script>
import ListDetector from '@/components/ListDetector.vue'
import {
  CAT_TCP
} from '@/constants/category'

export default {
  name: 'ListTCP',
  components: {
    ListDetector
  },
  data () {
    return {
      category: CAT_TCP
    }
  }
}
</script>
