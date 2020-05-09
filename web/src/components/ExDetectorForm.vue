<template lang="pug">
  el-form(
    ref="form"
    :model="form"
    :label-width="$props.labelWidth"
  )
    el-row(
      :gutter="$props.gutter"
    )
      el-col(
        v-for="field in $props.fields"
        :key="field.name"
        :span="field.span || 12"
      )
        el-form-item(
          :label="field.label + '：'"
        )
          el-select.selector(
            v-if="field.type === 'select'"
            v-model="form[field.name]"
          )
            el-option(
              v-for="opt in field.options"
              :key="opt.value"
              :label="opt.label"
              :value="opt.value"
            )
          el-input(
            v-else-if="field.type === 'textarea'"
            type="textarea"
            v-model="form[field.name]"
            :placeholder="field.placeholder"
            :rows="5"
          )
          el-select.selector(
            v-else-if="field.type === 'users'"
            v-model="form[field.name]"
            multiple
            filterable
            remote
            :placeholder="field.placeholder"
            :remote-method="searchUser"
            :loading="fetchingUsers"
          )
            el-option(
              v-for="item in users"
              :key="item.account"
              :label="item.account"
              :value="item.account"
            )
          el-input(
            v-else
            v-model="form[field.name]"
            :type="field.inputType"
            :placeholder="field.placeholder"
            clearable
          )
      el-col(
        :span="12"
      )
        el-form-item
          el-button.submit(
            @click="onSubmit"
            type='primary'
          ) 提交
      el-col(
        :span="12"
      )
        el-form-item
          el-button.back(
            @click="handleBack"
          ) 返回

</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'ExDetectorForm',
  props: {
    originalDetector: Object,
    fields: {
      type: Array,
      required: true
    },
    labelWidth: {
      type: String,
      default: '80px'
    },
    gutter: {
      type: Number,
      default: 20
    },
    submit: {
      type: Function,
      required: true
    },
    back: Function
  },
  data () {
    const {
      fields,
      originalDetector
    } = this.$props
    const data = {
      form: {}
    }
    fields.forEach((item) => {
      const {
        name
      } = item
      if (originalDetector) {
        data.form[name] = originalDetector[name]
      } else {
        data.form[name] = null
      }
    })
    return data
  },
  computed: mapState({
    users: state => state.user.users,
    fetchingUsers: state => state.user.fetchingUsers
  }),
  methods: {
    ...mapActions([
      'listUser'
    ]),
    async searchUser (keyword) {
      try {
        await this.listUser({
          keyword,
          limit: 10
        })
      } catch (err) {
        this.$message.error(err.message)
      }
    },
    handleBack () {
      if (this.$props.back) {
        this.$props.back()
        return
      }
      this.$router.back()
    },
    onSubmit () {
      const {
        form
      } = this
      const {
        fields
      } = this.$props
      const data = {}
      Object.keys(form).forEach((key) => {
        const value = form[key]
        if (value !== null && (!Array.isArray(value) || value.length !== 0)) {
          data[key] = value
        }
      })
      Object.keys(data).forEach((key) => {
        fields.forEach((field) => {
          if (field.name === key && field.inputType === 'number') {
            data[key] = Number(data[key])
          }
        })
      })
      const emptyFields = []
      this.$props.fields.forEach((item) => {
        if (item.required && !data[item.name]) {
          emptyFields.push(item.label)
        }
      })
      if (emptyFields.length !== 0) {
        this.$message({
          type: 'warning',
          message: `${emptyFields.join('，')}不能为空`
        })
        return
      }
      this.$props.submit(data)
    }
  },
  mounted () {
    this.searchUser()
  }
}
</script>
<style lang="sass" scoped>
.selector, .submit, .back
  width: 100%
</style>
