import { defineComponent } from "vue";

import FilterTable from "../../mixins/FilterTable";

export default defineComponent({
  name: "DetectorBase",
  mixins: [FilterTable],
  methods: {
    // 状态变化
    changeStatus(value) {
      this.modifiedItem.status = value;
    },
    // 接收人变化
    changeReceivers(values) {
      this.modifiedItem.receivers = values;
    },
    // generateModifyHandler 生成修改的处理函数
    generateModifyHandler(item) {
      return () => this.modify(item);
    },
    // modify 修改
    modify(item) {
      this.editing = true;
      this.originalItem = item;
      Object.assign(this.modifiedItem, this.convertToModified(item));
    },
    // add 添加
    add() {
      this.editing = true;
      this.originalItem = null;
      Object.assign(this.modifiedItem, this.convertToModified({}));
    },
    back() {
      this.editing = false;
    },
  },
});
