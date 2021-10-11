import { NCard, NPageHeader, NSpin, useMessage } from "naive-ui";
import { defineComponent, PropType, ref } from "vue";

import { ConfigStatus } from "../states/configs";
import ExForm from "./ExForm";
import { FormItem, FormItemTypes } from "./ExFormInterface";
import ExLoading from "./ExLoading";
import { diff, showError, showWarning } from "../helpers/util";

function noop(): void {
  // 无操作
}

export function getDefaultForItems(): FormItem[] {
  return [
    {
      name: "名称：",
      key: "name",
      placeholder: "请输入检测配置名称",
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
      name: "超时时长：",
      key: "timeout",
      type: FormItemTypes.InputDuration,
      placeholder: "请输入超时时长",
    },
    {
      name: "用户列表：",
      key: "owners",
      type: FormItemTypes.MultiUserSelect,
      placeholder: "请选择可以修改此配置的用户",
      span: 12,
    },
    {
      name: "告警接收：",
      key: "receivers",
      type: FormItemTypes.MultiUserSelect,
      placeholder: "请选择接收此告警的用户",
      span: 12,
    },
  ];
}

function fillItems(items: FormItem[], data: Record<string, unknown>) {
  items.forEach((item) => {
    if (!item.key) {
      return;
    }
    const value = data[item.key];
    item.defaultValue = value;
  });
}

export default defineComponent({
  name: "ExDetectorEditor",
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
    onBack: {
      type: Function as PropType<() => void>,
      default: noop,
    },
    findByID: {
      type: Function as PropType<
        (id: number) => Promise<Record<string, unknown>>
      >,
      required: true,
    },
    updateByID: {
      type: Function as PropType<
        (id: number, updateData: Record<string, unknown>) => Promise<unknown>
      >,
      required: true,
    },
    create: {
      type: Function as PropType<
        (data: Record<string, unknown>) => Promise<unknown>
      >,
      required: true,
    },
  },
  setup(props) {
    const message = useMessage();
    const fetching = ref(false);
    const processing = ref(false);
    const items: FormItem[] = [];
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
    let originalData: Record<string, unknown>;
    const fetch = async () => {
      if (fetching.value) {
        return;
      }
      fetching.value = true;
      try {
        originalData = await props.findByID(props.id);

        fillItems(items, originalData);
      } catch (err) {
        showError(message, err);
        ``;
      } finally {
        fetching.value = false;
      }
    };
    const submit = async (data: Record<string, unknown>) => {
      if (processing.value) {
        return;
      }
      try {
        processing.value = true;
        // 添加数据
        if (!props.id) {
          await props.create(data);
          props.onBack();
          return;
        }
        const result = diff(data, originalData);
        if (!result.modifiedCount) {
          showWarning(message, "数据未修改无需要更新");
          return;
        }
        await props.updateByID(props.id, result.data);
        props.onBack();
      } catch (err) {
        showError(message, err);
      } finally {
        processing.value = false;
      }
    };
    if (props.id) {
      fetch();
    }
    return {
      fetching,
      processing,
      items,
      submit,
    };
  },
  render() {
    const { title, description, id, onBack } = this.$props;
    const { fetching, processing, items, submit } = this;
    // 如果指定了id，则展示加载中
    if (fetching && id) {
      return <ExLoading />;
    }

    return (
      <NSpin show={fetching || processing}>
        <NCard>
          <NPageHeader
            title={title}
            onBack={onBack == noop ? undefined : onBack}
            subtitle={description}
          >
            <ExForm
              formItems={items}
              onSubmit={submit}
              submitText={id !== 0 ? "更新" : "添加"}
            />
          </NPageHeader>
        </NCard>
      </NSpin>
    );
  },
});
