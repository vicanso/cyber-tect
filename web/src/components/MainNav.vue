<template lang="pug">
nav.mainNav
  el-menu(
    :default-active="active"
  )
    el-submenu(
      v-for="(nav, i) in navs"
      :index="`${i}`"
      :key="`${i}`"
    )
      template(
        slot="title"
      )
        span {{nav.name}}
      el-menu-item.menu(
        v-for="(subItem, j) in nav.children"
        :index="`${i}-${j}`"
        :key="`${i}-${j}`"
      )
        router-link.link(
          :to="subItem.path"
        )
          span {{subItem.name}}
</template>

<script>
import {
  LIST_HTTP,
  LIST_HTTP_RESULT,
  LIST_DNS,
  LIST_DNS_RESULT,
  LIST_TCP,
  LIST_TCP_RESULT,
  LIST_PING,
  LIST_PING_RESULT
} from '@/paths'

export default {
  name: 'MainNav',
  data () {
    return {
      active: '',
      navs: [
        {
          name: 'HTTP',
          children: [
            {
              name: 'HTTP检测配置',
              path: LIST_HTTP
            },
            {
              name: 'HTTP检测结果',
              path: LIST_HTTP_RESULT
            }
          ]
        },
        {
          name: 'DNS',
          children: [
            {
              name: 'DNS检测配置',
              path: LIST_DNS
            },
            {
              name: 'DNS检测结果',
              path: LIST_DNS_RESULT
            }
          ]
        },
        {
          name: 'TCP',
          children: [
            {
              name: 'TCP检测配置',
              path: LIST_TCP
            },
            {
              name: 'TCP检测结果',
              path: LIST_TCP_RESULT
            }
          ]
        },
        {
          name: 'PING',
          children: [
            {
              name: 'PING检测配置',
              path: LIST_PING
            },
            {
              name: 'Ping检测结果',
              path: LIST_PING_RESULT
            }
          ]
        }
      ]
    }
  },
  watch: {
    '$route' (to, from) {
      const {
        navs
      } = this
      let active = ''
      navs.forEach((nav, i) => {
        nav.children.forEach((item, j) => {
          if (item.path === to.path) {
            active = `${i}-${j}`
          }
        })
      })
      this.active = active
    }
  },
  beforeRouteUpdate (to, from, next) {
    next()
  }
}
</script>

<style lang="sass" scoped>
@import "@/common.sass"
.mainNav
  overflow-y: auto
.link
  display: block
  color: $darkGray
.menu.is-active
  background-color: $lightBlue
  .link
    color: $blue
</style>
