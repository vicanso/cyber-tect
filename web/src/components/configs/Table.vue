<template lang="pug">
el-card.configurationList
  template(
    #header
    v-if="!$props.hiddenHeader"
  )
    i.el-icon-s-tools
    span {{ $props.name || "系统配置" }}
    span.filters
      el-checkbox(
        title="仅展示有效的"
        v-model="available"
      ) 仅展示有效的
      el-checkbox(
        title="展开所有配置"
        v-model="expanded"
      ) 展开所有配置
  el-table(
    v-loading="configs.processing"
    :data="filterConfigs"
    row-key="id"
    stripe
    :default-sort="{ prop: 'updatedAt', order: 'descending'}"
  )
    el-table-column(
      prop="id"
      key="id"
      label="ID"
      width="80"
    )
    //- 名称
    el-table-column(
      prop="name"
      key="name"
      label="名称"
      width="150"
    )
    //- 分类
    el-table-column(
      prop="category"
      key="category"
      label="分类"
      width="150"
    )
    //- 状态
    el-table-column(
      sortable
      prop="status"
      key="status"
    ): template(
      #default="scope"
    ) {{ getStatusDesc(scope.row.status) }}
    //- 开始时间
    el-table-column(
      sortable
      prop="startedAt"
      key="startedAt"
      label="开始时间"
      width="180"
    ): template(
      #default="scope"
    ): time-formater(
      :time="scope.row.startedAt"
    )
    //- 结束时间
    el-table-column(
      sortable
      prop="endedAt"
      key="endedAt"
      label="结束时间"
      width="180"
    ): template(
      #default="scope"
    ): time-formater(
      :time="scope.row.endedAt"
    )
    //- 配置数据
    el-table-column(
      prop="data"
      key="data"
      label="配置数据"
      :width="configWidth"
    ): template(
      #default="scope"
    ): config-data(
      :content="scope.row.data"
    )
    //- 创建者
    el-table-column(
      prop="owner"
      key="owner"
      label="创建者"
      width="150"
    )
    //- 更新时间
    el-table-column(
      sortable
      prop="updatedAt"
      key="updatedAt"
      label="更新时间"
      width="180"
    ): template(
      #default="scope"
    ): time-formater(
      :time="scope.row.updatedAt"
    )
    //- 操作
    el-table-column(
      fixed="right"
      label="操作"
      v-if="!$props.hiddenOp"
    ): template(
      #default="scope"
    )
      div(
        v-if="scope.row.owner === userInfo.account"
      )
        ex-button(
          category="smallText"
          :onClick="generateModifyHandler(scope.row)"
        ) 编辑
      span(
        v-else
      ) --
</template>

<script lang="ts">
import { defineComponent } from "vue";

import TimeFormater from "../TimeFormater.vue";
import ConfigData from "./ConfigData.vue";
import ExButton from "../ExButton.vue";
import { useConfigStore, useUserStore } from "../../store";
import { CONFIG_ENABLED, CONFIG_EDIT_MODE } from "../../constants/common";

export default defineComponent({
  name: "ConfigurationList",
  components: {
    TimeFormater,
    ConfigData,
    ExButton,
  },
  props: {
    name: {
      type: String,
      default: "",
    },
    // 是否隐藏header
    hiddenHeader: {
      type: Boolean,
      default: false,
    },
    // 是否隐藏操作栏
    hiddenOp: {
      type: Boolean,
      default: false,
    },
    category: {
      type: String,
      required: true,
    },
  },
  setup() {
    const configStore = useConfigStore();
    const userStore = useUserStore();
    return {
      userInfo: userStore.state.info,
      configs: configStore.state.configs,
      list: (params) => configStore.dispatch("list", params),
      getStatusDesc: (status) => {
        let desc = "";
        configStore.state.statuses.forEach((item) => {
          if (item.value === status) {
            desc = item.label;
          }
        });
        return desc;
      },
    };
  },
  data() {
    const data = {
      expanded: false,
      available: false,
      query: {
        category: this.$props.category,
      },
    };
    if (data.query.category === "*") {
      data.query.category = "";
    }
    return data;
  },
  computed: {
    configWidth() {
      if (this.expanded) {
        return 200;
      }
      return 80;
    },
    filterConfigs() {
      const { configs, available } = this;
      const arr = (configs.items || []).filter((item) => {
        // 如果非选择仅展示有效的
        if (!available) {
          return true;
        }
        if (item.status !== CONFIG_ENABLED) {
          return false;
        }
        const now = Date.now();
        const beginDate = new Date(item.startedAt).getTime();
        const endDate = new Date(item.endedAt).getTime();
        // 如果未到开始时间或者已结束
        if (beginDate > now || endDate < now) {
          return false;
        }
        return true;
      });
      return arr;
    },
  },
  beforeMount() {
    this.fetch();
  },
  methods: {
    // 拉取数据
    async fetch() {
      const { query } = this;
      try {
        await this.list(query);
      } catch (err) {
        this.$error(err);
      }
    },
    // generateModifyHandler 生成修改的处理函数
    generateModifyHandler(item) {
      return () => this.modify(item);
    },
    // 修改
    modify(item) {
      this.$router.push({
        query: {
          mode: CONFIG_EDIT_MODE,
          id: item.id,
        },
      });
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../../common";
.configurationList
  margin: $mainMargin
  i
    margin-right: 3px
  .op
    margin: 0 10px
  .filters
    margin-left: 20px
</style>
