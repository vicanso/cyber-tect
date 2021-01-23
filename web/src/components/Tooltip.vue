<template lang="pug">
el-tooltip(
  v-if="$props.content"
  :content="$props.content"
)
  i(
    v-if="$props.icon"
    :class="$props.icon"
  )
  span(
    v-else
  ) {{$props.content.substring($props.content.length - $props.viewSize)}}
span(
  v-else
) --
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "BaseTooltip",
  props: {
    icon: {
      type: String,
      default: "",
    },
    content: {
      type: String,
      default: "",
    },
    viewSize: {
      type: Number,
      default: 5,
    },
  },
});
</script>
