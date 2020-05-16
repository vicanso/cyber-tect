<template lang="pug">
LoginRegister.register(
  type="register"
)
</template>
<script>
import LoginRegister from '@/components/LoginRegister.vue'

export default {
  name: 'Register',
  components: {
    LoginRegister
  }
}
</script>
<style lang="sass" scoped>
.register
  margin-top: 100px
</style>
