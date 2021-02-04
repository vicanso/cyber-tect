<template lang="pug">
.mainNav
  //- 切换侧边栏
  a.toggleNav(
    href="#"
    @click.prevent="toggleNav"
  )
    i(
      :class=`$props.shrinking ? "el-icon-s-unfold" : "el-icon-s-fold"`
    )
  //- 应用图标
  h1
    router-link(
      v-if="!$props.shrinking"
      :to='{name: homeRoute}'
    )
      i.el-icon-eleme
      | Cyber Tect 
  //- 菜单栏
  nav: el-menu.menu(
    :collapse="$props.shrinking"
    :default-active="active"
    background-color="#000c17"
    text-color="#fff"
    active-text-color="#fff"
  )
    el-submenu.submenu(
      v-for="(nav, i) in navs"
      :index="`${i}`"
      :key="`${i}`"
    )
      template(
        #title
      )
        i(
          :class="nav.icon"
        )
        span {{nav.name}}
      //- 子菜单栏
      el-menu-item.menuItem(
        v-for="(subItem, j) in nav.children"
        :index="`${i}-${j}`"
        :key="`${i}-${j}`"
        @click="goTo(subItem)"
      )
        span {{subItem.name}}
</template>

<script lang="ts">
import { defineComponent } from "vue";
import {
  getHomeRouteName,
  getLoginsRouteName,
  getUsersRouteName,
  getTrackersRouteName,
  getMockTimeRouteName,
  getBlockIPRouteName,
  getSignedKeyRouteName,
  getRouterMockRouteName,
  getRouterConcurrencyRouteName,
  getSessionInterceptorRouteName,
  getConfigurationRouteName,
  getOthersRouteName,
  getHTTPErrorsRouteName,
  getDetectorHTTPRouteName,
  getDetectorDNSRouteName,
  getDetectorTCPRouteName,
  getDetectorPingRouteName,
  getDetectorResultHTTPRouteName,
  getActionsRouteName,
} from "../router";
import { USER_ADMIN, USER_SU } from "../constants/user";
import { useUserStore } from "../store";
import { isAllowedUser } from "../helpers/util";

const navs = [
  {
    name: "检测配置",
    icon: "el-icon-monitor",
    roles: [],
    groups: [],
    children: [
      {
        name: "HTTP",
        route: getDetectorHTTPRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "DNS",
        route: getDetectorDNSRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "TCP",
        route: getDetectorTCPRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "Ping",
        route: getDetectorPingRouteName(),
        roles: [],
        groups: [],
      },
    ],
  },
  {
    name: "检测结果",
    icon: "el-icon-s-operation",
    roles: [],
    groups: [],
    children: [
      {
        name: "HTTP",
        route: getDetectorResultHTTPRouteName(),
        roles: [],
        groups: [],
      },
    ],
  },
  {
    name: "用户",
    icon: "el-icon-user",
    roles: [USER_ADMIN, USER_SU],
    groups: [],
    children: [
      {
        name: "用户列表",
        route: getUsersRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "登录记录",
        route: getLoginsRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "用户行为",
        route: getTrackersRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "出错记录",
        route: getHTTPErrorsRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "客户端行为记录",
        route: getActionsRouteName(),
        roules: [],
        groups: [],
      },
    ],
  },
  {
    name: "配置",
    icon: "el-icon-setting",
    roles: [USER_SU],
    groups: [],
    children: [
      {
        name: "所有配置",
        route: getConfigurationRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "MockTime配置",
        route: getMockTimeRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "黑名单IP",
        route: getBlockIPRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "SignedKey配置",
        route: getSignedKeyRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "路由Mock配置",
        route: getRouterMockRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "路由并发配置",
        route: getRouterConcurrencyRouteName(),
        roles: [],
        groups: [],
      },
      {
        name: "Session拦截配置",
        route: getSessionInterceptorRouteName(),
        roles: [],
        groups: [],
      },
    ],
  },
  {
    name: "其它",
    icon: "el-icon-setting",
    roles: [USER_SU],
    groups: [],
    children: [
      {
        name: "其它",
        route: getOthersRouteName(),
        roles: [],
        groups: [],
      },
    ],
  },
];

export default defineComponent({
  name: "MainNav",
  props: {
    shrinking: {
      type: Boolean,
      default: false,
    },
    onToggle: {
      type: Function,
      default: null,
    },
  },
  emits: ["toggle"],
  setup() {
    const userStore = useUserStore();
    return {
      user: userStore.state.info,
    };
  },
  data() {
    return {
      homeRoute: getHomeRouteName(),
      active: "",
    };
  },
  computed: {
    navs() {
      const { user } = this;
      if (!user || !user.account) {
        return [];
      }
      const { roles, groups } = user;
      const filterNavs = [];
      navs.forEach((item) => {
        // 如果该栏目有配置权限，而且用户无该权限
        if (item.roles && !isAllowedUser(item.roles, roles)) {
          return;
        }
        // 如果该栏目配置了允许分级，而该用户不属于该组
        if (item.groups && !isAllowedUser(item.groups, groups)) {
          return;
        }
        const clone = Object.assign({}, item);
        const children = item.children.map((subItem) =>
          Object.assign({}, subItem)
        );
        clone.children = children.filter((subItem) => {
          // 如果未配置色色与分组限制
          if (!subItem.roles && !subItem.groups) {
            return true;
          }
          if (subItem.roles && !isAllowedUser(subItem.roles, roles)) {
            return false;
          }
          if (subItem.groups && !isAllowedUser(subItem.groups, groups)) {
            return false;
          }
          return true;
        });
        filterNavs.push(clone);
      });
      return filterNavs;
    },
  },
  watch: {
    // 如果nav变化时，根据当前route定位
    navs() {
      this.updateActive(this.$route.name);
    },
    // 路由变化时设置对应的导航为活动状态
    $route(to) {
      this.updateActive(to.name);
    },
  },
  beforeMount() {
    this.updateActive(this.$route.name);
  },
  methods: {
    toggleNav() {
      this.$emit("toggle");
    },
    goTo({ route }) {
      if (!route || this.$route.name === route) {
        return;
      }
      this.$router.push({
        name: route,
      });
    },
    // 查询定位当前选中菜单
    updateActive(name: string) {
      const { navs } = this;
      let active = "";
      navs.forEach((nav, i) => {
        nav.children.forEach((item, j) => {
          if (item.route === name) {
            active = `${i}-${j}`;
          }
        });
      });
      this.active = active;
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
$mainNavColor = #000c17
.mainNav
  min-height: 100vh
  overflow-y: auto
  background-color: $mainNavColor
.toggleNav
  height: $mainHeaderHeight
  line-height: $mainHeaderHeight
  display: block
  float: right
  width: $mainNavShrinkingWidth
  text-align: center
h1
  height: $mainHeaderHeight
  line-height: $mainHeaderHeight
  color: $white
  padding-left: 20px
  font-size: 18px
  margin-right: $mainNavShrinkingWidth
  i
    font-weight: bold
    margin-right: 5px
nav
  border-top: 1px solid rgba($white, 0.3)
.menu
  border-right: 1px solid $mainNavColor
.menuItem
  color: rgba($white, 0.65)
  &.is-active
    background-color: $darkBlue !important
</style>
