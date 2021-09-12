import { useMessage } from "naive-ui";
import { defineComponent, onMounted } from "vue";
import ExConfigEditorList from "../components/ExConfigEditorList";
import { FormItemTypes, FormItem } from "../components/ExForm";
import ExLoading from "../components/ExLoading";
import { showError } from "../helpers/util";
import useCommonState, { commonListRouter } from "../states/common";
import { ConfigCategory } from "../states/configs";

export default defineComponent({
  name: "SessionInterceptors",
  setup() {
    const message = useMessage();
    const { routers } = useCommonState();

    onMounted(async () => {
      try {
        await commonListRouter();
      } catch (err) {
        showError(message, err);
      }
    });

    return {
      routers,
    };
  },
  render() {
    const { routers } = this;
    if (routers.processing) {
      return <ExLoading />;
    }
    const routes: string[] = [];
    routers.items.forEach((item) => {
      if (!routes.includes(item.route)) {
        routes.push(item.route);
      }
    });
    routes.sort();
    const extraFormItems: FormItem[] = [
      {
        type: FormItemTypes.Blank,
        name: "",
        key: "",
      },
      {
        name: "提示信息：",
        key: "data.message",
        placeholder: "请输入出错提示信息",
      },
      {
        name: "允许账号：",
        key: "data.allowAccount",
        placeholder: "请输入允许账号，多个账号以,分割",
      },
      {
        name: "允许路由：",
        key: "data.allowRoutes",
        type: FormItemTypes.MultiSelect,
        placeholder: "请输入允许路由，可以多选",
        options: routes.map((item) => {
          return {
            label: item,
            value: item,
          };
        }),
      },
    ];
    return (
      <ExConfigEditorList
        listTitle="Session拦截配置"
        editorTitle="添加/更新Session拦截配置"
        editorDescription="设置Session拦截的相关配置"
        category={ConfigCategory.SessionInterceptor}
        extraFormItems={extraFormItems}
      />
    );
  },
});
