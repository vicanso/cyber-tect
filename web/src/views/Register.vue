<template lang="pug">
login-register(
  type="register"
)
</template>
<script lang="ts">
import { defineComponent } from "vue";

import LoginRegister from "../components/LoginRegister.vue";
export default defineComponent({
  name: "Register",
  components: {
    LoginRegister,
  },
});
</script>
