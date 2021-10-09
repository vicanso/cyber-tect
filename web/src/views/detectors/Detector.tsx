import { defineComponent, ref, PropType } from "vue";
import { NButton, NCard } from "naive-ui";
import { css } from "@linaria/core";
import { TableColumn } from "naive-ui/lib/data-table/src/interface";

import ExDetectorEditor from "../../components/ExDetectorEditor";
import { getDefaultForItems } from "../../components/ExDetectorEditor";
import { FormItem } from "../../components/ExFormInterface";
import { Mode } from "../../states/common";
import ExTable, { newOPColumn } from "../../components/ExTable";
import { padding } from "../../constants/style";
import { getDefaultColumns } from "../../components/ExDetectorTable";
import useUserState from "../../states/user";

const addButtonClass = css`
  width: 100%;
  margin-top: ${2 * padding}px;
`;

export default defineComponent({
  name: "Detector",
  props: {
    title: {
      type: String,
      required: true,
    },
    description: {
      type: String,
      required: true,
    },
    columns: {
      type: Array as PropType<TableColumn[]>,
      required: true,
    },
    formItems: {
      type: Array as PropType<FormItem[]>,
      required: true,
    },
    fetch: {
      type: Function,
      required: true,
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
    data: {
      type: Object,
      required: true,
    },
  },
  setup() {
    const mode = ref(Mode.List);
    const updatedID = ref(0);

    return {
      mode,
      updatedID,
      userInfo: useUserState().info,
    };
  },
  render() {
    const {
      data,
      findByID,
      updateByID,
      create,
      title,
      description,
      formItems,
      columns,
      fetch,
    } = this.$props;
    const { userInfo } = this;
    const columnsClone = getDefaultColumns().slice(0);
    columnsClone.splice(1, 0, ...columns);
    columnsClone.push(
      newOPColumn((row) => {
        this.mode = Mode.Update;
        this.updatedID = row.id as number;
      })
    );
    const { mode, updatedID } = this;
    if (mode === Mode.List) {
      return (
        <NCard title={title}>
          <ExTable columns={columnsClone} data={data} fetch={fetch}></ExTable>
          <NButton
            size="large"
            onClick={() => {
              this.updatedID = 0;
              this.mode = Mode.Add;
            }}
            class={addButtonClass}
          >
            增加检测配置
          </NButton>
        </NCard>
      );
    }
    const formItemsClone = getDefaultForItems();
    formItemsClone.forEach((item) => {
      if (["owners", "receivers"].includes(item.key)) {
        item.defaultValue = [userInfo.account];
      }
    });
    formItemsClone.push(...formItems);
    return (
      <ExDetectorEditor
        id={updatedID}
        findByID={findByID}
        updateByID={updateByID}
        create={create}
        title={title}
        description={description}
        formItems={formItemsClone}
        onBack={() => {
          this.mode = Mode.List;
        }}
      />
    );
  },
});
