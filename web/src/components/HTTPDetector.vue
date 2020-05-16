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
  CAT_HTTP
} from '@/constants/category'

export default {
  name: 'HTTPDetector',
  props: {
    originalDetector: Object
  },
  components: {
    BaseDetector
  },
  data () {
    return {
      category: CAT_HTTP
    }
  }
}
</script>
