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

export default defineComponent({
  name: "MainHeader",
  setup() {
    const userStore = useUserStore();
    return {
      user: userStore.state.info,
      logout: () => userStore.dispatch("logout"),
    };
  },
  data() {
    return {
      profileRoute: getProfileRouteName(),
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
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
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
</style>
