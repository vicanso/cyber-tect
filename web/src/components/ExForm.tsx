import {
  NButton,
  NDatePicker,
  NDynamicInput,
  NForm,
  NFormItem,
  NGrid,
  NGridItem,
  NInput,
  NInputGroup,
  NInputGroupLabel,
  NInputNumber,
  NSelect,
} from "naive-ui";
import { Value } from "naive-ui/lib/select/src/interface";
import { Component, defineComponent, PropType, ref } from "vue";

import ExUserSelect from "./ExUserSelect";
import { FormItem, FormItemTypes } from "./ExFormInterface";

function durationToSeconds(d: string) {
  if (!d || d.length < 2) {
    return null;
  }
  const units = ["s", "m", "h"];
  const seconds = [1, 60, 3600];
  const index = units.indexOf(d[d.length - 1]);
  if (index === -1) {
    return 0;
  }

  return Number.parseInt(d) * seconds[index];
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
        case FormItemTypes.MultiUserSelect:
          component = (
            <ExUserSelect
              formItem={item}
              multiple={true}
              size={size}
              onUpdateValue={(value) => {
                params[item.key] = value;
              }}
            />
          );
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
        case FormItemTypes.InputNumberGroup:
          {
            component = (
              <NInputGroup>
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
                {item.suffix && (
                  <NInputGroupLabel size={size}>{item.suffix}</NInputGroupLabel>
                )}
              </NInputGroup>
            );
          }
          break;
        case FormItemTypes.InputDuration:
          {
            component = (
              <NInputGroup>
                <NInputNumber
                  class="widthFull"
                  disabled={item.disabled || false}
                  size={size}
                  placeholder={item.placeholder}
                  defaultValue={
                    (durationToSeconds(item.defaultValue as string) ||
                      null) as number
                  }
                  onUpdate:value={(value) => {
                    params[item.key] = `${value}s`;
                  }}
                />
                <NInputGroupLabel size={size}>秒</NInputGroupLabel>
              </NInputGroup>
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
        case FormItemTypes.DynamicInput:
          {
            component = (
              <NDynamicInput
                placeholder={item.placeholder}
                defaultValue={(item.defaultValue || []) as []}
                onUpdateValue={(value) => {
                  params[item.key] = value;
                }}
                min={item.min}
                max={item.max}
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
