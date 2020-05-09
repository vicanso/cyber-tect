<template lang="pug">
  .userInfo.grayBorder
    h3 用户信息
    el-form(
      ref="form"
      :model="form"
      label-width="80px"
      v-loading="submitting"
    )
      el-form-item(
        label="Email："
      )
        el-input(
          v-model="form.email"
          clearable
        )
      el-form-item
        el-button.submit(
          @click="onSubmit"
          type='primary'
        ) 更新
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'UserInfo',
  data () {
    return {
      submitting: false,
      form: {
        email: ''
      }
    }
  },
  methods: {
    ...mapActions([
      'updateUser'
    ]),
    async onSubmit () {
      const {
        email
      } = this.form
      if (this.submitting) {
        return
      }
      if (!email) {
        this.$message({
          message: '邮箱不能为空',
          type: 'warning'
        })
        return
      }
      this.submitting = true
      try {
        await this.updateUser({
          email
        })
        this.$message({
          message: '已成功更新信息',
          type: 'success'
        })
      } catch (err) {
        this.$message.error(err.message)
      } finally {
        this.submitting = false
      }
    }
  },
  computed: mapState({
    email: state => state.user.email
  }),
  mounted () {
    // 此方法有可能数据未加载完成，暂时使用
    this.form.email = this.email
  }
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
$userInfoWidth: 600px
.userInfo
  width: $userInfoWidth
  margin: 100px auto 30px
  background-color: $white
  padding: 30px 20px
h3
  line-height: 50px
  margin-bottom: 10px
  padding-left: 15px
.submit
  width: 100%
</style>
