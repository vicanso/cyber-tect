<template lang="pug">
.detectorResultSummary.clearfix(
  v-loading="processing"
)
  template(
    v-if="detectorResults && detectorResults.length !== 0"
  )
    el-tooltip(
      v-for="item in detectorResults"
      :key="item.id"
    )
      div(
        slot="content"
        v-html="formatContent(item)"
      )
      .summary(
        @click="showTaskReuslts(item)"
        :style="{backgroundColor: item.color}"
      )
  .tips(
    v-else
  )
    i.el-icon-warning-outline
    | 请先设置自定义的检测配置
</template>
<script>
import { mapState, mapActions } from "vuex";

// 图示方块宽度
const boxWidth = 20;

class Color {
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

export default {
  name: "DetectorResultSummary",
  props: {
    category: {
      type: String,
      required: true,
    },
    route: {
      type: String,
      required: true,
    },
    mime: Boolean,
  },
  data() {
    return {
      filterTasks: "",
    };
  },
  computed: mapState({
    userAccount: (state) => state.user.account,
    resultSucess: (state) => state.detector.resultSucess,
    resultFail: (state) => state.detector.resultFail,
    processing: function (state) {
      const { category } = this.$props;
      const { detector } = state;
      return (
        detector.mime[category].processing ||
        detector[`${category}ListResult`].processing
      );
    },
    detectors: function (state) {
      return state.detector.mime[this.$props.category].detectors;
    },
    detectorResults: function (state) {
      const { results } = state.detector[`${this.$props.category}ListResult`];
      const { resultFail } = this;
      if (results) {
        results.forEach((item) => {
          if (item.result === resultFail) {
            item.color = defaultFailColor.toString();
          } else {
            if (item.duration >= 3000) {
              item.color = defaultSuccessSlowerColor.toString();
            } else if (item.duration >= 1000) {
              item.color = defaultSuccessSlowColor.toString();
            } else {
              item.color = defaultSuccessColor.toString();
            }
          }
        });
      }
      return results;
    },
  }),
  methods: {
    ...mapActions([
      "resetDetectorResults",
      "listMimeDetector",
      "listDetectorResult",
    ]),
    showTaskReuslts(item) {
      this.$router.push({
        name: this.route,
        query: {
          task: `${item.task}`,
        },
      });
    },
    formatContent(item) {
      let content = `任务：${item.task}<br />
      结果：${item.updatedAtDesc} 检测${item.resultDesc}<br />
        耗时：${item.durationDesc}
      `;
      if (item.result === this.resultSucess) {
        return content;
      }
      content += `<br />
      失败原因：${item.message}
      `;
      return content;
    },
    async fetchMimeDetector() {
      const { category } = this.$props;
      const params = {
        // 偷懒，直接取前100个
        limit: 100,
        offset: 0,
        owner: this.userAccount,
        order: "-id",
        fields: "id",
        status: 1,
      };
      try {
        await this.listMimeDetector({
          category,
          params,
        });
        const arr = this.detectors.map((item) => item.id);
        if (arr.length === 0) {
          return;
        }
        this.filterTasks = arr.join(",");
        await this.fetchDetectorSummary();
      } catch (err) {
        this.$message.error(err.error);
      }
    },
    async fetchDetectorSummary() {
      const eachColumnBoxCount = Math.floor(this.$el.clientWidth / boxWidth);
      const { category } = this.$props;
      const params = {
        fields: "result,id,message,task,duration,updatedAt",
        limit: eachColumnBoxCount * 8,
        offset: 0,
        order: "-id",
        // 不获取总数
        count: "0",
      };
      if (this.filterTasks) {
        params.tasks = this.filterTasks;
      }
      try {
        await this.listDetectorResult({
          category,
          params,
        });
      } catch (err) {
        this.$message.error(err.message);
      }
    },
  },
  mounted() {
    const { mime, category } = this.$props;
    this.resetDetectorResults({
      category,
    });
    if (mime) {
      this.fetchMimeDetector();
    } else {
      this.fetchDetectorSummary();
    }
  },
  beforeDestroy() {
    const { category } = this.$props;
    this.resetDetectorResults({
      category,
    });
  },
};
</script>
<style lang="sass" scoped>
@import '@/common'
$summaryWidth: 20px
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
.tips
  font-size: 13px
  i
    margin-right: 3px
</style>
