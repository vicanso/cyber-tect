<template lang="pug">
  ListDetector(
    :category="category"
  )
</template>
<script>
import ListDetector from '@/components/ListDetector.vue'
import {
  CAT_DNS
} from '@/constants/category'
export default {
  name: 'ListDNS',
  components: {
    ListDetector
  },
  data () {
    return {
      category: CAT_DNS
    }
  }
}
</script>
