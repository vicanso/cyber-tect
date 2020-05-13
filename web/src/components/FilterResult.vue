<template lang="pug">
.filterResult
  el-form(
    ref="form"
    :model="form"
    label-width="100px"
  )
    el-row(
      :gutter="20"
    )
      el-col(
        :span="6"
      )
        el-form-item(
          label="任务ID："
        )
          el-input(
            clearable
            v-model="form.task"
            placeholder="请输入要查看的任务ID"
          )
      el-col(
        :span="6"
      )
        el-form-item(
          label="状态："
        )
          el-select.selector(
            v-model="form.result"
          )
            el-option(
              v-for="item in results"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            )
      el-col(
        :span="6"
      )
        el-form-item(
          label="耗时大于："
        )
          el-input(
            clearable
            v-model="form.duration"
            placeholder="请输入时间"
          )
            template(
              slot="append"
            ) 秒
      el-col(
        :span="6"
      )
        el-button.filter(
          @click="filter"
        ) 查询
</template>
<script>
export default {
  name: 'FilterResult',
  props: {
    onFilter: {
      type: Function,
      required: true
    },
    task: String
  },
  data () {
    const {
      task
    } = this.$props
    return {
      form: {
        task,
        result: 0,
        duration: ''
      },
      results: [
        {
          label: '所有',
          value: 0
        },
        {
          label: '成功',
          value: 1
        },
        {
          label: '失败',
          value: 2
        }
      ]
    }
  },
  methods: {
    filter () {
      const {
        task,
        result,
        duration
      } = this.form
      const params = {
        task: '',
        result: null,
        duration: null
      }
      if (task) {
        params.task = task
      }
      if (result) {
        params.result = result
      }
      if (duration) {
        params.duration = `${duration}s`
      }
      this.$props.onFilter(params)
    }
  }
}
</script>
<style lang="sass" scoped>
.selector
  width: 100%
.filter
  width: 100%
</style>
