<template lang="pug">
.tcp
  el-card(
    v-if="!editing"
  )
    template(
      #header
    ) TCP检测配置
    el-table(
      v-loading="tcps.processing"
      :data="tcps.items"
      row-key="id"
      stripe
      @sort-change="handleSortChange"
    )
      //- ID
      el-table-column(
        prop="id"
        key="id"
        label="ID"
        width="80"
      )
      el-table-column(
        prop="name"
        key="name"
        label="名称"
        width="150"
      )
      //- 检测地址
      el-table-column(
        prop="addrs"
        key="addrs"
        label="地址"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="item in scope.row.addrs"
        ) {{item}}
      //- 状态
      el-table-column(
        prop="status"
        key="status"
        label="状态"
        width="80"
      ): template(
        #default="scope"
      ): detector-status-desc(
        :status="scope.row.status"
      )
      //- 超时
      el-table-column(
        prop="timeout"
        key="timeout"
        label="超时"
        width="80"
      )
      //- 告警接收者列表
      el-table-column(
        prop="receivers"
        key="receivers"
        label="告警接收者"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="item in scope.row.receivers"
        ) {{item}}
      //- 描述
      el-table-column(
        prop="description"
        key="description"
        label="描述"
      )
      //- 创建者
      el-table-column(
        prop="owner"
        key="owner"
        label="创建者"
      )
      //- 更新时间
      el-table-column(
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
      ): template(
        #default="scope"
      )
        div(
          v-if="scope.row.owner === userInfo.account"
        ): ex-button(
          :onClick="generateModifyHandler(scope.row)"
          category="smallText"
        ) 编辑
        span(
          v-else
        ) --
    //- 分页
    el-pagination.pagination(
      layout="prev, pager, next, sizes"
      :current-page="currentPage"
      :page-size="query.limit"
      :page-sizes="pageSizes"
      :total="tcps.count"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    )
    el-button.btn(
      type="primary"
      @click="add" 
    ) 添加
  el-card(
    v-else
  )
    template(
      #header
    ) {{ modifiedItem.id ? "修改TCP配置" : "添加TCP配置" }}
    el-form(
      label-width="90px"
      ref="form"
    ): el-row(
      :gutter="15"
    )
      //- 配置名称
      el-col(
        :span="8"
      ): el-form-item(
        label="名称："
      ): el-input(
        placeholder="请输入TCP配置名称"
        v-model="modifiedItem.name"
        clearable
      )
      //- 检测地址
      el-col(
        :span="8"
      ): el-form-item(
        label="检测地址："
      ): el-input(
        placeholder="请输入检测地址，多个地址以,分隔"
        v-model="modifiedItem.addrs"
        clearable
      )
      //- 配置状态
      el-col(
        :span="8"
      ): detector-status-selector(
        :status="modifiedItem.status"
        @change="changeStatus"
      )
      //- 超时设置
      el-col(
        :span="8"
      ): el-form-item(
        label="超时："
      ): el-input(
        type="number"
        placeholder="请输入超时设置"
        v-model="modifiedItem.timeout"
        clearable
      ): template(
        #append
      ) 秒
      //- 接收者列表
      el-col(
        :span="8"
      ): detector-receiver-selector(
        :receivers="modifiedItem.receivers"
        @change="changeReceivers"
      )
      //- 描述
      el-col(
        :span="24"
      ): el-form-item(
        label="描述："
      ): el-input(
        type="textarea"
        v-model="modifiedItem.description"
        placeholder="请输入检测配置描述"
        :autosize="{ minRows: 2 }"
      )
      el-col(
        :span="12"
      ): ex-button(
        :onClick="submit"
      ) 提交
      el-col(
        :span="12"
      ): el-button.btn(
        @click="back"
      ) 返回
</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useDetectorStore, useUserStore } from "../../store";
import DetectorBase from "./Detector";
import TimeFormater from "../../components/TimeFormater.vue";
import DetectorStatusDesc from "../../components/detectors/StatusDesc.vue";
import DetectorStatusSelector from "../../components/detectors/StatusSelector.vue";
import DetectorReceiverSelector from "../../components/detectors/ReceiverSelector.vue";
import ExButton from "../../components/ExButton.vue";
import { PAGE_SIZES } from "../../constants/common";
import { diff } from "../../helpers/util";

export default defineComponent({
  name: "DetectorTCP",
  components: {
    TimeFormater,
    DetectorStatusDesc,
    ExButton,
    DetectorStatusSelector,
    DetectorReceiverSelector,
  },
  mixins: [DetectorBase],
  setup() {
    const detectorStore = useDetectorStore();
    const userStore = useUserStore();
    return {
      userInfo: userStore.state.info,
      tcps: detectorStore.state.tcps,
      listTCP: (params) => detectorStore.dispatch("listTCP", params),
      addTCP: (params) => detectorStore.dispatch("addTCP", params),
      updateTCPByID: (params) =>
        detectorStore.dispatch("updateTCPByID", params),
    };
  },
  data() {
    return {
      originalItem: null,
      editing: false,
      modifiedItem: {
        id: null,
        name: "",
        status: null,
        timeout: null,
        receivers: [],
        description: "",
        addrs: "",
      },
      pageSizes: PAGE_SIZES,
      query: {
        offset: 0,
        limit: PAGE_SIZES[0],
        order: "-updatedAt",
      },
    };
  },
  methods: {
    async fetch() {
      const { query, tcps } = this;
      if (tcps.processing) {
        return;
      }
      try {
        await this.listTCP(query);
      } catch (err) {
        this.$error(err);
      }
    },
    // submit 提交
    async submit(): Promise<boolean> {
      let isSuccess = false;
      const { modifiedItem, tcps, originalItem } = this;
      if (tcps.processing) {
        return isSuccess;
      }
      const item = this.convertToTCP(modifiedItem);
      try {
        // 添加
        if (!item.id) {
          await this.addTCP(item);
        } else {
          const updateData = diff(item, originalItem);
          if (updateData.modifiedCount === 0) {
            throw new Error("请先修改再提交");
          }
          await this.updateTCPByID({
            id: item.id,
            data: updateData.data,
          });
        }
        this.back();
        isSuccess = true;
      } catch (err) {
        this.$error(err);
      }
      return isSuccess;
    },
    convertToModified(item) {
      let timeout = null;
      if (item.timeout) {
        const arr = /\d+/.exec(item.timeout);
        timeout = Number(arr[0]);
      }
      return {
        id: item.id,
        name: item.name,
        status: item.status,
        timeout,
        receivers: item.receivers,
        description: item.description,
        addrs: (item.addrs || []).join(","),
      };
    },
    convertToTCP(item) {
      return {
        id: item.id,
        name: item.name,
        status: item.status,
        timeout: `${item.timeout}s`,
        receivers: item.receivers,
        description: item.description,
        addrs: item.addrs.split(","),
      };
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../../common";

.tcp
  margin: $mainMargin
ul
  list-style-position: inside
.btn
  margin-top: $mainMargin
  width: 100%
.pagination
  text-align: right
  margin-top: 15px
</style>
