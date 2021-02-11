<template lang="pug">
#app(
  :class="{ shrinking: shrinking }"
  v-loading="loadingSetting"
)
  //- 主头部
  main-header.header(
    :shrinking="shrinking"
    v-if="!loadingSetting"
  )
  //- 主导航
  main-nav.nav(
    :shrinking="shrinking"
    @toggle="toggleNav"
    v-if="!loadingSetting"
  )
  //- 内容区域
  .mainContent
    router-view(
      v-if="inited && !loadingSetting"
    )
    p.tac(
      v-else
    ) ...

</template>

<script lang="ts">
import { defineComponent } from "vue";
import MainHeader from "./components/MainHeader.vue";
import MainNav from "./components/MainNav.vue";

import { useUserStore } from "./store";
import { getLoginRouteName } from "./router";
import { loadSetting, getSetting, saveSetting } from "./services/setting";

export default defineComponent({
  name: "App",
  components: {
    MainHeader,
    MainNav,
  },
  setup() {
    const userStore = useUserStore();
    return {
      userInfo: userStore.state.info,
      fetchUserInfo: () => userStore.dispatch("fetch"),
      updateUserInfo: (params) => userStore.dispatch("update", params),
    };
  },
  data() {
    return {
      shrinking: false,
      loadingSetting: false,
      // 是否初始化完成
      inited: false,
    };
  },
  async beforeMount() {
    this.loadingSetting = true;
    try {
      await loadSetting();
      const setting = getSetting();
      this.shrinking = setting.mainNavShrinking;
    } catch (err) {
      this.$error(err);
    } finally {
      this.loadingSetting = false;
    }
  },
  mounted() {
    this.fetch();
  },
  methods: {
    toggleNav() {
      this.shrinking = !this.shrinking;
      const setting = getSetting();
      setting.mainNavShrinking = this.shrinking;
      saveSetting(setting);
    },
    async fetch() {
      const { userInfo, $router } = this;
      try {
        await this.fetchUserInfo();
        // 如果未登录则跳转至登录
        if (!userInfo.account) {
          $router.push({
            name: getLoginRouteName(),
          });
        } else {
          // 如果已登录，刷新cookie有效期（不关注刷新是否成功，因此不用await）
          this.updateUserInfo({});
        }
      } catch (err) {
        this.$error(err);
      } finally {
        this.inited = true;
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "./common";
.shrinking
  .header
    left: $mainNavShrinkingWidth
  .nav
    width: $mainNavShrinkingWidth
  .mainContent
    padding-left: $mainNavShrinkingWidth
.header
  position: fixed
  left: $mainNavWidth
  top: 0
  right: 0
  z-index: 9
.nav
  position: fixed
  width: $mainNavWidth
  top: 0
  bottom: 0
  left: 0
  overflow: hidden
  overflow-y: auto
.mainContent
  padding-left: $mainNavWidth
  padding-top: $mainHeaderHeight
</style>
