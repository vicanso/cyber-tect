<template lang="pug">
  header.header
    .userInfo
      span(
        v-if="fetchingUserInfo"
      ) 加载用户信息...
      router-link(
        to="/login"
        v-else-if="!userAccount"
      )
        el-link(type="primary") 请先登录
    h3 Cyber Detect

</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'MainHeader',
  computed: mapState({
    fetchingUserInfo: state => state.user.fetching,
    userAccount: state => state.user.account
  })
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.header
  height: $mainHeaderHeight
  background-color: $darkBlue
  line-height: $mainHeaderHeight
  color: $white
h3
  margin-left: 15px
.userInfo
  font-size: 13px
  float: right
  margin: 10px 15px
  line-height: $mainHeaderHeight - 20
</style>
