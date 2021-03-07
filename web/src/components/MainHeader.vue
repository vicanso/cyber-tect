<template lang="pug">
header.header
  //- 用户信息
  .userInfo
    span(
      v-if="user.processing"
    ) 正在加载...
    .functions(
      v-else-if="user.account"
    )
      el-popover(
        placement="bottom-end"
        :width="250"
        trigger="click"
      )
        template(
          #reference
        ): el-button.setting(
          type="text"
          size="small"
        ) 首页设置
        el-form
          el-form-item(
            label="拉取数据："
          ): el-input(
            placeholder="请输入每个分类展示的检测数量"
            v-model="querySize"
            type="number"
          )
          el-form-item(
            label="定时刷新："
          ): el-input(
            placeholder="请输入定时刷新间隔，单位秒"
            v-model="refreshInterval"
            type="number"
          )
          el-form-item(
            label="结果筛选："
          ): el-checkbox(
            v-model="onlyFailure"
          ) 仅展示失败
          el-form-item(
            label="查询区间："
          ): el-input(
            placeholder="仅支持使用单位h,如2h(最近2小时)"
            v-model="timeRange"
          )
          el-button.btn(
            type="primary"
            @click="doSaveSetting"
          ) 确认
      span.divided |
      router-link(
        :to="{ name: profileRoute }"
      )
        i.el-icon-user
        span {{user.account}}
      span.divided |
      a.logout(
        href="#"
        title="退出登录"
        @click.preventDefault="onLogout"
      )
        i.el-icon-switch-button
    div(
      v-else
    )
      router-link.login(
        :to="{ name: loginRoute }"
      )
        i.el-icon-user
        | 登录
      span.divided |
      router-link.register(
        :to="{ name: registerRoute }"
      )
        i.el-icon-circle-plus
        | 注册
  //- 应用图标
  h1(
    v-if="$props.shrinking"
  ): router-link(
    :to='{name: homeRoute}'
  )
    i.el-icon-cpu
    | Cyber Tect
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useUserStore } from "../store";
import {
  getHomeRouteName,
  getProfileRouteName,
  getLoginRouteName,
  getRegisterRouteName,
} from "../router";
import { getSetting, saveSetting } from "../services/setting";

export default defineComponent({
  name: "MainHeader",
  props: {
    shrinking: {
      type: Boolean,
      default: false,
    },
  },
  setup() {
    const userStore = useUserStore();
    return {
      user: userStore.state.info,
      logout: () => userStore.dispatch("logout"),
    };
  },
  data() {
    const setting = getSetting();
    return {
      querySize: setting.mainDetectorResultCount || 0,
      refreshInterval: setting.mainDetectorRefreshInterval || 0,
      onlyFailure: setting.mainDetectorOnlyFailure || false,
      timeRange: setting.mainDetectorTimeRange || "",
      profileRoute: getProfileRouteName(),
      homeRoute: getHomeRouteName(),
      loginRoute: getLoginRouteName(),
      registerRoute: getRegisterRouteName(),
    };
  },
  methods: {
    async onLogout() {
      try {
        await this.logout();
        this.$router.push({
          name: getHomeRouteName(),
        });
      } catch (err) {
        this.$error(err);
      }
    },
    async doSaveSetting() {
      const setting = getSetting();
      setting.mainDetectorResultCount = Number(this.querySize);
      setting.mainDetectorRefreshInterval = Number(this.refreshInterval);
      setting.mainDetectorOnlyFailure = Boolean(this.onlyFailure);
      setting.mainDetectorTimeRange = this.timeRange;
      try {
        await saveSetting(setting);
        this.$message.info("已成功更新，请刷新首页");
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
h1
  font-size: 18px
  margin-left: 10px
  width: 200px
  i
    margin-right: 5px
  a
    color: $dark
.header
  height: $mainHeaderHeight
  background-color: $white
  padding: 5px 0
  line-height: $mainHeaderHeight - 10
  color: $darkBlue
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08)
.userInfo
  float: right
  font-size: 13px
  margin-right: $mainMargin
  i
    margin-right: 3px
    font-weight: bold
.divided
  margin: 0 15px
.btn
  width: 100%
</style>
