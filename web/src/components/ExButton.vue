<template lang="pug">
template(
  v-if=`$props.category === "primary"`
)
  el-button.btn(
    type="primary"
    :icon="$props.icon"
    @click="handleClick"
    :class="{ isProcessing: processing}"
  )
    slot
    i.el-icon-loading.loading(
      v-if="processing"
    )
template(
  v-else-if=`$props.category === "smallText"`
)
  el-button(
    type="text"
    size="small"
    @click="handleClick"
    :class="{ isProcessing: processing}"
  )
    slot
    span(
      v-if="processing"
    ) ...
</template>
<script lang="ts">
// 此button的扩展可记录用户行为，防止重复点击等
import { defineComponent } from "vue";

import { getCurrentLocation } from "../router";
import { addUserAction, SUCCESS, FAIL, CLICK } from "../services/action";

export default defineComponent({
  name: "ExButton",
  props: {
    onClick: {
      type: Function,
      required: true,
    },
    category: {
      type: String,
      default: "primary",
    },
    icon: {
      type: String,
      default: "",
    },
    extra: {
      type: Object,
      default: function () {
        return {};
      },
    },
  },
  data() {
    return {
      // 是否处理中，避免重复点击
      processing: false,
    };
  },
  methods: {
    async handleClick() {
      const { processing, $props } = this;
      if (processing) {
        return;
      }
      this.processing = true;
      const currentLocation = getCurrentLocation();
      const data = Object.assign(
        {
          route: currentLocation.name,
          path: currentLocation.path,
          time: Math.floor(Date.now() / 1000),
        },
        $props.extra
      );
      // 如果未设置分类，则设置为click
      if (!data.category) {
        data.category = CLICK;
      }
      // 由于在onClick会捕获异常处理，因此在此处无法判断是否成功
      data.result = SUCCESS;
      try {
        const isSuccess = await $props.onClick();
        // 如果onclick返回fail，则表示失败，其它均表示成功
        if (isSuccess === false) {
          data.result = FAIL;
        }
      } finally {
        this.processing = false;
      }
      addUserAction(data);
    },
  },
});
</script>
<style lang="stylus" scoped>
.isProcessing
  opacity: 0.5
.btn
  width: 100%
.loading
  margin-left: 10px
  font-weight: 900
</style>
