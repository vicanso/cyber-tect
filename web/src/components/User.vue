<template lang="pug">
.user: base-editor(
  v-if="!processing && fields"
  title="更新用户信息"
  icon="el-icon-user"
  :id="id"
  :findByID="findByID"
  :updateByID="updateByID"
  :fields="fields"
)

</template>

<script lang="ts">
import { defineComponent } from "vue";

import { useUserStore, useCommonStore } from "../store";
import BaseEditor from "./base/Editor.vue";

const roleSelectList = [];
const statusSelectList = [];
const fields = [
  {
    label: "账号：",
    key: "account",
    disabled: true,
  },
  {
    label: "用户角色：",
    key: "roles",
    type: "select",
    placeholder: "请选择用户角色",
    labelWidth: "100px",
    multiple: true,
    options: roleSelectList,
    rules: [
      {
        required: true,
        message: "用户角色不能为空",
      },
    ],
  },
  // {
  //   label: "用户组：",
  //   key: "groups",
  //   type: "select",
  //   placeholder: "请选择用户分组",
  //   multiple: true,
  //   options: userGroups
  //   // rules: [
  //   //   {
  //   //     required: true,
  //   //     message: "用户分组不能为空"
  //   //   }
  //   // ]
  // },
  {
    label: "用户状态：",
    key: "status",
    type: "select",
    placeholder: "请选择用户状态",
    labelWidth: "100px",
    options: statusSelectList,
    rules: [
      {
        required: true,
        message: "用户状态不能为空",
      },
    ],
  },
];

export default defineComponent({
  name: "User",
  components: {
    BaseEditor,
  },
  setup() {
    const userStore = useUserStore();
    const commonStore = useCommonStore();
    return {
      findByID: (id) =>
        userStore.dispatch("findByID", {
          id,
        }),
      updateByID: (params) => userStore.dispatch("updateByID", params),
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
      fields: null,
      processing: false,
      id: 0,
    };
  },
  async beforeMount() {
    this.processing = true;
    const { id } = this.$route.query;
    if (id) {
      this.id = Number(id);
    }
    try {
      await this.listRole();
      await this.listStatus();

      // 重置
      roleSelectList.length = 0;
      roleSelectList.push(...this.userRoles.items);

      // 重置
      statusSelectList.length = 0;
      statusSelectList.push(...this.statuses.items);

      this.fields = fields;
    } catch (err) {
      this.$error(err);
    } finally {
      this.processing = false;
    }
  },
});
</script>
