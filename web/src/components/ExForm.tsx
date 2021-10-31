import {
  FormRules,
  FormInst,
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
  useMessage,
} from "naive-ui";
import { Value } from "naive-ui/lib/select/src/interface";
import { Component, defineComponent, PropType, ref } from "vue";
import { set } from "lodash-es";

import ExUserSelect from "./ExUserSelect";
import { FormItem, FormItemTypes } from "./ExFormInterface";
import { durationToSeconds } from "../helpers/util";
import { showError } from "../helpers/util";

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
    rules: {
      type: Object as PropType<FormRules>,
      default: null,
    },
  },
  setup(props) {
    const formRef = ref({} as FormInst);
    const params = ref({} as Record<string, unknown>);
    props.formItems.forEach((item) => {
      if (item.defaultValue) {
        set(params.value, item.key, item.defaultValue);
      }
    });
    const formValidate = (): Promise<void> => {
      return new Promise((resolve, reject) => {
        formRef.value.validate((errors) => {
          if (errors) {
            const msgList = errors.map((arr) => {
              return arr.map((item) => item.message).join(",");
            });
            reject(new Error(msgList.join(";")));
            return;
          }
          resolve();
        });
      });
    };
    const message = useMessage();

    return {
      handleSubmit: async (data: Record<string, unknown>) => {
        try {
          await formValidate();
        } catch (err) {
          // 如果出错，则不再提交
          showError(message, err);
          return;
        }
        return props.onSubmit(data);
      },
      params,
      formRef,
    };
  },
  render() {
    const { submitText, rules } = this.$props;
    const { params, handleSubmit } = this;
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
            set(params, item.key, value);
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
                set(params, item.key, value);
              }}
              clearable
            />
          );
          break;
      }
      return (
        <NGridItem span={item.span || 8}>
          <NFormItem label={item.name} path={item.key}>
            {component}
          </NFormItem>
        </NGridItem>
      );
    });
    arr.push(
      <NGridItem span={24}>
        <NButton
          size={size}
          class="widthFull"
          onClick={() => handleSubmit(params)}
        >
          {submitText}
        </NButton>
      </NGridItem>
    );
    return (
      <NForm labelPlacement="left" rules={rules} model={params} ref="formRef">
        <NGrid xGap={24}>{arr}</NGrid>
      </NForm>
    );
  },
});
