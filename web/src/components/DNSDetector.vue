<template lang="pug">
BaseDetector(
  :category="category"
  :originalDetector="$props.originalDetector"
  labelWidth="100px"
)
</template>
<script>
import BaseDetector from '@/components/BaseDetector.vue'
import {
  CAT_DNS
} from '@/constants/category'
export default {
  name: 'DNSDetector',
  components: {
    BaseDetector
  },
  props: {
    originalDetector: Object
  },
  data () {
    return {
      category: CAT_DNS
    }
  }
}
</script>
