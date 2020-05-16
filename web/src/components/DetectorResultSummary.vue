<template lang="pug">
.detectorResultSummary.clearfix(
  v-loading="processing"
)
  template(
    v-if="detectorResults && detectorResults.length !== 0"
  )
    el-tooltip(
      v-for="item in detectorResults"
      :key="item.id"
    )
      div(
        slot="content"
        v-html="formatContent(item)"
      )
      .summary(
        :class="{success: item.result === resultSucess, fail: item.result !== resultSucess}"
      )
  .tips(
    v-else
  )
    i.el-icon-warning-outline
    | 请先设置自定义的检测配置
</template>
<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'DetectorResultSummary',
  props: {
    category: {
      type: String,
      required: true
    },
    mime: Boolean
  },
  data () {
    return {
      filterTasks: ''
    }
  },
  computed: mapState({
    userAccount: state => state.user.account,
    resultSucess: state => state.detector.resultSucess,
    resultFail: state => state.detector.resultFail,
    processing: function (state) {
      const {
        category
      } = this.$props
      const {
        detector
      } = state
      return detector.mime[category].processing || detector[`${category}ListResult`].processing
    },
    detectors: function (state) {
      return state.detector.mime[this.$props.category].detectors
    },
    detectorResults: function (state) {
      return state.detector[`${this.$props.category}ListResult`].results
    }
  }),
  methods: {
    ...mapActions([
      'resetDetectorResults',
      'listMimeDetector',
      'listDetectorResult'
    ]),
    formatContent (item) {
      let content = `任务：${item.task}<br />
      结果：${item.updatedAtDesc} 检测${item.resultDesc}<br />
        耗时：${item.durationDesc}
      `
      if (item.result === this.resultSucess) {
        return content
      }
      content += `<br />
      失败原因：${item.message}
      `
      return content
    },
    async fetchMimeDetector () {
      const {
        category
      } = this.$props
      const params = {
        // 偷懒，直接取前100个
        limit: 100,
        offset: 0,
        owner: this.userAccount,
        order: '-id',
        fields: 'id',
        status: 1
      }
      try {
        await this.listMimeDetector({
          category,
          params
        })
        const arr = this.detectors.map(item => item.id)
        if (arr.length === 0) {
          return
        }
        this.filterTasks = arr.join(',')
        await this.fetchDetectorSummary()
      } catch (err) {
        this.$message.error(err.error)
      }
    },
    async fetchDetectorSummary () {
      const {
        category
      } = this.$props
      const params = {
        fields: 'result,id,message,task,duration,updatedAt',
        limit: 120,
        offset: 0,
        order: '-id'
      }
      if (this.filterTasks) {
        params.tasks = this.filterTasks
      }
      try {
        await this.listDetectorResult({
          category,
          params
        })
      } catch (err) {
        this.$message.error(err.message)
      }
    }
  },
  mounted () {
    const {
      mime,
      category
    } = this.$props
    this.resetDetectorResults({
      category
    })
    if (mime) {
      this.fetchMimeDetector()
    } else {
      this.fetchDetectorSummary()
    }
  }
}
</script>
<style lang="sass" scoped>
@import '@/common'
$summaryWidth: 30px
.detectorResultSummary
  padding: 15px
  padding-right: 0
.summary
  width: $summaryWidth
  height: $summaryWidth
  margin: 0 2px 2px 0
  float: left
  cursor: pointer
  &.success
    background-color: #26a0fc
  &.fail
    background-color: #c1c3c7
.tips
  font-size: 13px
  i
    margin-right: 3px
</style>
