<template lang="pug">
span {{value}}
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { getDetectorStatusList } from "../../helpers/util";

export default defineComponent({
  name: "DetectorStatusDesc",
  props: {
    status: {
      type: Number,
      default: 0,
    },
  },
  data() {
    let value = "--";
    const statusList = getDetectorStatusList();
    statusList.forEach((element) => {
      if (element.value === this.$props.status) {
        value = element.name;
      }
    });
    return {
      value,
    };
  },
});
</script>
