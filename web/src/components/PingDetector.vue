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
  CAT_PING
} from '@/constants/category'
export default {
  name: 'PingDetector',
  components: {
    BaseDetector
  },
  props: {
    originalDetector: Object
  },
  data () {
    return {
      category: CAT_PING
    }
  }
}
</script>
