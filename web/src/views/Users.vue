<template lang="pug">
.users(
  v-loading="!inited"
)
  el-card(
    v-if="!editMode"
  )
    template(
      #header
    )
      i.el-icon-user-solid
      span 用户列表
    //- 筛选条件
    base-filter(
      :fields="filterFields"
      v-if="inited"
      @filter="filter"
    )
    div(
      v-loading="users.processing"
    ): el-table(
      :data="users.items"
      row-key="id"
      stripe
      @sort-change="handleSortChange"
    )
      //- 用户ID
      el-table-column(
        prop="id"
        key="id"
        label="ID"
        width="80"
        sortable
      )
      //- 用户账号
      el-table-column(
        prop="account"
        key="account"
        label="账户"
        width="120"
      )
      //- 用户状态
      el-table-column(
        label="状态"
        width="100"
      ): template(
        #default="scope"
      ) {{ getStatusDesc(scope.row.status) }}
      //- 用户角色
      el-table-column(
        label="角色"
      ): template(
        #default="scope"
      ): ul
        li(
          v-for="role in scope.row.roles"
          :key="role"
        ) {{ role }}
      //- 更新时间
      el-table-column(
        prop="updatedAt"
        key="updatedAt"
        label="更新于"
        width="160"
        sortable
      ): template(
        #default="scope"
      ): time-formater(
        :time="scope.row.updatedAt"
      )
      //- 操作
      el-table-column(
        label="操作"
        width="120"
      ): template(
        #default="scope"
      ): el-button.op(
        type="text"
        size="small"
        @click="modify(scope.row)"
      ) 编辑
    //- 分页
    el-pagination.pagination(
      layout="prev, pager, next, sizes"
      :current-page="currentPage"
      :page-size="query.limit"
      :page-sizes="pageSizes"
      :total="users.count"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    )
  //- 用户编辑
  user(
    v-else
  )


</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useUserStore, useCommonStore } from "../store";

import BaseFilter from "../components/base/Filter.vue";
import BaseTooltip from "../components/Tooltip.vue";
import TimeFormater from "../components/TimeFormater.vue";
import User from "../components/User.vue";
import { PAGE_SIZES } from "../constants/common";
import FilterTable from "../mixins/FilterTable";

const roleSelectList = [];
const statusSelectList = [];
const filterFields = [
  {
    label: "用户角色：",
    key: "role",
    type: "select",
    options: roleSelectList,
    span: 6,
  },
  {
    label: "用户状态：",
    key: "status",
    type: "select",
    options: statusSelectList,
    span: 6,
  },
  {
    label: "关键字：",
    key: "keyword",
    placeholder: "请输入关键字",
    clearable: true,
    span: 6,
  },
  {
    label: "",
    type: "filter",
    span: 6,
    labelWidth: "0px",
  },
];

export default defineComponent({
  name: "Users",
  components: {
    BaseFilter,
    BaseTooltip,
    TimeFormater,
    User,
  },
  mixins: [FilterTable],
  setup() {
    const userStore = useUserStore();
    const commonStore = useCommonStore();
    return {
      users: userStore.state.users,
      list: (params) => userStore.dispatch("list", params),
      listRole: () => userStore.dispatch("listRole"),
      listStatus: () => commonStore.dispatch("listStatus"),
      userRoles: userStore.state.roles,
      statuses: commonStore.state.statuses,
      getStatusDesc: (status: number): string => {
        let desc = "";
        commonStore.state.statuses.items.forEach((item) => {
          if (item.value === status) {
            desc = item.name;
          }
        });
        return desc;
      },
    };
  },
  data() {
    return {
      inited: false,
      filterFields: null,
      pageSizes: PAGE_SIZES,
      query: {
        offset: 0,
        limit: PAGE_SIZES[0],
        order: "-updatedAt",
      },
    };
  },
  async beforeMount() {
    try {
      await this.listRole();
      await this.listStatus();

      // 重置
      roleSelectList.length = 0;
      roleSelectList.push({
        name: "所有",
        value: "",
      });
      roleSelectList.push(...this.userRoles.items);

      // 重置
      statusSelectList.length = 0;
      statusSelectList.push({
        name: "所有",
        value: "",
      });
      statusSelectList.push(...this.statuses.items);

      this.filterFields = filterFields;
    } catch (err) {
      this.$error(err);
    } finally {
      this.inited = true;
    }
  },
  methods: {
    async fetch() {
      const { users, query } = this;
      if (users.processing) {
        return;
      }
      try {
        await this.list(query);
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
.users
  margin: $mainMargin
  i
    margin-right: 5px
  ul
    li
      list-style: inside
.selector, .submit
  width: 100%
.pagination
  text-align: right
  margin-top: 15px
</style>
