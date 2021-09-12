import {
  NButton,
  NDatePicker,
  NForm,
  NFormItem,
  NGrid,
  NGridItem,
  NInput,
  NInputNumber,
  NSelect,
} from "naive-ui";
import { Value } from "naive-ui/lib/select/src/interface";
import { Component, defineComponent, PropType, ref } from "vue";

export interface FormItem {
  name: string;
  key: string;
  type?: string;
  placeholder?: string;
  span?: number;
  defaultValue?: unknown;
  disabled?: boolean;
  // TODO 确认是否有其它方式表示
  // eslint-disable-next-line
  options?: any[];
}

export enum FormItemTypes {
  Select = "select",
  MultiSelect = "multiselect",
  DateTime = "datetime",
  DateRange = "datrange",
  InputNumber = "inputnumber",
  TextArea = "textarea",
  Blank = "blank",
}

export default defineComponent({
  name: "ExForm",
  props: {
    formItems: {
      type: Array as PropType<FormItem[]>,
      required: true,
    },
    onSubmit: {
      type: Function as PropType<
        (data: Record<string, unknown>) => Promise<void>
      >,
      required: true,
    },
    submitText: {
      type: String,
      default: "提交",
    },
  },
  setup(props) {
    const params = ref({} as Record<string, unknown>);
    props.formItems.forEach((item) => {
      if (item.defaultValue) {
        params.value[item.key] = item.defaultValue;
      }
    });
    return {
      params,
    };
  },
  render() {
    const { onSubmit, submitText } = this.$props;
    const { params } = this;
    const size = "large";
    const createSelect = (item: FormItem, multiple: boolean) => {
      return (
        <NSelect
          filterable
          multiple={multiple}
          defaultValue={item.defaultValue as Value}
          size={size}
          options={item.options || []}
          placeholder={item.placeholder}
          onUpdateValue={(value) => {
            params[item.key] = value;
          }}
        />
      );
    };
    const formItems = this.$props.formItems as FormItem[];
    const arr = formItems.map((item) => {
      let component: Component;
      switch (item.type) {
        case FormItemTypes.Blank:
          component = <div />;
          break;
        case FormItemTypes.MultiSelect:
          component = createSelect(item, true);
          break;
        case FormItemTypes.Select:
          component = createSelect(item, false);
          break;
        case FormItemTypes.DateTime:
          {
            let defaultValue = null;
            if (item.defaultValue) {
              defaultValue = new Date(item.defaultValue as string).getTime();
            }
            component = (
              <NDatePicker
                type="datetime"
                class="widthFull"
                size={size}
                placeholder={item.placeholder}
                defaultValue={defaultValue}
                clearable
                onUpdateValue={(value) => {
                  if (!value) {
                    params[item.key] = "";
                  } else {
                    params[item.key] = new Date(value).toISOString();
                  }
                }}
              />
            );
          }
          break;
        case FormItemTypes.InputNumber:
          {
            component = (
              <NInputNumber
                class="widthFull"
                disabled={item.disabled || false}
                size={size}
                placeholder={item.placeholder}
                defaultValue={(item.defaultValue || null) as number}
                onUpdate:value={(value) => {
                  params[item.key] = value;
                }}
              />
            );
          }
          break;
        case FormItemTypes.TextArea:
          {
            component = (
              <NInput
                type="textarea"
                autosize={{
                  minRows: 3,
                  maxRows: 5,
                }}
                disabled={item.disabled || false}
                size={size}
                placeholder={item.placeholder}
                defaultValue={(item.defaultValue || "") as string}
                onUpdateValue={(value) => {
                  params[item.key] = value;
                }}
                clearable
              />
            );
          }
          break;
        default:
          component = (
            <NInput
              disabled={item.disabled || false}
              size={size}
              placeholder={item.placeholder}
              defaultValue={(item.defaultValue || "") as string}
              onUpdateValue={(value) => {
                params[item.key] = value;
              }}
              clearable
            />
          );
          break;
      }
      return (
        <NGridItem span={item.span || 8}>
          <NFormItem label={item.name}>{component}</NFormItem>
        </NGridItem>
      );
    });
    arr.push(
      <NGridItem span={24}>
        <NButton size={size} class="widthFull" onClick={() => onSubmit(params)}>
          {submitText}
        </NButton>
      </NGridItem>
    );
    return (
      <NForm labelPlacement="left">
        <NGrid xGap={24}>{arr}</NGrid>
      </NForm>
    );
  },
});
