<template lang="pug">
  nav.mainNav
    el-menu
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
  ADD_HTTP,
  LIST_HTTP,
  LIST_HTTP_RESULT,
  ADD_DNS,
  LIST_DNS,
  ADD_TCP,
  LIST_TCP,
  ADD_PING,
  LIST_PING
} from '@/paths'

export default {
  name: 'MainNav',
  data () {
    return {
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
            },
            {
              name: '添加HTTP检测',
              path: ADD_HTTP
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
              name: '添加DNS检测',
              path: ADD_DNS
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
              name: '添加TCP检测',
              path: ADD_TCP
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
              name: '添加PING检测',
              path: ADD_PING
            }
          ]
        }
      ]
    }
  }
}
</script>

<style lang="sass" scoped>
@import "@/common.sass"
.link
  display: block
  color: $darkGray
.menu.is-active
  background-color: $lightBlue
  .link
    color: $blue
</style>
