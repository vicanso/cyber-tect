<template lang="pug">
span(
  v-if="$props.status === 1"
) 成功
span(
  v-else-if="$props.status === 2"
) 失败
span(
  v-else
) --
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "DetectorResultStatus",
  props: {
    status: {
      type: Number,
      default: 0,
    },
  },
});
</script>
