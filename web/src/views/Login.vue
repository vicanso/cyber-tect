<template lang="pug">
login-register(
  type="login"
)
</template>

<script lang="ts">
import { defineComponent } from "vue";

import LoginRegister from "../components/LoginRegister.vue";
export default defineComponent({
  name: "Login",
  components: {
    LoginRegister,
  },
});
</script>
