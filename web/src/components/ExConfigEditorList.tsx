// 配置的列表展示与更新

import { NButton, NCard } from "naive-ui";
import { css } from "@linaria/core";
import { defineComponent, PropType, ref } from "vue";
import { padding } from "../constants/style";
import ExConfigEditor, { getDefaultFormItems } from "./ExConfigEditor";
import ExConfigTable from "./ExConfigTable";
import { FormItem } from "./ExForm";

const addButtonClass = css`
  width: 100%;
  margin-top: ${2 * padding}px;
`;

const listMode = "list";
const addMode = "add";
const updateMode = "update";

export default defineComponent({
  name: "ExConfigEditorList",
  props: {
    listTitle: {
      type: String,
      required: true,
    },
    editorTitle: {
      type: String,
      required: true,
    },
    editorDescription: {
      type: String,
      required: true,
    },
    category: {
      type: String,
      required: true,
    },
    extraFormItems: {
      type: Array as PropType<FormItem[]>,
      default: () => [],
    },
  },
  setup() {
    const mode = ref(listMode);
    const updatedID = ref(0);
    const toggle = (value: string) => {
      mode.value = value;
    };
    return {
      updatedID,
      toggle,
      mode,
    };
  },
  render() {
    const {
      listTitle,
      editorTitle,
      category,
      editorDescription,
      extraFormItems,
    } = this.$props;
    const { mode, toggle, updatedID } = this;
    if (mode === listMode) {
      return (
        <NCard title={listTitle}>
          <ExConfigTable
            category={category}
            onUpdate={(id: number) => {
              this.updatedID = id;
              toggle(updateMode);
            }}
          />
          <NButton
            size="large"
            class={addButtonClass}
            onClick={() => {
              this.updatedID = 0;
              toggle(addMode);
            }}
          >
            增加配置
          </NButton>
        </NCard>
      );
    }

    const formItems = getDefaultFormItems({
      category,
    });
    extraFormItems.forEach((item) => {
      const data = Object.assign({}, item);
      formItems.push(data as FormItem);
    });

    return (
      <ExConfigEditor
        title={editorTitle}
        description={editorDescription}
        id={updatedID}
        formItems={formItems}
        onSubmitDone={() => {
          toggle(listMode);
        }}
        onBack={() => {
          toggle(listMode);
        }}
      ></ExConfigEditor>
    );
  },
});
