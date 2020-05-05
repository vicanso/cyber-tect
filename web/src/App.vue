<template lang="pug">
  #app
    MainHeader
    MainNav.nav
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
      'userFetch'
    ])
  },
  async mounted () {
    try {
      await this.userFetch()
    } catch (err) {
      this.$message.error(err.message)
    }
  }
}
</script>
<style lang="sass" scoped>
@import "@/common.sass"
.nav
  position: fixed
  width:  200px
  top: $mainHeaderHeight
  bottom: 0
  left: 0
  background-color: $white
</style>
