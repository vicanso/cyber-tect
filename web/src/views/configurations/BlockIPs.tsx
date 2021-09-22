import { defineComponent } from "vue";
import ExConfigEditorList from "../../components/ExConfigEditorList";
import { ConfigCategory } from "../../states/configs";

export default defineComponent({
  name: "Configs",
  render() {
    const extraFormItems = [
      {
        name: "IP地址：",
        key: "data",
        placeholder: "请输入IP地址或网段",
      },
    ];
    return (
      <ExConfigEditorList
        listTitle="黑名单IP配置"
        editorTitle="添加/更新黑名单配置"
        editorDescription="用于拦截访问IP"
        category={ConfigCategory.BlockIP}
        extraFormItems={extraFormItems}
      />
    );
  },
});
