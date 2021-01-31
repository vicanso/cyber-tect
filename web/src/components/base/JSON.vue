<template lang="pug">
.jsonView
  pre(
    v-if="$props.expanded && json"
  ) {{ data }}
  el-tooltip(
    v-if="json"
    placement="bottom"
    :content="data"
  ) 
    template(
      #content
    )
      pre {{ data }}
    i(
      v-if="$props.icon"
      :class="$props.icon"
    )
    pre(
      v-else
    ) {{ data }}
  span(
    v-else
  ) {{ data }}
</template>

<script lang="ts">
import { defineComponent } from "vue";

const recursionFormat = (data) => {
  const json = JSON.parse(data);
  Object.keys(json).forEach((key) => {
    const value = json[key];
    if (value && (value[0] === "{" || value[0] === "[")) {
      json[key] = recursionFormat(value);
    }
  });
  return json;
};
const formatJSON = (data) => {
  if (data && (data[0] === "{" || data[0] === "[")) {
    return {
      json: true,
      data: JSON.stringify(recursionFormat(data), null, 2),
    };
  }
  return {
    json: false,
    data: data || "--",
  };
};

export default defineComponent({
  name: "BaseJson",
  props: {
    expanded: {
      type: Boolean,
      default: false,
    },
    content: {
      type: String,
      default: "",
    },
    icon: {
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
