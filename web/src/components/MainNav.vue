<template lang="pug">
  nav.mainNav
    el-menu
      el-submenu(
        v-for="nav in navs"
        :index="nav.name"
      )
        template(
          slot="title"
        )
          span {{nav.name}}
</template>

<script>
export default {
  name: 'MainNav',
  data () {
    return {
      navs: [
        {
          name: 'HTTP'
        },
        {
          name: 'TCP'
        },
        {
          name: 'DNS'
        },
        {
          name: 'PING'
        }
      ]
    }
  }
}
</script>
