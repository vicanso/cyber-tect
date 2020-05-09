<template lang="pug">
  #app
    MainHeader.header
    MainNav.nav
    .mainContent
      router-view
</template>
<script>
import { mapActions } from 'vuex'

import MainHeader from '@/components/MainHeader'
import MainNav from '@/components/MainNav'
export default {
  components: {
    MainHeader,
    MainNav
  },
  name: 'App',
  methods: {
    ...mapActions([
      'fetchUser'
    ])
  },
  async mounted () {
    try {
      await this.fetchUser()
    } catch (err) {
      this.$message.error(err.message)
    }
  }
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.header
  position: fixed
  left: 0
  top: 0
  right: 0
  z-index: 9
.nav
  position: fixed
  width: $mainNavWidth
  top: $mainHeaderHeight
  bottom: 0
  left: 0
  background-color: $white
.mainContent
  padding-left: $mainNavWidth
  padding-top: $mainHeaderHeight
</style>
