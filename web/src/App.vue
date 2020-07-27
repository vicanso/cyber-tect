<template lang="pug">
  #app
    MainHeader.header
    MainNav.nav
    .mainContent
      router-view
</template>
<script>
import { mapActions, mapState } from "vuex";

import MainHeader from "@/components/MainHeader";
import MainNav from "@/components/MainNav";
export default {
  components: {
    MainHeader,
    MainNav,
  },
  name: "App",
  computed: mapState({
    userAccount: (state) => state.user.account,
  }),
  methods: {
    ...mapActions(["fetchUser", "updateUser"]),
    refreshSessionTTL() {
      if (!this.userAccount) {
        return;
      }
      this.updateUser({});
    },
  },
  async mounted() {
    setInterval(() => {
      this.refreshSessionTTL();
    }, 5 * 60 * 1000);
    try {
      await this.fetchUser();
    } catch (err) {
      this.$message.error(err.message);
    }
  },
};
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
