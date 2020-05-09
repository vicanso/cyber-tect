<template lang="pug">
  .loginRegister.grayBorder
    h3 {{title}}
    el-form(
      v-loading="submitting"
      ref="form"
      :model="form"
      label-width="80px"
    )
      el-form-item(
        label="账号："
      )
        el-input(
          v-model="form.account"
          autofocus=true
          clearable
        )
      el-form-item(
        label="密码："
      )
        el-input(
          v-model="form.password"
          show-password
        )
      el-form-item(
        label="验证码："
      )
        el-row
          el-col(
            :span="18"
          )
            el-input.code(
              v-model="form.captcha"
              maxlength="4"
              clearable
            )
          el-col(
            :span="6"
          )
            .captchaImg(
              @click="getCaptcha"
            ): img(
              v-if="captchaData"
              :src="`data:image/jpeg;base64,${captchaData.data}`"
            )
      el-form-item
        el-button.submit(
          @click="onSubmit"
          type='primary'
        ) {{submitText}}
</template>

<script>
import { mapActions } from 'vuex'

import request from '@/request'
import {
  COMMONS_CAPTCHA
} from '@/constants/url'
import {
  ROUTE_LOGIN
} from '@/router'

const registerType = 'register'

export default {
  name: 'LoginRegister',
  props: {
    type: String
  },
  data () {
    const {
      type
    } = this.$props
    let title = '用户登录'
    let submitText = '立即登录'
    if (type === registerType) {
      title = '用户注册'
      submitText = '立即注册'
    }
    return {
      title,
      submitText,
      form: {
        account: '',
        password: '',
        captcha: ''
      },
      submitting: false,
      captchaData: null
    }
  },
  methods: {
    ...mapActions([
      'login',
      'register'
    ]),
    async getCaptcha () {
      this.captchaData = null
      try {
        const { data } = await request.get(COMMONS_CAPTCHA)
        this.captchaData = data
      } catch (err) {
        this.$message.error(err.message)
      }
    },
    async onSubmit () {
      const {
        account,
        password,
        captcha
      } = this.form
      if (!account || !password || !captcha) {
        this.$message({
          message: '账号、密码以及验证码不能为空',
          type: 'warning'
        })
        return
      }
      if (this.submitting) {
        return
      }
      this.submitting = true
      const params = {
        account,
        password,
        captcha: `${this.captchaData.id}:${captcha}`
      }
      try {
        if (this.$props.type === registerType) {
          await this.register(params)
          this.$router.replace({
            name: ROUTE_LOGIN
          })
        } else {
          await this.login(params)
          this.$router.back()
        }
      } catch (err) {
        // 图形验证码只可校验一次，因此出错则刷新
        this.getCaptcha()
        this.$message.error(err.message)
      } finally {
        this.submitting = false
      }
    }
  },
  mounted () {
    this.getCaptcha()
  }
}
</script>

<style lang="sass" scoped>
@import "@/common.sass"
$loginRegisterWidth: 600px
.loginRegister
  width: $loginRegisterWidth
  margin: 30px auto
  background-color: $white
  padding: 30px 20px
h3
  line-height: 50px
  margin-bottom: 10px
  padding-left: 15px
.el-input
  width: $loginRegisterWidth - 150
.captchaImg
  padding-left: 5px
  cursor: pointer
  overflow: hidden
  height: 40px
.code
  width: 100%
.submit
  width: 100%
</style>
