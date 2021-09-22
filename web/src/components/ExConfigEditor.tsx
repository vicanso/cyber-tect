import { NCard, NPageHeader, NSpin, useMessage } from "naive-ui";
import { defineComponent, PropType, ref, Ref } from "vue";
import { showError, showWarning } from "../helpers/util";
import {
  Config,
  configAdd,
  configFindByID,
  ConfigStatus,
  configUpdateByID,
} from "../states/configs";
import ExForm from "./ExForm";
import { FormItem, FormItemTypes } from "./ExFormInterface";
import ExLoading from "./ExLoading";

export function getDefaultFormItems(params: {
  category: string;
  name?: string;
}): FormItem[] {
  return [
    {
      name: "名称：",
      key: "name",
      disabled: params.name != null,
      defaultValue: params.name,
      placeholder: "请输入配置名称",
    },
    {
      name: "分类：",
      key: "category",
      disabled: true,
      defaultValue: params.category,
    },
    {
      name: "状态：",
      key: "status",
      type: FormItemTypes.Select,
      placeholder: "请选择配置状态",
      options: [
        {
          label: "启用",
          value: ConfigStatus.Enabled,
        },
        {
          label: "禁用",
          value: ConfigStatus.Disabled,
        },
      ],
    },
    {
      name: "开始时间：",
      key: "startedAt",
      type: FormItemTypes.DateTime,
      placeholder: "请选择配置生效开始时间",
    },
    {
      name: "结束时间：",
      key: "endedAt",
      type: FormItemTypes.DateTime,
      placeholder: "请选择配置生效结束时间",
    },
  ];
}

function convertDataToConfig(data: Record<string, unknown>): Config {
  // 转换配置数据
  const dataKeyPrefix = "data.";
  const configData: Record<string, unknown> = {};
  Object.keys(data).forEach((key) => {
    if (!key.startsWith(dataKeyPrefix)) {
      return;
    }
    configData[key.substring(dataKeyPrefix.length)] = data[key];
  });
  let configDataStr = data.data as string;
  if (Object.keys(configData).length !== 0) {
    configDataStr = JSON.stringify(configData);
  }
  return {
    name: data.name,
    status: data.status,
    category: data.category,
    startedAt: data.startedAt,
    endedAt: data.endedAt,
    data: configDataStr,
    description: data.description,
  } as Config;
}

function diffConfig(
  newConfig: Config,
  current: Config
): Record<string, unknown> {
  const data: Record<string, unknown> = {};
  if (newConfig.name != current.name) {
    data.name = newConfig.name;
  }
  if (newConfig.status != current.status) {
    data.status = newConfig.status;
  }
  if (newConfig.category !== current.category) {
    data.category = newConfig.category;
  }
  if (newConfig.startedAt !== current.startedAt) {
    data.startedAt = newConfig.startedAt;
  }
  if (newConfig.endedAt !== current.endedAt) {
    data.endedAt = newConfig.endedAt;
  }
  if (newConfig.data !== current.data) {
    data.data = newConfig.data;
  }
  if (newConfig.description !== current.description) {
    data.description = newConfig.description;
  }
  return data;
}

function noop(): void {
  // 无操作
}

export default defineComponent({
  name: "ExConfigEditor",
  props: {
    id: {
      type: Number,
      default: 0,
    },
    title: {
      type: String,
      required: true,
    },
    description: {
      type: String,
      required: true,
    },
    formItems: {
      type: Array as PropType<FormItem[]>,
      required: true,
    },
    onSubmitDone: {
      type: Function as PropType<() => void>,
      default: noop,
    },
    onBack: {
      type: Function as PropType<() => void>,
      default: noop,
    },
  },
  setup(props) {
    const message = useMessage();
    const isUpdatedMode = props.id !== 0;
    const processing = ref(false);
    const currentConfig: Ref<Config> = ref({} as Config);
    // 提交数据
    const onSubmit = async (data: Record<string, unknown>) => {
      if (processing.value) {
        return;
      }
      const { name, category, status, startedAt, endedAt } = data;
      if (!name || !category || !status || !startedAt || !endedAt) {
        showWarning(
          message,
          "配置名称、分类、状态、开始时间、结束时间均不能为空"
        );
        return;
      }

      try {
        processing.value = true;
        const configData = convertDataToConfig(data);
        if (isUpdatedMode) {
          const updateData = diffConfig(configData, currentConfig.value);
          if (Object.keys(updateData).length === 0) {
            showWarning(message, "数据未修改无需要更新");
            return;
          }
          await configUpdateByID({
            id: props.id,
            data: updateData,
          });
          currentConfig.value = configData;
        } else {
          await configAdd(configData);
        }
        props.onSubmitDone();
      } catch (err) {
        showError(message, err);
      } finally {
        processing.value = false;
      }
    };
    const items: FormItem[] = [];
    // 由于会对form item的元素填写默认值，因此重新clone
    props.formItems.forEach((item) => {
      items.push(Object.assign({}, item));
    });
    items.push({
      name: "配置描述：",
      key: "description",
      type: FormItemTypes.TextArea,
      placeholder: "请输入配置描述",
      span: 24,
    });
    // 如果指定了id，则拉取配置
    const fetch = async () => {
      processing.value = true;
      try {
        const data = await configFindByID(props.id);
        currentConfig.value = data;
        items.forEach((item) => {
          if (!item.key) {
            return;
          }
          switch (item.key) {
            case "name":
              item.defaultValue = data.name;
              break;
            case "category":
              item.defaultValue = data.category;
              break;
            case "status":
              item.defaultValue = data.status;
              break;
            case "startedAt":
              item.defaultValue = data.startedAt;
              break;
            case "endedAt":
              item.defaultValue = data.endedAt;
              break;
            case "data":
              item.defaultValue = data.data;
              break;
            case "description":
              item.defaultValue = data.description;
              break;
            default:
              {
                const arr = item.key.split(".");
                if (arr.length === 2 && arr[0] === "data") {
                  try {
                    item.defaultValue = JSON.parse(data.data)[arr[1]];
                  } catch (err) {
                    console.error(err);
                  }
                }
              }
              break;
          }
        });
      } finally {
        processing.value = false;
      }
    };
    if (isUpdatedMode) {
      fetch();
    }

    return {
      currentConfig,
      processing,
      onSubmit,
      items,
    };
  },
  render() {
    const { title, description, id, onBack } = this.$props;
    const { onSubmit, processing, items, currentConfig } = this;
    // 如果指定了id，则展示加载中
    if (processing && id && !currentConfig.id) {
      return <ExLoading />;
    }
    return (
      <NSpin show={processing}>
        <NCard>
          <NPageHeader
            title={title}
            onBack={onBack == noop ? undefined : onBack}
            subtitle={description}
          >
            <ExForm
              formItems={items}
              onSubmit={onSubmit}
              submitText={id !== 0 ? "更新" : "添加"}
            />
          </NPageHeader>
        </NCard>
      </NSpin>
    );
  },
});
