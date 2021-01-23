<template lang="pug">
.loginRegister
  el-card
    //- 头部
    template(
      #header
    )
      .clearfix(
        slot="header"
      )
        i.el-icon-user-solid
        span {{title}}
    //- 输入表单
    el-form(
      ref="form"
      :model="form"
      label-width="80px"
    )
      //- 账号
      el-form-item(
        label="账号："
      ): el-input(
        placeholder="请输入账号"
        v-model="form.account"
        autofocus="true"
        clearable
      )
      //- 密码
      el-form-item(
        label="密码："
      ): el-input(
        v-model="form.password"
        show-password
        placeholder="请输入密码"
      )
      //- 验证码
      el-form-item(
        label="验证码："
      ): el-row
        //- 验证码输入框
        el-col(
          :span="18"
        ): el-input.code(
          v-model="form.captcha"
          maxlength="4"
          clearable
          @keyup.enter.native="onSubmit"
          placeholder="请输入验证码"
        )
        //- 验证码图片
        el-col(
          :span="6"
        ): .captcha(
          @click="refreshCaptcha"
        )
          img(
            v-if="captchaData"
            :src="`data:image/${captchaData.type};base64,${captchaData.data}`"
          )
          span(
            v-else
          ) ...
      //- 确认按钮
      el-form-item: ex-button(
        :onClick="onSubmit"
        :extra="exBtnExtra"
      ) {{submitText}}

</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useUserStore, useCommonStore } from "../store";
import { getLoginRouteName } from "../router";
import { LOGIN, REGISTER } from "../services/action";

const registerType = "register";

export default defineComponent({
  name: "LoginRegister",
  props: {
    type: {
      type: String,
      default: "login",
    },
  },
  setup() {
    const userStore = useUserStore();
    const commonStore = useCommonStore();
    return {
      user: userStore.state.info,
      getCaptcha: () => commonStore.dispatch("getCaptcha"),
      login: (params) => userStore.dispatch("login", params),
      register: (params) => userStore.dispatch("register", params),
    };
  },
  data() {
    const { type } = this.$props;
    let title = "用户登录";
    let submitText = "立即登录";
    let category = LOGIN;
    if (type === registerType) {
      title = "用户注册";
      submitText = "立即注册";
      category = REGISTER;
    }
    return {
      title,
      submitText,
      captchaData: null,
      exBtnExtra: {
        category,
      },
      form: {
        account: "",
        password: "",
        captcha: "",
      },
    };
  },
  mounted() {
    this.refreshCaptcha();
  },
  methods: {
    async refreshCaptcha() {
      try {
        this.captchaData = null;
        const data = await this.getCaptcha();
        this.captchaData = data;
      } catch (err) {
        this.$error(err);
      }
    },
    async onSubmit(): Promise<boolean> {
      let isSuccess = false;
      const { account, password, captcha } = this.form;
      if (!account || !password || !captcha) {
        this.$message.warning("账号、密码以及验证码不能为空");
        return isSuccess;
      }
      const params = {
        account,
        password,
        captcha: `${this.captchaData.id}:${captcha}`,
      };
      try {
        const { type } = this.$props;
        // 注册
        if (type == registerType) {
          await this.register(params);
          this.$router.replace({
            name: getLoginRouteName(),
          });
        } else {
          // 登录
          await this.login(params);
          this.$router.back();
        }
        isSuccess = true;
      } catch (err) {
        this.refreshCaptcha();
        this.$error(err);
      }
      return isSuccess;
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
.loginRegister
  margin: 100px auto
  max-width: 600px
  i
    margin-right: 5px
.captcha
  cursor: pointer
  overflow: hidden
  height: 40px
  text-align: center
.code
  width: 100%
.submit
  width: 100%
</style>
