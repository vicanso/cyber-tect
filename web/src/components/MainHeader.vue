<template lang="pug">
header.header
  .userInfo
    span(
      v-if="fetchingUserInfo"
    ) 加载用户信息...
    div(
      v-else-if="!userAccount"
    )
      router-link(
        :to="LOGIN"
      )
        el-link(type="primary") 登录
      span.divided |
      router-link(
        :to="REGISTER"
      )
        el-link(type="info") 注册
    .functions(
      v-else
    )
      router-link.userinfoUpdate(
        :to="USER_INFO"
        title="用户信息"
      )
        span {{userAccount}}
        i.el-icon-user
      span.divided |
      a.logout(
        href="#"
        title="退出登录"
        @click="logout"
      )
        i.el-icon-switch-button
  h3
    router-link(
      :to="HOME"
    )
      i.el-icon-cpu
      | Cyber Detect
</template>

<script>
import { mapState, mapActions } from "vuex";
import { HOME, LOGIN, REGISTER, USER_INFO } from "@/paths";

export default {
  name: "MainHeader",
  data() {
    return {
      HOME,
      LOGIN,
      REGISTER,
      USER_INFO,
    };
  },
  computed: mapState({
    fetchingUserInfo: (state) => state.user.fetching,
    userAccount: (state) => state.user.account,
  }),
  methods: {
    ...mapActions(["logout"]),
  },
};
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.header
  height: $mainHeaderHeight
  background-color: $dark
  line-height: $mainHeaderHeight
  color: $white
h3
  margin-left: 15px
  a
    font-size: 16px
    color: $white
    i
      font-weight: bold
      margin-right: 5px
.userInfo
  font-size: 13px
  float: right
  margin: 10px 30px
  line-height: $mainHeaderHeight - 20
.divided
  margin: 0 15px
  color: $lightGray
.userinfoUpdate
  color: $white
  i
    margin-left: 3px
.logout
  color: $white
</style>
