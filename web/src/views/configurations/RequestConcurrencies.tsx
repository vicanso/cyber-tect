import { useMessage } from "naive-ui";
import { defineComponent, onMounted } from "vue";
import ExConfigEditorList from "../../components/ExConfigEditorList";
import { FormItemTypes, FormItem } from "../../components/ExFormInterface";
import ExLoading from "../../components/ExLoading";
import { showError } from "../../helpers/util";
import useCommonState, { commonListRequestInstance } from "../../states/common";
import { ConfigCategory } from "../../states/configs";

export default defineComponent({
  name: "RequestCOncurrencies",
  setup() {
    const { requestInstances } = useCommonState();
    const message = useMessage();

    onMounted(async () => {
      try {
        await commonListRequestInstance();
      } catch (err) {
        showError(message, err);
      }
    });

    return {
      requestInstances,
    };
  },

  render() {
    const { requestInstances } = this;
    if (requestInstances.processing) {
      return <ExLoading />;
    }
    const extraFormItems: FormItem[] = [
      {
        type: FormItemTypes.Blank,
        name: "",
        key: "",
      },
      {
        name: "实例：",
        key: "data.name",
        type: FormItemTypes.Select,
        placeholder: "请选择限制并发数的实例",
        options: requestInstances.items.map((item) => {
          return {
            label: item.name,
            value: item.name,
          };
        }),
      },
      {
        name: "并发数：",
        key: "data.max",
        type: FormItemTypes.InputNumber,
        placeholder: "请输入并发限制",
      },
    ];

    return (
      <ExConfigEditorList
        listTitle="HTTP请求实例并发配置"
        editorTitle="添加/更新HTTP请求实例并发限制"
        editorDescription="设置各HTTP请求实例的并发请求数"
        category={ConfigCategory.RequestConcurrency}
        extraFormItems={extraFormItems}
      />
    );
  },
});
