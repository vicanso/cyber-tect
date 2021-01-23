<template lang="pug">
//- 配置的数据展示
.configData
  pre(
    v-if="$props.expanded && json"
  ) {{ data }}
  el-tooltip(
    v-else-if="json"
    placement="bottom"
    :content="data"
  )
    template(
      #content
    )
      pre {{ data }}
    i.el-icon-info
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
  name: "ConfigData",
  props: {
    expanded: {
      type: Boolean,
      default: false,
    },
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
