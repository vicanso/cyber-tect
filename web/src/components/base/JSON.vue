<template lang="pug">
.jsonView
  pre(
    v-if="json"
  ) {{ data }}
  span(
    v-else
  ) {{ data }}
</template>

<script lang="ts">
import { defineComponent } from "vue";

const formatJSON = (data) => {
  if (data && (data[0] === "{" || data[0] === "[")) {
    return {
      json: true,
      data: JSON.stringify(JSON.parse(data), null, 2),
    };
  }
  return {
    json: false,
    data: data,
  };
};

export default defineComponent({
  name: "BaseJson",
  props: {
    content: {
      type: String,
      default: "",
    },
  },
  data() {
    return formatJSON(this.$props.content);
  },
});
</script>
<style lang="stylus" scoped>
.jsonView
  overflow-x: scroll
</style>
