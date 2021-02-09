<template lang="pug">
.detectorResultSummary.clearfix
  el-tooltip(
    v-for="item in items"
    :key="item.id"
  )
    template(
      #content
    )
      div(
        v-html="formatContent(item)"
      )
    .summary(
      @click="showTaskReuslts(item)"
      :style="{backgroundColor: item.color}"
    )
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { formatDate } from "../../helpers/util";

// 图示方块宽度
const boxWidth = 20;
class Color {
  r: number;
  g: number;
  b: number;
  constructor(r, g, b) {
    this.r = r;
    this.g = g;
    this.b = b;
  }
  toString() {
    return `rgba(${this.r}, ${this.g}, ${this.b})`;
  }
}

const defaultFailColor = new Color(193, 195, 199);
const defaultSuccessColor = new Color(3, 134, 134);
const defaultSuccessSlowColor = new Color(3, 185, 185);
const defaultSuccessSlowerColor = new Color(165, 208, 208);

const resultFail = 2;

export default defineComponent({
  name: "DetectorResultSummary",
  props: {
    route: {
      type: String,
      required: true,
    },
    results: {
      type: Array,
      default: [],
    },
  },
  data() {
    const items = [];
    this.$props.results.forEach((item: any) => {
      if (item.result === resultFail) {
        item.color = defaultFailColor.toString();
      } else {
        if (item.maxDuration >= 3000) {
          item.color = defaultSuccessSlowerColor.toString();
        } else if (item.maxDuration >= 1000) {
          item.color = defaultSuccessSlowColor.toString();
        } else {
          item.color = defaultSuccessColor.toString();
        }
      }
      items.push(item);
    });
    return {
      items,
    };
  },
  methods: {
    showTaskReuslts(item) {
      this.$router.push({
        name: this.$props.route,
        query: {
          task: `${item.task}`,
        },
      });
    },
    formatContent(item) {
      const resultDesc = item.result !== resultFail ? "成功" : "失败";
      let content = `任务：${item.name}<br />
      结果：${formatDate(item.updatedAt)} 检测${resultDesc}<br />
        耗时：${item.maxDuration || 0}ms
      `;
      if (item.result !== resultFail) {
        return content;
      }
      content += `<br />
      失败原因：${item.messages.join(",")}
      `;
      return content;
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../../common";

$summaryWidth = 20px
.detectorResultSummary
  margin: 15px
  margin-right: 0
  overflow: hidden
.summary
  width: $summaryWidth
  height: $summaryWidth
  float: left
  cursor: pointer
  border: 1px solid white
  border-radius: 3px
.tips
  font-size: 13px
  i
    margin-right: 3px
</style>
