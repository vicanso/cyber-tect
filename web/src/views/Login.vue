<template lang="pug">
LoginRegister.login(
  type="login"
)
</template>
<script>
import LoginRegister from "@/components/LoginRegister.vue";

export default {
  name: "Login",
  components: {
    LoginRegister,
  },
};
</script>
<style lang="sass" scoped>
.login
  margin-top: 100px
</style>
