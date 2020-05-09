<template lang="pug">
  .updateDetector(
    v-loading="fetching"
  )
    div(
      v-if="!fetching"
    )
      HTTPDetector(
        v-if="$props.category === httpCat"
        :originalDetector="originalDetector"
      )
      DNSDetector(
        v-else-if="$props.category === dnsCat"
        :originalDetector="originalDetector"
      )
      TCPDetector(
        v-else-if="$props.category === tcpCat"
        :originalDetector="originalDetector"
      )
      PingDetector(
        v-else
        :originalDetector="originalDetector"
      )
</template>

<script>
import { mapState, mapActions } from 'vuex'

import HTTPDetector from '@/components/HTTPDetector.vue'
import DNSDetector from '@/components/DNSDetector.vue'
import TCPDetector from '@/components/TCPDetector.vue'
import PingDetector from '@/components/PingDetector.vue'
import {
  CAT_HTTP,
  CAT_DNS,
  CAT_PING,
  CAT_TCP
} from '@/constants/category'

export default {
  name: 'UpdateDetector',
  props: {
    category: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      httpCat: CAT_HTTP,
      dnsCat: CAT_DNS,
      pingCat: CAT_PING,
      tcpCat: CAT_TCP
    }
  },
  components: {
    DNSDetector,
    HTTPDetector,
    TCPDetector,
    PingDetector
  },
  computed: mapState({
    fetching: state => state.detector.processing,
    originalDetector: state => state.detector.updateDetector
  }),
  methods: {
    ...mapActions([
      'clearUpdateDetector',
      'getUpdateDetector'
    ])
  },
  beforeDestroy () {
    this.clearUpdateDetector()
  },
  async mounted () {
    const {
      category
    } = this.$props
    try {
      await this.getUpdateDetector({
        category,
        id: Number(this.$route.params.id)
      })
    } catch (err) {
      this.$message.error(err.messagel)
    }
  }
}
</script>
