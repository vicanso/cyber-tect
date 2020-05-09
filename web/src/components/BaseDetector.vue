<template lang="pug">
  ExDetectorForm(
    :fields="fields"
    :submit="submit"
    :originalDetector="$props.originalDetector"
    :labelWidth="$props.labelWidth"
    v-loading="processing"
  )
</template>

<script>
import { mapState, mapActions } from 'vuex'

import ExDetectorForm from '@/components/ExDetectorForm.vue'
import {
  diff
} from '@/helpers/util'
import {
  getHTTPFields,
  getDNSFields,
  getTCPFields,
  getPingFields
} from '@/helpers/field'
import {
  ROUTE_LIST_HTTP,
  ROUTE_LIST_DNS,
  ROUTE_LIST_PING,
  ROUTE_LIST_TCP
} from '@/router'
import {
  CAT_HTTP,
  CAT_PING,
  CAT_DNS,
  CAT_TCP
} from '@/constants/category'

function getFields (category) {
  switch (category) {
    case CAT_DNS:
      return getDNSFields()
    case CAT_HTTP:
      return getHTTPFields()
    case CAT_TCP:
      return getTCPFields()
    default:
      return getPingFields()
  }
}

function getRouterName (category) {
  let routerName = ''
  switch (category) {
    case CAT_HTTP:
      routerName = ROUTE_LIST_HTTP
      break
    case CAT_DNS:
      routerName = ROUTE_LIST_DNS
      break
    case CAT_PING:
      routerName = ROUTE_LIST_PING
      break
    default:
      routerName = ROUTE_LIST_TCP
      break
  }
  return routerName
}

export default {
  name: 'HTTPDetector',
  components: {
    ExDetectorForm
  },
  props: {
    originalDetector: Object,
    labelWidth: String,
    category: {
      type: String,
      required: true
    }
  },
  data () {
    const {
      category
    } = this.$props
    return {
      routerName: getRouterName(category),
      fields: getFields(category)
    }
  },
  computed: mapState({
    processing: state => state.detector.processing
  }),
  methods: {
    ...mapActions([
      'addDetector',
      'updateDetector'
    ]),
    async add (detector) {
      const {
        category
      } = this.$props
      try {
        await this.addDetector({
          category,
          detector
        })
        this.$message({
          message: `添加${category.toUpperCase()}检测成功`
        })
        this.$router.push({
          name: this.routerName
        })
      } catch (err) {
        this.$message.error(err.message)
      }
    },
    async update (detector) {
      const {
        originalDetector,
        category
      } = this.$props
      const {
        data,
        modifiedCount
      } = diff(detector, originalDetector)
      if (modifiedCount === 0) {
        this.$message({
          message: '相关配置未修改，请修改后再提交'
        })
        return
      }
      try {
        await this.updateDetector({
          id: originalDetector.id,
          category,
          data
        })
        this.$router.back()
      } catch (err) {
        this.$message.error(err.message)
      }
    },
    submit (detector) {
      if (this.$props.originalDetector) {
        this.update(detector)
        return
      }
      this.add(detector)
    }
  }
}
</script>
